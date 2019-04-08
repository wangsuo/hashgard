package keepers

import (
	"unsafe"
)

// Key for getting a the next available proposalID from the store
var (
	KeyDelimiter = []byte(":")
)

func BytesString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// Key for getting a specific issuer from the store
func KeyIssuer(issueID string) []byte {
	return []byte(issueID)
}
