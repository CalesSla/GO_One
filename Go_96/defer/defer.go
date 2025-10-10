package main

import "fmt"

func main() {

	process(10)
}

func process(i int) {
	defer fmt.Println("Deferred i value: ", i)
	defer fmt.Println("Deferred: This will run first")
	defer fmt.Println("Deferred: This will run second")
	defer fmt.Println("Deferred: This will run third")
	fmt.Println("This will run first")
	i++
	fmt.Println("Value of i:", i)
	// exit with an error
	panic("An error occurred!")
	fmt.Println("This will run last")
}
