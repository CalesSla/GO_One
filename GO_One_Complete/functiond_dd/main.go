package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := sumup(1, 10, 15, 40, -5)
	anotherSum := sumup(1, numbers...)
	fmt.Println("Sum: ", sum)
	fmt.Println("Another Sum: ", anotherSum)
}

func sumup(startingValue int, numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}
