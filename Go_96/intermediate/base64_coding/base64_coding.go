package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	data := []byte("He~lo, base64 encoding")
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error decoding base64 string:", err)
		return
	}
	fmt.Println(string(decoded))

	urlSafeEncoded := base64.URLEncoding.EncodeToString(data)
	fmt.Println("URL Safe Encoded:", urlSafeEncoded)

}
