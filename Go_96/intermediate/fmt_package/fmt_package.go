package main

import "fmt"

func main() {

	fmt.Print("Hello")
	fmt.Print(" World")
	fmt.Print(12, 456)

	name := "John"
	age := 25
	fmt.Printf("Name: %s, Age: %d\n", name, age)
	fmt.Printf("Binary: %b, Hex: %X\n", age, age)

}
