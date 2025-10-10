package main

import "fmt"

func main() {

	process()
}

func process() {
	defer fmt.Println("Deferred: This will run first")
	defer fmt.Println("Deferred: This will run second")
	defer fmt.Println("Deferred: This will run third")
	fmt.Println("This will run first")
}
