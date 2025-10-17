package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	data := []byte("Hello World!\n")
	numBites, err := file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Printf("Wrote %d bytes to file.\n", numBites)

	file, err = os.Create("write_string.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	numBites, err = file.WriteString("Hello Go!\n")
	if err != nil {
		fmt.Println("Error writing string to file:", err)
		return
	}
	fmt.Printf("Wrote %d bytes to file.\n", numBites)

}
