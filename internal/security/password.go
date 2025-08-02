package security

import (
	"crypto/md5"
	"encoding/hex"
)

func HashPassword(raw string) string {
	sum := md5.Sum([]byte(raw))
	return hex.EncodeToString(sum[:])
}

func CheckPassword(raw, hashed string) bool {
	return HashPassword(raw) == hashed
}
