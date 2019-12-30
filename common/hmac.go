package common

import (
	"crypto/hmac"
	"crypto/sha256"
	"io"
)

// HmacGenerate : generate hmac
func HmacGenerate(clientID, clientKey string) []byte {
	hash := hmac.New(sha256.New, []byte(clientKey))
	io.WriteString(hash, clientID)

	return hash.Sum(nil)
}

// HmacValidate : validate hmac
func HmacValidate(clientID string, clientHmac []byte) bool {
	// TODO : get client key from env
	clientKey := "XW2inbzgaFSsdCycV9Ht6hfv4GmU0oZu"
	hash := HmacGenerate(clientID, clientKey)

	return hmac.Equal(clientHmac, hash)
}
