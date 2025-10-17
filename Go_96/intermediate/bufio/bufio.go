package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// reader := bufio.NewReader(strings.NewReader("Hello, bufio package!\n"))

	// data := make([]byte, 20)
	// n, err := reader.Read(data)
	// if err != nil {
	// 	fmt.Println("Error reading data:", err)
	// 	return
	// }
	// fmt.Printf("Read %d bytes: %s\n", n, string(data[:n]))
	// fmt.Println(string(data))

	// line, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Error reading line:", err)
	// 	return
	// }
	// fmt.Printf("Read line: %s", line)

	writer := bufio.NewWriter(os.Stdout)

	data := []byte("Hello, bufio package!\n")
	n, err := writer.Write(data)

	if err != nil {
		fmt.Println("Error writing data:", err)
		return
	}
	fmt.Printf("Wrote %d bytes.\n", n)
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing data:", err)
		return
	}

	str := "This is a string.\n"
	n, err = writer.WriteString(str)
	if err != nil {
		fmt.Println("Error writing string:", err)
		return
	}
	fmt.Printf("Wrote %d bytes.\n", n)
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing data:", err)
		return
	}
}
