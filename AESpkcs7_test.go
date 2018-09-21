package utils

import (
	"encoding/base64"
	"git.henghajiang.com/backend/golang_utils/log"
	"testing"
)

const (
	key = "henghajiangwillbeonlineatshanghaiin20180912"
)

var (
	pkcsLogger = log.New()
)

func TestEncrypt(t *testing.T) {
	byteKey, err := base64.StdEncoding.DecodeString(key + "=")
	if err != nil {
		pkcsLogger.Exception(err)
	}
	secret, err := Encrypt([]byte("aabbccddeeffgg"), byteKey)
	if err != nil {
		pkcsLogger.Exception(err)
	}
	pkcsLogger.Info(secret)
	pkcsLogger.Info(Decrypt(secret, byteKey))
}
