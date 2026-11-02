package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		log.Panicf("Error hashing password: %v", err)
	}
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HmacSHA256(key, message string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

func VerifyHmacSHA256(key, message, hash string) bool {
	expectedHash := HmacSHA256(key, message)
	return hmac.Equal([]byte(expectedHash), []byte(hash))
}
