package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"log"
)

//go:embed example.txt
var content string

//go:embed custom_error
var custom_error embed.FS

func main() {

	fmt.Println("Embedded Content: ", content)
	content, err := custom_error.ReadFile("custom_error/hello.txt")
	if err != nil {
		fmt.Println("Error reading embedded file:", err)
		return
	}
	fmt.Println("Content from embedded file:", string(content))

	err = fs.WalkDir(custom_error, "custom_error", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		fmt.Println(path)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

}
