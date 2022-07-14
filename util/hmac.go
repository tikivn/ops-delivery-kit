package util

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/subtle"
	"encoding/hex"
)

func GenerateSignature(secretToken, payloadBody string) string {
	mac := hmac.New(sha1.New, []byte(secretToken))
	mac.Write([]byte(payloadBody))
	expectedMAC := mac.Sum(nil)
	return "sha1=" + hex.EncodeToString(expectedMAC)
}

func VerifySignature(secretToken, payloadBody string, signatureToCompareWith string) bool {
	signature := GenerateSignature(secretToken, payloadBody)
	return subtle.ConstantTimeCompare([]byte(signature), []byte(signatureToCompareWith)) == 1
}
