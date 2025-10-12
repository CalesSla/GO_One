package main

import "fmt"

func main() {
	sequence, total := sum(1, 20, 30, 40)
	fmt.Println("Sequence:", sequence, "Total:", total)

	sequence2, total2 := sum(2, 5, 10)
	fmt.Println("Sequence:", sequence2, "Total:", total2)

	numbers := []int{1, 2, 3, 4, 5}
	sequence3, total3 := sum(3, numbers...)
	fmt.Println("Sequence:", sequence3, "Total:", total3)

}

func sum(sequence int, numbers ...int) (int, int) {
	total := 0
	for _, v := range numbers {
		total += v
	}
	return sequence, total
}
