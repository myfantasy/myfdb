package main

import (
	"crypto/sha256"
	"encoding/base64"
)

// HashGet get sha256 convert into base64
func HashGet(s string) (res string) {
	sha := sha256.Sum256([]byte(s))
	return base64.StdEncoding.EncodeToString(sha[:])
}

// HashOfBytesGet get sha256 convert into base64
func HashOfBytesGet(b []byte) (res string) {
	sha := sha256.Sum256(b)
	return base64.StdEncoding.EncodeToString(sha[:])
}
