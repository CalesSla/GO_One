package main

import "fmt"

func main() {

	// process(10)

	process(-3)
}

func process(input int) {

	defer fmt.Println("Deferred: This will run first")
	defer fmt.Println("Deferred: This will run second")
	defer fmt.Println("Deferred: This will run third")

	if input < 0 {
		panic("Input must be a non-negative number")
	}
	fmt.Println("Processing input: ", input)

}
