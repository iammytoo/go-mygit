package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
)

func Sha1(text string) string {
	sha1 := sha1.New()
	io.WriteString(sha1, text)
	return hex.EncodeToString(sha1.Sum(nil))
}
