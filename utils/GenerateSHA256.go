package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateSHA256(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	shaSecret := hex.EncodeToString(hash.Sum(nil))

	return shaSecret
}
