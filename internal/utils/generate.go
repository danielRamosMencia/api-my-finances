package utils

import (
	"crypto/rand"
	"log"
	"time"
)

const (
	ALPHANUMERIC = "abcdefghijklmnopqrstubwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	NUMERIC      = "0123456789"
)

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

func RandomNumber(size uint, useTimestamp bool) string {
	if size == 0 {
		return ""
	}

	randomString := make([]byte, size)
	randomBytes := make([]byte, size)

	if _, err := rand.Read(randomBytes); err != nil {
		log.Panicf("[RandomNumber] Error reading random bytes: %v", err)
	}

	for i, randomByte := range randomBytes {
		randomString[i] = NUMERIC[int(randomByte)%10]
	}

	timestamp := time.Now().Format("20060102")

	if useTimestamp {
		return timestamp + string(randomString)
	}

	return string(randomString)
}
