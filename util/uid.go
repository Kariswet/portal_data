package util

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateID() (resp string) {
	resp = strings.ReplaceAll(uuid.New().String(), "-", "")
	return
}
