package utils

import (
	"encoding/hex"
	"math/rand"
	"time"
)

func GenerateSecureToken(length int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
