package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"
)

func readFromReader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf[:n]))
}

func writeToWriter(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Data written successfully.")
}

func closeResource(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func bufferExample() {
	var buf bytes.Buffer
	buf.WriteString("Hello Buffer!")
	fmt.Println(buf.String())
}

func multiReaderExample() {
	r1 := strings.NewReader("Hello")
	r2 := strings.NewReader(" World!")
	mr := io.MultiReader(r1, r2)
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(mr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())
}

func main() {
	fmt.Println("=== Read from Reader ===")
	readFromReader(strings.NewReader("Hello reader!"))

	fmt.Println("\n=== Write to Writer ===")
	var writer bytes.Buffer
	writeToWriter(&writer, "Hello Writer!")
	fmt.Println(writer.String())

	fmt.Println("\n=== Buffer Example ===")
	bufferExample()

	fmt.Println("\n=== MultiReader Example ===")
	multiReaderExample()
}
