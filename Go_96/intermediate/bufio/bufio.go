package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {

	reader := bufio.NewReader(strings.NewReader("Hello, bufio package!\n"))

	data := make([]byte, 20)
	n, err := reader.Read(data)
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}
	fmt.Printf("Read %d bytes: %s\n", n, string(data[:n]))
	fmt.Println(string(data))

	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading line:", err)
		return
	}
	fmt.Printf("Read line: %s", line)
}
