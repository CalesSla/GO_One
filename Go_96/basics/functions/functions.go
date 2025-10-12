package main

import "fmt"

func main() {

	// sum := add(3, 4)
	// fmt.Println("Sum:", sum)

	// greet := func() {
	// 	fmt.Println("This is an anonymous function")
	// }

	// greet()

	// operation := add

	// result := operation(5, 6)
	// fmt.Println("Result from operation:", result)

	result := applyOperation(5, 3, add)
	fmt.Println("5 + 3 =", result)

	multiplyBy2 := createMultiplier(2)
	resultMult := multiplyBy2(6)
	fmt.Println("6 * 2 = ", resultMult)
}

func add(a, b int) int {
	return a + b
}

func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
