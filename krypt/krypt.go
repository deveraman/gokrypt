package krypt

import (
	"crypto/aes"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"log"
)


func MD5(s string) {
	h := md5.New()
	_, err := io.WriteString(h, s)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%x\n", h.Sum(nil))
}

func SHA256(s string) {
	h := sha256.New()
	_, err := h.Write([]byte(s))
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%x\n", h.Sum(nil))
}

func SHA512_256(s string) {
	h := sha512.New512_256()
	_, err := h.Write([]byte(s))
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%x\n", h.Sum(nil))
}

func AES(key []byte, s string) string {
	c, err := aes.NewCipher(key)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(c.BlockSize())

	out := make([]byte, len(s))

	c.Encrypt(out, []byte(s))

	return hex.EncodeToString(out)
}