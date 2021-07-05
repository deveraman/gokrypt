package main

import (
	"fmt"

	"github.com/ramanverma2k/gokrypt/dkrypt"
	"github.com/ramanverma2k/gokrypt/krypt"
)

func main() {
	t0 := time.Now()

	file, err := os.Open("10-million-password-list-top-1000000.txt")
	if err != nil {
		log.Fatalf("Error:  %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := 0

	for scanner.Scan() {
		c := krypt.Encrypt(""this is a random key"", scanner.Text())
		fmt.Println(c)
		lines++
	}

	fmt.Printf("Encrypted %v passwords\n", lines)
	fmt.Println("It took ", time.Since(t0))
}
