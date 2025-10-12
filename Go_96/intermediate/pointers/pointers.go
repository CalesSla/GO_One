package main

import "fmt"

func main() {

	var ptr *int
	var ptr2 *int
	var a int = 10

	ptr = &a

	fmt.Println(a)
	fmt.Println(ptr)
	fmt.Println(*ptr)

	if ptr2 == nil {
		fmt.Println("ptr2 is nil")
	}

	modifyValue(ptr)
	fmt.Println(a)
}

func modifyValue(ptr *int) {
	*ptr++
}
