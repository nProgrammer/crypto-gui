package controllers

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func SHA256(str string) string {
	str = strings.TrimSpace(str)
	textToHash := []byte(str)
	cryptedH := sha256.Sum256(textToHash)
	text := hex.EncodeToString(cryptedH[:])
	return text
}

func MD5(str string) string {
	str = strings.TrimSpace(str)
	textToHash := []byte(str)
	cryptedH := md5.Sum(textToHash)
	text := hex.EncodeToString(cryptedH[:])
	return text
}
