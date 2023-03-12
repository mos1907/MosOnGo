package Structs

import (
	"encoding/xml"
	"time"
)

// Mos ...
type Mos struct {
	XMLName                 xml.Name                 `xml:"mos"`
	VersionAttr             string                   `xml:"version,attr,omitempty"`
	ChangeDateAttr          string                   `xml:"changeDate,attr,omitempty"`
	NcsAck                  *NcsAck                  `xml:"ncsAck"`
	NcsReqAppInfo           *NcsReqAppInfo           `xml:"ncsReqAppInfo"`
	NcsReqAppMode           *NcsReqAppMode           `xml:"ncsReqAppMode"`
	NcsStoryRequest         *NcsStoryRequest         `xml:"ncsStoryRequest"`
	NcsItemRequest          *NcsItemRequest          `xml:"ncsItemRequest"`
	NcsItem                 *NcsItem                 `xml:"ncsItem"`
	NcsReqAppClose          *NcsReqAppClose          `xml:"ncsReqAppClose"`
	NcsReqStoryAction       string                   `xml:"ncsReqStoryAction"`
	MosID                   string                   `xml:"mosID"`
	NcsID                   string                   `xml:"ncsID"`
	MessageID               int                      `xml:"messageID"`
	MosAck                  *MosAck                  `xml:"mosAck"`
	MosObj                  *MosObj                  `xml:"mosObj"`
	MosReqObj               *MosReqObj               `xml:"mosReqObj"`
	MosReqAll               *MosReqAll               `xml:"mosReqAll"`
	MosListAll              *MosListAll              `xml:"mosListAll"`
	MosObjCreate            *MosObjCreate            `xml:"mosObjCreate"`
	MosItemReplace          *MosItemReplace          `xml:"mosItemReplace"`
	NcsAppInfo              *NcsAppInfo              `xml:"ncsAppInfo"`
	RoAck                   *RoAck                   `xml:"roAck"`
	RoCreate                *RoCreate                `xml:"roCreate"`
	RoReplace               *RoReplace               `xml:"roReplace"`
	RoMetadataReplace       *RoMetadataReplace       `xml:"roMetadataReplace"`
	RoDelete                *RoDelete                `xml:"roDelete"`
	RoReq                   *RoReq                   `xml:"roReq"`
	RoList                  *RoList                  `xml:"roList"`
	RoReqAll                *RoReqAll                `xml:"roReqAll"`
	RoListAll               *RoListAll               `xml:"roListAll"`
	RoStat                  *RoStat                  `xml:"roStat"`
	RoReadyToAir            *RoReadyToAir            `xml:"roReadyToAir"`
	RoStoryAppend           *RoStoryAppend           `xml:"roStoryAppend"`
	RoStoryInsert           *RoStoryInsert           `xml:"roStoryInsert"`
	RoStoryReplace          *RoStoryReplace          `xml:"roStoryReplace"`
	RoStoryMove             *RoStoryMove             `xml:"roStoryMove"`
	RoStorySwap             *RoStorySwap             `xml:"roStorySwap"`
	RoStoryDelete           *RoStoryDelete           `xml:"roStoryDelete"`
	RoStoryMoveMultiple     *RoStoryMoveMultiple     `xml:"roStoryMoveMultiple"`
	RoItemInsert            *RoItemInsert            `xml:"roItemInsert"`
	RoItemReplace           *RoItemReplace           `xml:"roItemReplace"`
	RoItemMoveMultiple      *RoItemMoveMultiple      `xml:"roItemMoveMultiple"`
	RoItemDelete            *RoItemDelete            `xml:"roItemDelete"`
	RoElementAction         string                   `xml:"roElementAction"`
	RoItemStat              *RoItemStat              `xml:"roItemStat"`
	RoElementStat           *RoElementStat           `xml:"roElementStat"`
	RoItemCue               *RoItemCue               `xml:"roItemCue"`
	RoCtrl                  *RoCtrl                  `xml:"roCtrl"`
	RoStorySend             *RoStorySend             `xml:"roStorySend"`
	Heartbeat               *Heartbeat               `xml:"heartbeat"`
	ReqMachInfo             *ReqMachInfo             `xml:"reqMachInfo"`
	ListMachInfo            *ListMachInfo            `xml:"listMachInfo"`
	MosReqSearchableSchema  *MosReqSearchableSchema  `xml:"mosReqSearchableSchema"`
	MosListSearchableSchema *MosListSearchableSchema `xml:"mosListSearchableSchema"`
	MosReqObjList           *MosReqObjList           `xml:"mosReqObjList"`
	MosObjList              *MosObjList              `xml:"mosObjList"`
	MosReqObjAction         *MosReqObjAction         `xml:"mosReqObjAction"`
	RoReqStoryAction        *RoReqStoryAction        `xml:"roReqStoryAction"`
}

// MosAck ...
type MosAck struct {
	XMLName           xml.Name `xml:"mosAck"`
	ObjID             string   `xml:"objID"`
	ObjRev            int      `xml:"objRev"`
	Status            string   `xml:"status"`
	StatusDescription string   `xml:"statusDescription"`
}

// MosObj ...
type MosObj struct {
	XMLName     xml.Name     `xml:"mosObj"`
	ObjID       string       `xml:"objID"`
	ObjSlug     string       `xml:"objSlug"`
	MosAbstract *MosAbstract `xml:"mosAbstract"`
	ObjGroup    string       `xml:"objGroup"`
	ObjType     string       `xml:"objType"`
	ObjTB       int          `xml:"objTB"`
	ObjRev      int          `xml:"objRev"`
	ObjDur      int          `xml:"objDur"`
	Status      string       `xml:"status"`
	ObjAir      string       `xml:"objAir"`
	ObjPaths    *ObjPaths    `xml:"objPaths"`
	CreatedBy   string       `xml:"createdBy"`
	Created     string       `xml:"created"`
	ChangedBy   string       `xml:"changedBy"`
	Changed     string       `xml:"changed"`
	Description *Description `xml:"description"`
}

// MosReqObj ...
type MosReqObj struct {
	XMLName xml.Name `xml:"mosReqObj"`
	ObjID   string   `xml:"objID"`
}

// MosReqAll ...
type MosReqAll struct {
	XMLName xml.Name `xml:"mosReqAll"`
	Pause   int      `xml:"pause"`
}

// MosListAll ...
type MosListAll struct {
	XMLName xml.Name  `xml:"mosListAll"`
	MosObj  []*MosObj `xml:"mosObj"`
}

// MosReqSearchableSchema ...
type MosReqSearchableSchema struct {
	XMLName      xml.Name `xml:"mosReqSearchableSchema"`
	UsernameAttr string   `xml:"username,attr,omitempty"`
}

// MosListSearchableSchema ...
type MosListSearchableSchema struct {
	XMLName      xml.Name `xml:"mosListSearchableSchema"`
	UsernameAttr string   `xml:"username,attr,omitempty"`
	MosSchema    string   `xml:"mosSchema"`
}

// MosReqObjList ...
type MosReqObjList struct {
	XMLName         xml.Name       `xml:"mosReqObjList"`
	UsernameAttr    string         `xml:"username,attr,omitempty"`
	QueryID         string         `xml:"queryID"`
	ListReturnStart int            `xml:"listReturnStart"`
	ListReturnEnd   int            `xml:"listReturnEnd"`
	GeneralSearch   string         `xml:"generalSearch"`
	MosSchema       string         `xml:"mosSchema"`
	SearchGroup     []*SearchGroup `xml:"searchGroup"`
}

// SearchGroup ...
type SearchGroup struct {
	XMLName     xml.Name       `xml:"searchGroup"`
	SearchField []*SearchField `xml:"searchField"`
}

// SearchField ...
type SearchField struct {
	XMLName         xml.Name `xml:"searchField"`
	XPathAttr       string   `xml:"XPath,attr"`
	SortByOrderAttr string   `xml:"sortByOrder,attr,omitempty"`
	SortTypeAttr    string   `xml:"sortType,attr,omitempty"`
}

// MosObjList ...
type MosObjList struct {
	XMLName          xml.Name `xml:"mosObjList"`
	UsernameAttr     string   `xml:"username,attr,omitempty"`
	QueryID          string   `xml:"queryID"`
	ListReturnStart  int      `xml:"listReturnStart"`
	ListReturnEnd    int      `xml:"listReturnEnd"`
	ListReturnTotal  int      `xml:"listReturnTotal"`
	ListReturnStatus string   `xml:"listReturnStatus"`
	List             *List    `xml:"list"`
}

// List ...
type List struct {
	XMLName xml.Name  `xml:"list"`
	MosObj  []*MosObj `xml:"mosObj"`
}

// GeneralSearch ...
type GeneralSearch string

// ListReturnStart ...
type ListReturnStart int

// ListReturnEnd ...
type ListReturnEnd int

// ListReturnTotal ...
type ListReturnTotal int

// ListReturnStatus ...
type ListReturnStatus string

// QueryID ...
type QueryID string

// MosObjCreate ...
type MosObjCreate struct {
	XMLName             xml.Name               `xml:"mosObjCreate"`
	ObjSlug             string                 `xml:"objSlug"`
	ObjGroup            string                 `xml:"objGroup"`
	ObjType             string                 `xml:"objType"`
	ObjTB               int                    `xml:"objTB"`
	ObjDur              int                    `xml:"objDur"`
	Time                time.Time              `xml:"time"`
	CreatedBy           string                 `xml:"createdBy"`
	Description         *Description           `xml:"description"`
	MosExternalMetadata []*MosExternalMetadata `xml:"mosExternalMetadata"`
}

// MosItemReplace ...
type MosItemReplace struct {
	XMLName xml.Name `xml:"mosItemReplace"`
	RoID    string   `xml:"roID"`
	StoryID string   `xml:"storyID"`
	Item    *Item    `xml:"item"`
}

// NcsAck ...
type NcsAck struct {
	XMLName           xml.Name `xml:"ncsAck"`
	Status            string   `xml:"status"`
	StatusDescription string   `xml:"statusDescription"`
}

// NcsAppInfo ...
type NcsAppInfo struct {
	XMLName        xml.Name        `xml:"ncsAppInfo"`
	MosObj         *MosObj         `xml:"mosObj"`
	RoID           string          `xml:"roID"`
	StoryID        string          `xml:"storyID"`
	Item           *Item           `xml:"item"`
	NcsInformation *NcsInformation `xml:"ncsInformation"`
}

// RoAck ...
type RoAck struct {
	XMLName     xml.Name `xml:"roAck"`
	RoID        string   `xml:"roID"`
	RoStatus    string   `xml:"roStatus"`
	StoryID     string   `xml:"storyID"`
	ItemID      string   `xml:"itemID"`
	ObjID       string   `xml:"objID"`
	ItemChannel string   `xml:"itemChannel"`
	Status      string   `xml:"status"`
}

// RoCreate ...
type RoCreate struct {
	XMLName             xml.Name               `xml:"roCreate"`
	RoID                string                 `xml:"roID"`
	RoSlug              string                 `xml:"roSlug"`
	RoChannel           string                 `xml:"roChannel"`
	RoEdStart           string                 `xml:"roEdStart"`
	RoEdDur             string                 `xml:"roEdDur"`
	RoTrigger           string                 `xml:"roTrigger"`
	MacroIn             string                 `xml:"macroIn"`
	MacroOut            string                 `xml:"macroOut"`
	MosExternalMetadata []*MosExternalMetadata `xml:"mosExternalMetadata"`
	Story               []*Story               `xml:"story"`
}

// RoReplace ...
type RoReplace struct {
	XMLName             xml.Name               `xml:"roReplace"`
	RoID                string                 `xml:"roID"`
	RoSlug              string                 `xml:"roSlug"`
	RoChannel           string                 `xml:"roChannel"`
	RoEdStart           string                 `xml:"roEdStart"`
	RoEdDur             string                 `xml:"roEdDur"`
	RoTrigger           string                 `xml:"roTrigger"`
	MacroIn             string                 `xml:"macroIn"`
	MacroOut            string                 `xml:"macroOut"`
	MosExternalMetadata []*MosExternalMetadata `xml:"mosExternalMetadata"`
	Story               []*Story               `xml:"story"`
}

// RoMetadataReplace ...
type RoMetadataReplace struct {
	XMLName             xml.Name             `xml:"roMetadataReplace"`
	RoID                string               `xml:"roID"`
	RoSlug              string               `xml:"roSlug"`
	RoChannel           string               `xml:"roChannel"`
	RoEdStart           string               `xml:"roEdStart"`
	RoEdDur             string               `xml:"roEdDur"`
	RoTrigger           string               `xml:"roTrigger"`
	MacroIn             string               `xml:"macroIn"`
	MacroOut            string               `xml:"macroOut"`
	MosExternalMetadata *MosExternalMetadata `xml:"mosExternalMetadata"`
}

// RoDelete ...
type RoDelete struct {
	XMLName xml.Name `xml:"roDelete"`
	RoID    string   `xml:"roID"`
}

// RoReq ...
type RoReq struct {
	XMLName xml.Name `xml:"roReq"`
	RoID    string   `xml:"roID"`
}

// RoList ...
type RoList struct {
	XMLName             xml.Name               `xml:"roList"`
	RoID                string                 `xml:"roID"`
	RoSlug              string                 `xml:"roSlug"`
	RoChannel           string                 `xml:"roChannel"`
	RoEdStart           string                 `xml:"roEdStart"`
	RoEdDur             string                 `xml:"roEdDur"`
	RoTrigger           string                 `xml:"roTrigger"`
	MacroIn             string                 `xml:"macroIn"`
	MacroOut            string                 `xml:"macroOut"`
	MosExternalMetadata []*MosExternalMetadata `xml:"mosExternalMetadata"`
	Story               []*Story               `xml:"story"`
}

// RoListAll ...
type RoListAll struct {
	XMLName xml.Name `xml:"roListAll"`
	Ro      []*Ro    `xml:"ro"`
}

// Ro ...
type Ro struct {
	XMLName             xml.Name               `xml:"ro"`
	RoID                string                 `xml:"roID"`
	RoSlug              string                 `xml:"roSlug"`
	RoChannel           string                 `xml:"roChannel"`
	RoEdStart           string                 `xml:"roEdStart"`
	RoEdDur             string                 `xml:"roEdDur"`
	RoTrigger           string                 `xml:"roTrigger"`
	MosExternalMetadata []*MosExternalMetadata `xml:"mosExternalMetadata"`
}

// RoStat ...
type RoStat struct {
	XMLName xml.Name  `xml:"roStat"`
	RoID    string    `xml:"roID"`
	Status  string    `xml:"status"`
	Time    time.Time `xml:"time"`
}

// RoReadyToAir ...
type RoReadyToAir struct {
	XMLName xml.Name `xml:"roReadyToAir"`
	RoID    string   `xml:"roID"`
	RoAir   string   `xml:"roAir"`
}

// RoStoryAppend ...
type RoStoryAppend struct {
	XMLName xml.Name `xml:"roStoryAppend"`
	RoID    string   `xml:"roID"`
	Story   []*Story `xml:"story"`
}

// RoStoryInsert ...
type RoStoryInsert struct {
	XMLName xml.Name `xml:"roStoryInsert"`
	RoID    string   `xml:"roID"`
	StoryID string   `xml:"storyID"`
	Story   []*Story `xml:"story"`
}

// RoStoryReplace ...
type RoStoryReplace struct {
	XMLName xml.Name `xml:"roStoryReplace"`
	RoID    string   `xml:"roID"`
	StoryID string   `xml:"storyID"`
	Story   []*Story `xml:"story"`
}

// RoStoryMove ...
type RoStoryMove struct {
	XMLName xml.Name `xml:"roStoryMove"`
	RoID    string   `xml:"roID"`
	StoryID string   `xml:"storyID"`
}

// RoStorySwap ...
type RoStorySwap struct {
	XMLName xml.Name `xml:"roStorySwap"`
	RoID    string   `xml:"roID"`
	StoryID string   `xml:"storyID"`
}

// RoStoryDelete ...
type RoStoryDelete struct {
	XMLName xml.Name `xml:"roStoryDelete"`
	RoID    string   `xml:"roID"`
	StoryID []string `xml:"storyID"`
}

// RoStoryMoveMultiple ...
type RoStoryMoveMultiple struct {
	XMLName xml.Name `xml:"roStoryMoveMultiple"`
	RoID    string   `xml:"roID"`
	StoryID []string `xml:"storyID"`
}

// RoItemInsert ...
type RoItemInsert struct {
	XMLName xml.Name `xml:"roItemInsert"`
	RoID    string   `xml:"roID"`
	StoryID string   `xml:"storyID"`
	ItemID  string   `xml:"itemID"`
	Item    []*Item  `xml:"item"`
}

// RoItemReplace ...
type RoItemReplace struct {
	XMLName xml.Name `xml:"roItemReplace"`
	RoID    string   `xml:"roID"`
	StoryID string   `xml:"storyID"`
	ItemID  string   `xml:"itemID"`
	Item    []*Item  `xml:"item"`
}

// RoItemMoveMultiple ...
type RoItemMoveMultiple struct {
	XMLName xml.Name `xml:"roItemMoveMultiple"`
	RoID    string   `xml:"roID"`
	StoryID string   `xml:"storyID"`
	ItemID  []string `xml:"itemID"`
}

// RoItemDelete ...
type RoItemDelete struct {
	XMLName xml.Name `xml:"roItemDelete"`
	RoID    string   `xml:"roID"`
	StoryID string   `xml:"storyID"`
	ItemID  []string `xml:"itemID"`
}

// RoElementAction ...
type RoElementAction struct {
	XMLName         xml.Name       `xml:"roElementAction"`
	OperationAttr   string         `xml:"operation,attr"`
	RoID            string         `xml:"roID"`
	Elementtarget   *Elementtarget `xml:"element_target"`
	RoElementAction string         `xml:"roElementAction"`
}

// Elementtarget ...
type Elementtarget struct {
	XMLName xml.Name `xml:"element_target"`
	StoryID string   `xml:"storyID"`
	ItemID  string   `xml:"itemID"`
}

// Elementsource ...
type Elementsource struct {
	XMLName xml.Name `xml:"element_source"`
	Story   []*Story `xml:"story"`
	Item    []*Item  `xml:"item"`
	StoryID []string `xml:"storyID"`
	ItemID  []string `xml:"itemID"`
}

// MosReqObjAction ...
type MosReqObjAction struct {
	XMLName             xml.Name               `xml:"mosReqObjAction"`
	OperationAttr       string                 `xml:"operation,attr"`
	ObjIDAttr           string                 `xml:"objID,attr,omitempty"`
	ObjSlug             string                 `xml:"objSlug"`
	MosAbstract         *MosAbstract           `xml:"mosAbstract"`
	ObjGroup            string                 `xml:"objGroup"`
	ObjType             string                 `xml:"objType"`
	ObjTB               int                    `xml:"objTB"`
	ObjDur              int                    `xml:"objDur"`
	Time                time.Time              `xml:"time"`
	CreatedBy           string                 `xml:"createdBy"`
	ChangedBy           string                 `xml:"changedBy"`
	Changed             string                 `xml:"changed"`
	Description         *Description           `xml:"description"`
	MosExternalMetadata []*MosExternalMetadata `xml:"mosExternalMetadata"`
}

// RoReqStoryAction ...
type RoReqStoryAction struct {
	XMLName       xml.Name     `xml:"roReqStoryAction"`
	OperationAttr string       `xml:"operation,attr"`
	LeaseLockAttr string       `xml:"leaseLock,attr,omitempty"`
	UsernameAttr  string       `xml:"username,attr,omitempty"`
	RoStorySend   *RoStorySend `xml:"roStorySend"`
}

// RoStorySend ...
type RoStorySend struct {
	XMLName             xml.Name               `xml:"roStorySend"`
	RoID                string                 `xml:"roID"`
	StoryID             string                 `xml:"storyID"`
	StorySlug           string                 `xml:"storySlug"`
	StoryNum            string                 `xml:"storyNum"`
	StoryBody           *StoryBody             `xml:"storyBody"`
	MosExternalMetadata []*MosExternalMetadata `xml:"mosExternalMetadata"`
}

// RoItemStat ...
type RoItemStat struct {
	XMLName     xml.Name  `xml:"roItemStat"`
	RoID        string    `xml:"roID"`
	StoryID     string    `xml:"storyID"`
	ItemID      string    `xml:"itemID"`
	ObjID       string    `xml:"objID"`
	ItemChannel string    `xml:"itemChannel"`
	Status      string    `xml:"status"`
	Time        time.Time `xml:"time"`
}

// RoElementStat ...
type RoElementStat struct {
	XMLName     xml.Name  `xml:"roElementStat"`
	RoID        string    `xml:"roID"`
	StoryID     string    `xml:"storyID"`
	ItemID      string    `xml:"itemID"`
	ObjID       string    `xml:"objID"`
	ItemChannel string    `xml:"itemChannel"`
	Status      string    `xml:"status"`
	Time        time.Time `xml:"time"`
}

// RoItemCue ...
type RoItemCue struct {
	XMLName             xml.Name               `xml:"roItemCue"`
	MosID               string                 `xml:"mosID"`
	RoID                string                 `xml:"roID"`
	StoryID             string                 `xml:"storyID"`
	ItemID              string                 `xml:"itemID"`
	RoEventType         string                 `xml:"roEventType"`
	RoEventTime         string                 `xml:"roEventTime"`
	MosExternalMetadata []*MosExternalMetadata `xml:"mosExternalMetadata"`
}

// RoCtrl ...
type RoCtrl struct {
	XMLName             xml.Name               `xml:"roCtrl"`
	RoID                string                 `xml:"roID"`
	StoryID             string                 `xml:"storyID"`
	ItemID              string                 `xml:"itemID"`
	Command             string                 `xml:"command"`
	MosExternalMetadata []*MosExternalMetadata `xml:"mosExternalMetadata"`
}

// Heartbeat ...
type Heartbeat struct {
	XMLName xml.Name  `xml:"heartbeat"`
	Time    time.Time `xml:"time"`
}

// ListMachInfo ...
type ListMachInfo struct {
	XMLName             xml.Name               `xml:"listMachInfo"`
	Manufacturer        string                 `xml:"manufacturer"`
	Model               string                 `xml:"model"`
	HwRev               string                 `xml:"hwRev"`
	SwRev               string                 `xml:"swRev"`
	DOM                 string                 `xml:"DOM"`
	SN                  string                 `xml:"SN"`
	ID                  string                 `xml:"ID"`
	Time                time.Time              `xml:"time"`
	OpTime              string                 `xml:"opTime"`
	MosRev              string                 `xml:"mosRev"`
	SupportedProfiles   string                 `xml:"SupportedProfiles"`
	DefaultActiveX      []*DefaultActiveX      `xml:"defaultActiveX"`
	MosExternalMetadata []*MosExternalMetadata `xml:"mosExternalMetadata"`
}

// DefaultActiveX ...
type DefaultActiveX struct {
	XMLName              xml.Name `xml:"defaultActiveX"`
	Mode                 string   `xml:"mode"`
	ControlFileLocation  string   `xml:"controlFileLocation"`
	ControlSlug          string   `xml:"controlSlug"`
	ControlName          string   `xml:"controlName"`
	ControlDefaultParams string   `xml:"controlDefaultParams"`
}

// Story ...
type Story struct {
	XMLName             xml.Name               `xml:"story"`
	StoryID             string                 `xml:"storyID"`
	StorySlug           string                 `xml:"storySlug"`
	StoryNum            string                 `xml:"storyNum"`
	MosExternalMetadata []*MosExternalMetadata `xml:"mosExternalMetadata"`
	Item                []*Item                `xml:"item"`
}

// Description ...
type Description struct {
	XMLName xml.Name `xml:"description"`
	P       []*P     `xml:"p"`
	Em      []string `xml:"em"`
	Tab     []string `xml:"tab"`
}

// StoryBody ...
type StoryBody struct {
	XMLName              xml.Name `xml:"storyBody"`
	Read1stMEMasBodyAttr string   `xml:"Read1stMEMasBody,attr,omitempty"`
	P                    []*P     `xml:"p"`
}

// StoryItem ...
type StoryItem struct {
	XMLName             xml.Name               `xml:"storyItem"`
	ItemID              string                 `xml:"itemID"`
	ItemSlug            string                 `xml:"itemSlug"`
	ObjID               string                 `xml:"objID"`
	MosID               string                 `xml:"mosID"`
	MosAbstract         *MosAbstract           `xml:"mosAbstract"`
	ItemChannel         string                 `xml:"itemChannel"`
	ItemEdStart         int                    `xml:"itemEdStart"`
	ItemEdDur           int                    `xml:"itemEdDur"`
	ItemUserTimingDur   int                    `xml:"itemUserTimingDur"`
	ItemTrigger         string                 `xml:"itemTrigger"`
	MacroIn             string                 `xml:"macroIn"`
	MacroOut            string                 `xml:"macroOut"`
	MosExternalMetadata []*MosExternalMetadata `xml:"mosExternalMetadata"`
}

// Item ...
type Item struct {
	XMLName             xml.Name               `xml:"item"`
	ItemID              string                 `xml:"itemID"`
	ItemSlug            string                 `xml:"itemSlug"`
	ObjID               string                 `xml:"objID"`
	MosID               string                 `xml:"mosID"`
	MosPlugInID         string                 `xml:"mosPlugInID"`
	MosAbstract         *MosAbstract           `xml:"mosAbstract"`
	ObjPaths            *ObjPaths              `xml:"objPaths"`
	ItemChannel         string                 `xml:"itemChannel"`
	ItemEdStart         int                    `xml:"itemEdStart"`
	ItemEdDur           int                    `xml:"itemEdDur"`
	ItemUserTimingDur   int                    `xml:"itemUserTimingDur"`
	ItemTrigger         string                 `xml:"itemTrigger"`
	MacroIn             string                 `xml:"macroIn"`
	MacroOut            string                 `xml:"macroOut"`
	MosExternalMetadata []*MosExternalMetadata `xml:"mosExternalMetadata"`
}

// MosExternalMetadata ...
type MosExternalMetadata struct {
	XMLName    xml.Name    `xml:"mosExternalMetadata"`
	MosScope   string      `xml:"mosScope"`
	MosSchema  string      `xml:"mosSchema"`
	MosPayload *MosPayload `xml:"mosPayload"`
}

// NcsInformation ...
type NcsInformation struct {
	XMLName    xml.Name    `xml:"ncsInformation"`
	UserID     string      `xml:"userID"`
	RunContext string      `xml:"runContext"`
	Software   *Software   `xml:"software"`
	UImetric   []*UImetric `xml:"UImetric"`
}

// UImetric ...
type UImetric struct {
	NumAttr  string `xml:"num,attr"`
	Startx   string `xml:"startx"`
	Starty   string `xml:"starty"`
	Endx     string `xml:"endx"`
	Endy     string `xml:"endy"`
	Mode     string `xml:"mode"`
	CanClose bool   `xml:"canClose"`
}

// NcsItem ...
type NcsItem struct {
	XMLName xml.Name `xml:"ncsItem"`
	Item    *Item    `xml:"item"`
}

// NcsReqStoryAction ...
type NcsReqStoryAction struct {
	XMLName           xml.Name `xml:"ncsReqStoryAction"`
	OperationAttr     string   `xml:"operation,attr"`
	NcsReqStoryAction string   `xml:"ncsReqStoryAction"`
}

// NcsReqAppMode ...
type NcsReqAppMode struct {
	XMLName  xml.Name  `xml:"ncsReqAppMode"`
	UImetric *UImetric `xml:"UImetric"`
}

// Software ...
type Software struct {
	XMLName      xml.Name `xml:"software"`
	Manufacturer string   `xml:"manufacturer"`
	Product      string   `xml:"product"`
	Version      string   `xml:"version"`
}

// B ...
type B struct {
	XMLName xml.Name `xml:"b"`
	I       []*I     `xml:"i"`
	U       []*U     `xml:"u"`
}

// I ...
type I struct {
	XMLName xml.Name `xml:"i"`
	B       []*B     `xml:"b"`
	U       []*U     `xml:"u"`
}

// P ...
type P struct {
	XMLName          xml.Name     `xml:"p"`
	Em               []string     `xml:"em"`
	Tab              []string     `xml:"tab"`
	Pi               []*Pi        `xml:"pi"`
	Pkg              []*Pkg       `xml:"pkg"`
	B                []*B         `xml:"b"`
	I                []*I         `xml:"i"`
	U                []*U         `xml:"u"`
	StoryItem        []*StoryItem `xml:"storyItem"`
	StoryPresenter   []string     `xml:"storyPresenter"`
	StoryPresenterRR []string     `xml:"storyPresenterRR"`
}

// Pi ...
type Pi struct {
	XMLName xml.Name `xml:"pi"`
	B       []*B     `xml:"b"`
	I       []*I     `xml:"i"`
	U       []*U     `xml:"u"`
}

// Pkg ...
type Pkg struct {
	XMLName xml.Name `xml:"pkg"`
	B       []*B     `xml:"b"`
	I       []*I     `xml:"i"`
	U       []*U     `xml:"u"`
}

// U ...
type U struct {
	XMLName xml.Name `xml:"u"`
	B       []*B     `xml:"b"`
	I       []*I     `xml:"i"`
}

// MosID ...
type MosID string

// NcsID ...
type NcsID string

// CanClose ...
type CanClose bool

// Changed ...
type Changed string

// ChangedBy ...
type ChangedBy string

// Command ...
type Command string

// ControlDefaultParams ...
type ControlDefaultParams string

// ControlFileLocation ...
type ControlFileLocation string

// ControlName ...
type ControlName string

// ControlSlug ...
type ControlSlug string

// Created ...
type Created string

// CreatedBy ...
type CreatedBy string

// DOM ...
type DOM string

// Em ...
type Em string

// Endx ...
type Endx string

// Endy ...
type Endy string

// HwRev ...
type HwRev string

// ID ...
type ID string

// ItemChannel ...
type ItemChannel string

// ItemEdDur ...
type ItemEdDur int

// ItemEdStart ...
type ItemEdStart int

// ItemID ...
type ItemID string

// ItemSlug ...
type ItemSlug string

// ItemTrigger ...
type ItemTrigger string

// ItemUserTimingDur ...
type ItemUserTimingDur int

// MacroIn ...
type MacroIn string

// MacroOut ...
type MacroOut string

// Manufacturer ...
type Manufacturer string

// Mode ...
type Mode string

// Model ...
type Model string

// MosAbstract ...
type MosAbstract struct {
	XMLName xml.Name `xml:"mosAbstract"`
}

// MosPayload ...
type MosPayload struct {
	XMLName xml.Name `xml:"mosPayload"`
}

// MosPlugInID ...
type MosPlugInID string

// SupportedProfiles ...
type SupportedProfiles struct {
	DeviceTypeAttr    string `xml:"deviceType,attr"`
	SupportedProfiles string `xml:"SupportedProfiles"`
}

// MosProfile ...
type MosProfile struct {
	XMLName    xml.Name `xml:"mosProfile"`
	NumberAttr int      `xml:"number,attr"`
	Value      bool     `xml:",chardata"`
}

// MosSchema ...
type MosSchema string

// MosScope ...
type MosScope string

// MosRev ...
type MosRev string

// NcsItemRequest ...
type NcsItemRequest struct {
	XMLName xml.Name `xml:"ncsItemRequest"`
}

// NcsReqAppClose ...
type NcsReqAppClose struct {
	XMLName xml.Name `xml:"ncsReqAppClose"`
}

// NcsReqAppInfo ...
type NcsReqAppInfo struct {
	XMLName xml.Name `xml:"ncsReqAppInfo"`
}

// NcsStoryRequest ...
type NcsStoryRequest struct {
	XMLName xml.Name `xml:"ncsStoryRequest"`
}

// ObjAir ...
type ObjAir string

// ObjDur ...
type ObjDur int

// ObjGroup ...
type ObjGroup string

// ObjID ...
type ObjID string

// ObjPaths ...
type ObjPaths struct {
	XMLName         xml.Name           `xml:"objPaths"`
	ObjPath         []*ObjPath         `xml:"objPath"`
	ObjProxyPath    []*ObjProxyPath    `xml:"objProxyPath"`
	ObjMetadataPath []*ObjMetadataPath `xml:"objMetadataPath"`
}

// ObjPath ...
type ObjPath struct {
	XMLName             xml.Name `xml:"objPath"`
	TechDescriptionAttr string   `xml:"techDescription,attr"`
	Value               string   `xml:",chardata"`
}

// ObjProxyPath ...
type ObjProxyPath struct {
	XMLName             xml.Name `xml:"objProxyPath"`
	TechDescriptionAttr string   `xml:"techDescription,attr"`
	Value               string   `xml:",chardata"`
}

// ObjMetadataPath ...
type ObjMetadataPath struct {
	XMLName             xml.Name `xml:"objMetadataPath"`
	TechDescriptionAttr string   `xml:"techDescription,attr"`
	Value               string   `xml:",chardata"`
}

// ObjRev ...
type ObjRev int

// ObjSlug ...
type ObjSlug string

// ObjTB ...
type ObjTB int

// ObjType ...
type ObjType string

// OpTime ...
type OpTime string

// Pause ...
type Pause int

// Product ...
type Product string

// MessageID ...
type MessageID int

// ReqMachInfo ...
type ReqMachInfo struct {
	XMLName xml.Name `xml:"reqMachInfo"`
}

// RoAir ...
type RoAir string

// RoID ...
type RoID string

// RoChannel ...
type RoChannel string

// RoCtrlCmd ...
type RoCtrlCmd string

// RoCtrlTime ...
type RoCtrlTime string

// RoEdDur ...
type RoEdDur string

// RoEdStart ...
type RoEdStart string

// RoEventTime ...
type RoEventTime string

// RoEventType ...
type RoEventType string

// RoReqAll ...
type RoReqAll struct {
	XMLName xml.Name `xml:"roReqAll"`
}

// RoSlug ...
type RoSlug string

// RoStatus ...
type RoStatus string

// RoTrigger ...
type RoTrigger string

// RunContext ...
type RunContext string

// SN ...
type SN string

// Startx ...
type Startx string

// Starty ...
type Starty string

// Status ...
type Status string

// StatusDescription ...
type StatusDescription string

// StoryID ...
type StoryID string

// StoryNum ...
type StoryNum string

// StoryPresenter ...
type StoryPresenter string

// StoryPresenterRR ...
type StoryPresenterRR string

// StorySlug ...
type StorySlug string

// SwRev ...
type SwRev string

// Tab ...
type Tab string

// Time ...
type Time string

// UserID ...
type UserID string

// Version ...
type Version string
