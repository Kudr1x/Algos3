package main

import (
	"Algos3/src/AES"
	"fmt"
)

func main() {
	testAES()
}

func testAES() {
	key, err := AES.GenerateRandomKey()
	if err != nil {
		fmt.Println("Error generating key:", err)
		return
	}

	plaintext := []byte("122121112122112112121212")
	ciphertext, err := AES.Encrypt(plaintext, key)
	if err != nil {
		fmt.Println("Error encrypting:", err)
		return
	}

	decrypted, err := AES.Decrypt(ciphertext, key)
	if err != nil {
		fmt.Println("Error decrypting:", err)
		return
	}

	fmt.Printf("Plaintext: %s\n", plaintext)
	fmt.Printf("Ciphertext: %x\n", ciphertext)
	fmt.Printf("Decrypted: %s\n", decrypted)
	fmt.Printf("key: %x\n", key)
}
