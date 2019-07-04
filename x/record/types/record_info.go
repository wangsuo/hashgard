package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Record interface
type Record interface {
	GetID() string
	SetID(string)

	GetHash() string
	SetHash(string)

	GetRecordNo() string
	SetRecordNo(string)

	GetSender() sdk.AccAddress
	SetSender(sdk.AccAddress)

	GetRecordTime() int64
	SetRecordTime(int64)

	GetName() string
	SetName(string)

	GetAuthor() string
	SetAuthor(string)

	GetType() string
	SetType(string)

	String() string
}

// Records is an array of Record
type Records []RecordInfo

// Record Info
type RecordInfo struct {
	ID          	   string         `json:"id"`
	Hash               string         `json:"hash"`
	RecordNo           string         `json:"record_number"`
	Sender             sdk.AccAddress `json:"sender"`
	RecordTime         int64          `json:"record_time"`
	Name               string         `json:"name"`
	Author             string         `json:"author"`
	Type               string         `json:"type"`
}

// Implements Record Interface
var _ Record = (*RecordInfo)(nil)

//nolint
func (ci RecordInfo) GetID() string {
	return ci.ID
}
func (ci *RecordInfo) SetID(id string) {
	ci.ID = id
}
func (ci RecordInfo) GetHash() string {
	return ci.Hash
}
func (ci *RecordInfo) SetHash(hash string) {
	ci.Hash = hash
}
func (ci RecordInfo) GetRecordNo() string {
	return ci.RecordNo
}
func (ci *RecordInfo) SetRecordNo(number string) {
	ci.RecordNo = number
}
func (ci RecordInfo) GetSender() sdk.AccAddress {
	return ci.Sender
}
func (ci *RecordInfo) SetSender(operator sdk.AccAddress) {
	ci.Sender = operator
}
func (ci RecordInfo) GetRecordTime() int64 {
	return ci.RecordTime
}
func (ci *RecordInfo) SetRecordTime(time int64) {
	ci.RecordTime = time
}
func (ci RecordInfo) GetName() string {
	return ci.Name
}
func (ci *RecordInfo) SetName(name string) {
	ci.Name = name
}
func (ci RecordInfo) GetAuthor() string {
	return ci.Author
}
func (ci *RecordInfo) SetAuthor(author string) {
	ci.Author = author
}
func (ci RecordInfo) GetType() string {
	return ci.Type
}
func (ci *RecordInfo) SetType(recordType string) {
	ci.Type = recordType
}

//nolint
func (ci RecordInfo) String() string {
	return fmt.Sprintf(`Record:
  ID:		          		    %s
  Hash:       		   		    %s
  RecordNo:          		    %s
  Sender:           			%s
  Name:             			%s
  Author:             	 		%s
  Type:    	    		 		%s`,
		ci.ID, ci.Hash, ci.RecordNo, ci.Sender.String(), ci.Name, ci.Author, ci.Type)
}

//nolint
func (records Records) String() string {
	out := fmt.Sprintf("%-17s|%-17s|%-10s|%-44s|%-10s|%-6s|%-6s|%s\n",
		"Id", "Hash", "RecordNo", "Sender", "Name", "Author", "Type", "RecordTime")
	for _, record := range records {
		out += fmt.Sprintf("%-17s|%-17s|%-10s|%-44s|%-10s|%-6s|%-6s|%d\n",
			record.ID, record.Hash, record.RecordNo, record.GetSender().String(), record.Name, record.Author, record.Type, record.RecordTime)
	}
	return strings.TrimSpace(out)
}
