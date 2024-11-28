package models

import (
	"github.com/gofrs/uuid"
)

// UUIDsToBytes
func UUIDsToBytes(uuids []uuid.UUID) (bytes [][]byte) {
	for _, u := range uuids {
		bytes = append(bytes, u.Bytes())
	}
	return
}

// BytesToUUIDs
func BytesToUUIDs(bytes [][]byte) (uuids []uuid.UUID) {
	for _, b := range bytes {
		uuids = append(uuids, uuid.FromBytesOrNil(b))
	}
	return
}
