package main

import "fmt"

func canReturnNil() string {
	return ""
}

func modString(a *string) {
	fmt.Println("Inside function: ", a)
	*a = "modified"
}

func main() {
	a := "hello"
	fmt.Println("Before: ", a)

	modString(&a)
	fmt.Println("After: ", a)
}
