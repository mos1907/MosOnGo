package main

import (
	"ProjectsOnGo/Structs"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/basgys/goxml2json"
	"github.com/natefinch/lumberjack"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"time"
	"unicode/utf16"
)

const Upperport = 10541
const Lowerport = 10540
const MaxBufferSize = 65535

var msgID int
var MosID string
var NcsID string
var MosServerIP string
var RundownExportPath string
var config Config

func main() {
	ReadConfig()
	//go HeartBeatEvent()
	listenFromServer()
}

type Config struct {
	MsgID             int    `json:"msgID"`
	MosID             string `json:"MosID"`
	NcsID             string `json:"NcsID"`
	MosServerIP       string `json:"MosServerIP"`
	RundownExportPath string `json:"RundownExportPath"`
}

// Reads the config file and assigns the values to the variables
func ReadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Error reading file:", err.Error())
		return
	}
	defer file.Close()
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err.Error())
		return
	}
	json.Unmarshal(byteValue, &config)
	msgID = config.MsgID
	MosID = config.MosID
	NcsID = config.NcsID
	MosServerIP = config.MosServerIP
	RundownExportPath = config.RundownExportPath
}
func HeartBeatEvent() {
	logFile := &lumberjack.Logger{
		Filename:   "Logs/HeratBeatEvent.log",
		MaxSize:    1,    // 1MB'dan büyük olunca log dosyası rotate ediliyor
		MaxBackups: 250,  // En fazla 3 adet log dosyası tutulacak
		MaxAge:     28,   // 28 gün önce oluşturulmuş log dosyaları silinecek
		Compress:   true, // Gzip ile sıkıştırma yapılacak
	}
	defer logFile.Close()
	// log dosyasını kullanarak logger oluşturma
	logger := log.New(logFile, "", log.LstdFlags)
	// log mesajları yazma
	conn, err := net.Dial("tcp", MosServerIP+":10541")
	if err != nil {
		fmt.Printf("Error connecting to Upperport %d: %v\n", Upperport, err)
		logger.Printf("Error connecting to Upperport %d: %v\n", Upperport, err)
		return
	}
	defer conn.Close()
	conn2, err := net.Dial("tcp", MosServerIP+":10540")
	if err != nil {
		fmt.Printf("Error connecting to Lowerport %d: %v\n", Lowerport, err)
		logger.Printf("Error connecting to Lowerport %d: %v\n", Lowerport, err)
		return
	}
	defer conn2.Close()
	for {
		now := time.Now()
		timeString := now.Format("2006-01-02T15:04:05")
		parsedTime, err := time.Parse("2006-01-02T15:04:05", timeString)
		//xmlString := fmt.Sprintf(XML_STRING, timeString)
		var mos = &Structs.Mos{
			MosID: MosID,
			NcsID: NcsID,
			Heartbeat: &Structs.Heartbeat{
				Time: parsedTime,
			},
		}
		data, err := xml.Marshal(mos)
		// Convert string to UCS-2 encoding
		runes := utf16.Encode([]rune(string(data)))
		bytes := make([]byte, len(runes)*2)
		for i, r := range runes {
			binary.BigEndian.PutUint16(bytes[i*2:], r)
		}
		_, err = conn.Write(bytes)
		if err != nil {
			fmt.Printf("Error sending MOS message to Upperport %d: %v\n", Upperport, err)
			logger.Printf("Error sending MOS message to Upperport ")
			return
		}
		_, err = conn2.Write(bytes)
		if err != nil {
			fmt.Printf("Error sending MOS message to Lowerport %d: %v\n", Lowerport, err)
			logger.Printf("Error sending MOS message to Lowerport")
			return
		}

		respBytes := make([]byte, 1024)
		n, err := conn.Read(respBytes)
		if err != nil {
			fmt.Printf("Error reading MOS response from Upperport %d: %v\n", Upperport, err)
			logger.Printf("Error reading MOS response from Upperport")
			return
		}
		n, err = conn2.Read(respBytes)
		if err != nil {
			fmt.Printf("Error reading MOS response from Lowerport %d: %v\n", Lowerport, err)
			logger.Printf("Error reading MOS response from Lowerport ")
			return
		}

		respString := string(respBytes[:n])
		fmt.Printf("Response from Upperport %d: %s\n", Upperport, respString)
		logger.Printf("HeartBeat Response came from Upperport")
		respString2 := string(respBytes[:n])
		fmt.Printf("Response from Lowerport %d: %s\n", Lowerport, respString2)
		logger.Printf("heartBeat Response came from Lowerport")
		time.Sleep(5 * time.Second)
	}
}
func listenFromServer() {

	// log mesajları yazma

	fmt.Println("Server Running...")
	logx("Server Running...")

	server, err := net.Listen("tcp", ":10541")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		logx("Error listening:" + err.Error())
		os.Exit(1)
	}
	defer server.Close()

	fmt.Println("Listening on  :10541")
	logx("Listening on  :10541")
	fmt.Println("Waiting for client...")
	logx("Waiting for client...")
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			logx("Error accepting: " + err.Error())
			os.Exit(1)
		}
		fmt.Println("client connected")
		go processClient(connection)
	}

}
func processClient(connection net.Conn) {
	defer connection.Close()
	// read Data from client
	buf := make([]byte, 65535)
	n, err := connection.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		logx("Error reading:" + err.Error())
		return
	}
	data := buf[:n]
	// Clear null bytes
	data = bytes.ReplaceAll(data, []byte{0}, []byte{})
	checkData(data)

}
func checkData(data []byte) {
	var mos Structs.Mos

	err := xml.Unmarshal(data, &mos)
	if err != nil {
		fmt.Println("Error unmarshalling XML:", err.Error())
		logx("Error unmarshalling XML:" + err.Error())
		return
	}
	var RunID string
	switch true {
	case mos.RoCreate != nil:
		RunID = mos.RoCreate.RoID
		RecordRundown(RunID, data, err)
		sendAck(RunID) // When Get RoCreate, Send Ack
		fmt.Println("Rundown RoCreate " + RunID)
		logx("Rundown RoCreate " + RunID)
		break
	case mos.RoList != nil:
		RunID = mos.RoList.RoID // when Get RoList, Send Ack
		RecordRundown(RunID, data, err)
		sendAck(RunID)
		fmt.Println("Rundown RoList " + RunID)
		logx("Rundown RoList " + RunID)
		break
	case mos.RoReplace != nil:
		RunID = mos.RoReplace.RoID // when Get RoReplace, Send Ack
		sendAck(RunID)
		RecordRundown(RunID, data, err)
		fmt.Println("Rundown RoReplace " + RunID)
		logx("Rundown RoReplace " + RunID)
		break
	case mos.RoStoryAppend != nil:
		RunID = mos.RoStoryAppend.RoID // when Get RoStoryAppend, Send Ack
		SendRoreq(RunID)
		sendAck(RunID)
		fmt.Println("Rundown RoStoryAppend " + RunID)
		logx("Rundown RoStoryAppend " + RunID)

		break
	case mos.RoStoryDelete != nil:
		RunID = mos.RoStoryDelete.RoID // when Get RoStoryDelete, Send Ack
		SendRoreq(RunID)
		//sendAck(RunID)
		fmt.Println("Rundown RoStoryDelete " + RunID)
		logx("Rundown RoStoryDelete " + RunID)
		break
	case mos.RoStoryReplace != nil:
		RunID = mos.RoStoryReplace.RoID // when Get RoStoryReplace, Send Ack
		SendRoreq(RunID)
		sendAck(RunID)
		fmt.Println("Rundown RoStoryReplace " + RunID)
		logx("Rundown RoStoryReplace " + RunID)
		break
	case mos.RoStoryMove != nil:
		RunID = mos.RoStoryMove.RoID // when Get RoStoryMove, Send Ack
		SendRoreq(RunID)
		sendAck(RunID)
		fmt.Println("Rundown RoStoryMove " + RunID)
		logx("Rundown RoStoryMove " + RunID)
		break
	case mos.RoStorySwap != nil:
		RunID = mos.RoStorySwap.RoID // when Get RoStorySwap, Send Ack
		SendRoreq(RunID)
		sendAck(RunID)
		fmt.Println("Rundown RoStorySwap " + RunID)
		logx("Rundown RoStorySwap " + RunID)

		break
	case mos.RoStoryMoveMultiple != nil:
		RunID = mos.RoStoryMoveMultiple.RoID // when Get RoStoryMoveMultiple, Send Ack
		SendRoreq(RunID)
		sendAckRodelete(RunID)
		fmt.Println("Rundown RoStoryMoveMultiple " + RunID)
		logx("Rundown RoStoryMoveMultiple " + RunID)

		break
	case mos.RoStoryInsert != nil:
		RunID = mos.RoStoryInsert.RoID // when Get RoStoryInsert, Send Ack
		SendRoreq(RunID)
		sendAck(RunID)
		fmt.Println("Rundown RoStoryInsert " + RunID)
		logx("Rundown RoStoryInsert " + RunID)
		break
			
	case mos.RoDelete != nil:
		RunID = mos.RoDelete.RoID // when Get RoDelete, Send Ack
		sendAckRodelete(RunID)
		err := os.Remove(RundownExportPath + RunID + ".xml")
		if err != nil {
			fmt.Println("Error deleting file:", err.Error())
			logx("Error deleting file:" + err.Error())
			return
		} else {
			fmt.Println(RunID, "Rundwon Deleted ", RundownExportPath+RunID+".xml"+" deleted")
			logx(RunID + "Rundwon Deleted " + RundownExportPath + RunID + ".xml" + " deleted")
		}
		err = os.Remove(RundownExportPath + RunID + ".json")
		if err != nil {
			fmt.Println("Error deleting file:", err.Error())
			logx("Error deleting file:" + err.Error())
			return
		} else {
			fmt.Println(RunID, "Rundwon Deleted ", RundownExportPath+RunID+".json"+" deleted")
			logx(RunID + "Rundwon Deleted " + RundownExportPath + RunID + ".json" + " deleted")
		}
	default:

		break
	}
}
func RecordRundown(RunID string, xmlData []byte, err error) {

	filePath := RundownExportPath + RunID + ".xml"
	jsonpath := RundownExportPath + RunID + ".json"
	// write xml to file
	err = ioutil.WriteFile(filePath, xmlData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err.Error())
		logx("Error writing file:" + err.Error())
		return
	} else {
		fmt.Println("Data saved to", filePath)
		logx("Data saved to" + filePath)
	}
	jsonBuffer, err := xml2json.Convert(bytes.NewReader(xmlData))
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		logx("Error converting to JSON:" + err.Error())
		return
	}
	jsonString := jsonBuffer.String()
	err = ioutil.WriteFile(jsonpath, []byte(jsonString), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err.Error())
		logx("Error writing file:" + err.Error())
		return
	} else {
		fmt.Println("Data saved to ", jsonpath)
		logx("Data saved to " + jsonpath)
	}

}
func sendAck(RunID string) {
	conn, err := net.Dial("tcp", MosServerIP+":10541")
	if err != nil {
		fmt.Printf("Error connecting to Upperport %d: %v\n", Upperport, err)
		logx("Error connecting to Upperport %d: %v\n" + err.Error())
		return
	}
	//defer conn.Close()
	// create roAck

	now := time.Now()
	timeString := now.Format("2006-01-02T15:04:05")
	//parsedTime, err := time.Parse("2006-01-02T15:04:05", timeString)
	//xmlString := fmt.Sprintf(XML_STRING, timeString)
	var mos = &Structs.Mos{
		MosID:     MosID,
		NcsID:     NcsID,
		MessageID: msgID,
		RoAck: &Structs.RoAck{
			RoID:     RunID,
			RoStatus: "OK",
		},
	}
	data, err := xml.Marshal(mos)
	// Convert string to UCS-2 encoding
	runes := utf16.Encode([]rune(string(data)))
	bytes := make([]byte, len(runes)*2)
	for i, r := range runes {
		binary.BigEndian.PutUint16(bytes[i*2:], r)
	}
	_, err = conn.Write(bytes)
	if err != nil {
		fmt.Printf("Error sending MOS message to Upperport %d: %v\n", Upperport, err)
		logx("Error sending MOS message to Upperport %d: %v\n" + err.Error())
		return
	} else {
		fmt.Println("RoAck sent " + timeString + " for RundownID: " + RunID)
		logx("RoAck sent " + timeString + " for RundownID: " + RunID)

		msgID++
		config.MsgID = msgID
		jsonValue, _ := json.Marshal(config)
		ioutil.WriteFile("config.json", jsonValue, 0644)
		//write last msgID to log
		fmt.Println("Last msgID: ", msgID)
		logx("Last msgID: " + strconv.Itoa(msgID))
	}
	timer := time.After(1 * time.Second)
	fmt.Println("Waiting for 1 second...")
	logx("Waiting for 1 second Connection close...")
	defer conn.Close()
	// Kanal üzerinde bekleme yapılıyor
	<-timer

}
func sendAckRodelete(RunID string) {
	conn, err := net.Dial("tcp", MosServerIP+":10541")
	if err != nil {
		fmt.Printf("Error connecting to Upperport %d: %v\n", Upperport, err)
		logx("Error connecting to Upperport %d: %v\n" + err.Error())
		return
	}
	now := time.Now()
	timeString := now.Format("2006-01-02T15:04:05")
	Data := "<mos><mosID>" + MosID + "</mosID><ncsID>" + NcsID + "</ncsID><roAck><roID>" + RunID + "</roID><roStatus>OK</roStatus></roAck></mos>"
	runes := utf16.Encode([]rune(string(Data)))
	bytes := make([]byte, len(runes)*2)
	for i, r := range runes {
		binary.BigEndian.PutUint16(bytes[i*2:], r)
	}
	_, err = conn.Write(bytes)
	if err != nil {
		fmt.Printf("Error sending MOS message to Upperport %d: %v\n", Upperport, err)
		logx("Error sending MOS message to Upperport %d: %v\n" + err.Error())
		return
	} else {
		fmt.Println("RoAck sent " + timeString + " for RundownID: " + RunID)
		logx("RoAck sent " + timeString + " for RundownID: " + RunID)
	}
	timer := time.After(1 * time.Second)
	fmt.Println("Waiting for 1 second...")
	logx("Waiting for 1 second Connection close...")
	defer conn.Close()
	// Kanal üzerinde bekleme yapılıyor
	<-timer
}
func SendRoreq(RunID string) {

	conn, err := net.Dial("tcp", MosServerIP+":10541")
	if err != nil {
		fmt.Printf("Error connecting to Upperport %d: %v\n", Upperport, err)
		logx("Error connecting to Upperport %d: %v\n" + err.Error())
		return
	}
	now := time.Now()
	timeString := now.Format("2006-01-02T15:04:05")
	Data := "<mos><mosID>" + MosID + "</mosID><ncsID>" + NcsID + "</ncsID><messageID>" + strconv.Itoa(msgID) + "</messageID><roReq><roID>" + RunID + "</roID></roReq></mos>"
	runes := utf16.Encode([]rune(string(Data)))
	bytesx := make([]byte, len(runes)*2)
	for i, r := range runes {
		binary.BigEndian.PutUint16(bytesx[i*2:], r)
	}
	_, err = conn.Write(bytesx)
	if err != nil {
		fmt.Printf("Error sending MOS message to Upperport %d: %v\n", Upperport, err)
		logx("Error sending MOS message to Upperport %d: %v\n" + err.Error())
		return
	} else {
		fmt.Println("RoReq sent " + timeString + " for RundownID: " + RunID)
		logx("RoReq sent " + timeString + " for RundownID: " + RunID)
	}
	respBytes := make([]byte, 65535)
	n, err := conn.Read(respBytes)
	if err != nil {
		fmt.Printf("Error reading MOS response from Upperport %d: %v\n", Upperport, err)
		logx("Error reading MOS response from Upperport %d: %v\n" + err.Error())
		return
	} else {
		data := respBytes[:n]
		// Clear null bytes
		data = bytes.ReplaceAll(data, []byte{0}, []byte{})
		checkData(data)
	}
	msgID++
	config.MsgID = msgID
	jsonValue, _ := json.Marshal(config)
	ioutil.WriteFile("config.json", jsonValue, 0644)
	//write last msgID to log
	fmt.Println("Last msgID: ", msgID)
	logx("Last msgID: " + strconv.Itoa(msgID))
	timer := time.After(1 * time.Second)
	fmt.Println("Waiting for 1 second...")
	logx("Waiting for 1 second Connection close...")
	defer conn.Close()
	// Kanal üzerinde bekleme yapılıyor
	<-timer
}

func logx(string string) {
	// log mesajları yazma
	logFile := &lumberjack.Logger{
		Filename:   "Logs/MosEvent.log",
		MaxSize:    1,    // 1MB'dan büyük olunca log dosyası rotate ediliyor
		MaxBackups: 250,  // En fazla 3 adet log dosyası tutulacak
		MaxAge:     28,   // 28 gün önce oluşturulmuş log dosyaları silinecek
		Compress:   true, // Gzip ile sıkıştırma yapılacak
	}

	defer logFile.Close()

	// log dosyasını kullanarak logger oluşturma
	logger := log.New(logFile, "", log.LstdFlags)

	// log mesajları yazma
	logger.Println(string)
}
