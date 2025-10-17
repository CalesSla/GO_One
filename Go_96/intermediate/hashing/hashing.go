package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {

	password := "password123"

	// hash := sha256.Sum256([]byte(password))
	// hash512 := sha512.Sum512([]byte(password))

	// fmt.Println(password)
	// fmt.Println(hash)
	// fmt.Println(hash512)
	// fmt.Printf("SHA-256 Hash hex val: %x\n", hash)
	// fmt.Printf("SHA-512 Hash hex val: %x\n", hash512)

	salt, err := generateSalt()
	fmt.Println("Original Salt: ", salt)
	fmt.Printf("Original Salt hex: %x\n", salt)

	if err != nil {
		fmt.Println("Error generating salt:", err)
		return
	}

	signupHash := hashPassword(password, salt)

	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt (base64):", saltStr)
	fmt.Println("Hashed Password (base64):", signupHash)
	hashOriginalPassword := sha256.Sum256([]byte(password))
	fmt.Println("Hash of just the password without salt: ", base64.StdEncoding.EncodeToString(hashOriginalPassword[:]))

	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Error decoding salt:", err)
		return
	}

	loginPassword := "password123"

	loginHash := hashPassword(loginPassword, decodedSalt)

	if signupHash == loginHash {
		fmt.Println("Password match!")
	} else {
		fmt.Println("Password do not match!")
	}

}

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil

}

func hashPassword(password string, salt []byte) string {

	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])

}
