package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func generateShortUrl(originalUrl string) string {
	hasher := md5.New()
	hasher.Write([]byte(originalUrl))
	data := hasher.Sum(nil)

	hash := hex.EncodeToString(data)

	return hash[:6]
}
