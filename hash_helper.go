package main

import (
	"crypto/md5"
	"encoding/hex"
)

func HashContent(body []byte) string {
	hasher := md5.New()
	hasher.Write(body)
	return hex.EncodeToString(hasher.Sum(nil))
}
