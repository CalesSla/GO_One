package main

import "fmt"

const pi = 3.14
const GRAVITY = 9.81

func main() {
	fmt.Println(GRAVITY)

	const days int = 7

	const (
		monday        = 1
		tuesday       = 2
		wednesday     = 3
		thursday  int = 4
	)

	fmt.Println(monday)
}
