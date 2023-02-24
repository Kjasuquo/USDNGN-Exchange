package utils

import (
	"crypto/sha512"
	"encoding/hex"
)

func ComputeHash(password, salt string) string {
	hasher := sha512.New()
	// TODO: we should throw this error
	_, _ = hasher.Write([]byte(password + salt))
	result := hex.EncodeToString(hasher.Sum(nil))
	return result[:24]
}
