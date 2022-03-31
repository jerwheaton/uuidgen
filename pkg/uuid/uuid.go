package uuid

import (
	"github.com/atotto/clipboard"
	gouuid "github.com/google/uuid"
)

func NewV4() (string, error) {
	newUUID := gouuid.NewString()
	return newUUID, clipboard.WriteAll(newUUID)
}
