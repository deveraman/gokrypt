package krypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/pbkdf2"
)

func GetKey(password string, salt []byte) ([]byte, []byte) {
	if salt == nil {
		salt = make([]byte, 16)
		rand.Read(salt)
	}
	return pbkdf2.Key([]byte(password), salt, 1000, 32, sha256.New), salt
}

func Encrypt(passkey string, txt string) string {
	key, salt := GetKey(passkey, nil)
	iv := make([]byte, 12)
	// http://nvlpubs.nist.gov/nistpubs/Legacy/SP/nistspecialpublication800-38d.pdf
	// Section 8.2
	rand.Read(iv)
	newBlock, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(newBlock)
	encData := gcm.Seal(nil, iv, []byte(txt), nil)
	return hex.EncodeToString(salt) + "-" + hex.EncodeToString(iv) + "-" + hex.EncodeToString(encData)
}