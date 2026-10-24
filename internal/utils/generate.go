package utils

import (
	"crypto/rand"
	"log"
)

const ALPHANUMERIC = "abcdefghijklmnopqrstubwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString(size uint8) string {
	if size == 0 {
		return ""
	}

	randomString := make([]byte, size)
	randomBytes := make([]byte, size)

	if _, err := rand.Read(randomBytes); err != nil {
		log.Panicf("[RandomString] Error reading random bytes: %v", err)
	}

	for i, randomByte := range randomBytes {
		randomString[i] = ALPHANUMERIC[int(randomByte)%len(ALPHANUMERIC)]
	}

	return string(randomString)
}
