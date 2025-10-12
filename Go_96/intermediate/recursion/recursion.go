package main

import "fmt"

func main() {
	// res := factorial(5)
	// fmt.Println(res)

	// res2 := sumOfDigits(12345)
	// fmt.Println(res2)

	fibs := []int{0, 1}
	fibonacci(8, &fibs)
	fmt.Println(fibs)

}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func sumOfDigits(n int) int {
	if n < 10 {
		return n
	}
	return n%10 + sumOfDigits(n/10)
}

func fibonacci(n int, fibs *[]int) {
	if len(*fibs) >= n {
		return
	}
	nextFib := (*fibs)[len(*fibs)-1] + (*fibs)[len(*fibs)-2]
	*fibs = append(*fibs, nextFib)
	fibonacci(n, fibs)
}
