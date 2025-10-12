package main

import "fmt"

func main() {

	sequence := adder()
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())

	sequence2 := adder()
	fmt.Println(sequence2())

}

func adder() func() int {
	i := 0

	fmt.Println("Previous i:", i)
	return func() int {
		i++
		fmt.Println("added 1 to i")
		return i
	}
}
