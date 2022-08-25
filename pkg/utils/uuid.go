package utils

import (
	guuid "github.com/google/uuid"
)

func GenerateUUID() string {
	return guuid.New().String()
}
