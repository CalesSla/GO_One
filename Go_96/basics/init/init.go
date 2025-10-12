package main

import "fmt"

func init() {
	fmt.Println("First init function")
}

func init() {
	fmt.Println("Second init function")
}

func init() {
	fmt.Println("Third init function")
}

func main() {

	fmt.Println("Inside main function")

}
