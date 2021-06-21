package dkrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"strings"

	"github.com/ramanverma2k/gokrypt/krypt"
)

func Dkrypt(passkey string, encString string) string {
	arr := strings.Split(encString, "-")
	salt, _ := hex.DecodeString(arr[0])
	iv, _ := hex.DecodeString(arr[1])
	dkrypt, _ := hex.DecodeString(arr[2])
	key, _ := krypt.GetKey(passkey, salt)
	newBlock, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(newBlock)
	dkrypt, _ = gcm.Open(nil, iv, dkrypt, nil)
	return string(dkrypt)
}