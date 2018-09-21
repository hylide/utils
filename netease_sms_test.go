package utils

import (
	"git.henghajiang.com/backend/golang_utils/log"
	"testing"
)

var (
	smsLogger = log.New()
)

func TestSendSmsCode(t *testing.T) {
	smsLogger.Info(SendSmsCode("17688165079"))
}

func TestVerifySmsCode(t *testing.T) {
	smsLogger.Info(VerifySmsCode("17688165079", "4478"))
}
