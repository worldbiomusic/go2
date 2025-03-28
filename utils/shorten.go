package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"log"
	"strconv"
)

func ShortCode(origin string) string {
	URL_LENGTH, err := strconv.Atoi(ReadEnv("URL_LENGTH"))
	if err != nil {
		log.Println("URL_LENGTH 값 변환 오류")
		return ""
	}

	return makeShort(origin, URL_LENGTH)
}

func makeShort(origin string, length int) string {
	hash := sha256.Sum256([]byte(origin))
	base64Hash := base64.StdEncoding.EncodeToString(hash[:])
	return base64Hash[:length]
}
