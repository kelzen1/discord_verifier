package utils

import (
	"crypto/md5"
	"encoding/hex"
	"os"
)

func HashMD5(text string) string {
	text += os.Getenv("SECRET")
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
