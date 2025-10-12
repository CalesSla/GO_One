package main

import "fmt"

func main() {

	// var numbers [5]int
	// fmt.Println(numbers)

	// numbers[4] = 20
	// fmt.Println(numbers)

	// fruits := [4]string{"apple", "banana", "orange", "grape"}
	// fmt.Println(fruits)

	// originalArray := [3]int{1, 2, 3}
	// copiedArray := originalArray

	// copiedArray[0] = 100
	// fmt.Println("Original Array:", originalArray)
	// fmt.Println("Copied Array:", copiedArray)

	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println("Element at index, ", i, ":", numbers[i])
	// }

	// for _, v := range numbers {
	// 	fmt.Printf("Value: %d\n", v)
	// }

	// a, _ := someFunction()
	// fmt.Println(a)

	// fmt.Println("Length of numbers array:", len(numbers))

	// array1 := [3]int{1, 2, 3}
	// array2 := [3]int{1, 2, 3}
	// array3 := [3]int{4, 5, 6}

	// fmt.Println("array1 == array2:", array1 == array2)
	// fmt.Println("array1 == array3:", array1 == array3)

	// var matrix [3][3]int = [3][3]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }
	// fmt.Println("Matrix:", matrix)

	originalArray := [3]int{1, 2, 3}

	copiedArray := &originalArray
	copiedArray[0] = 100
	fmt.Println("Original Array:", originalArray)
	fmt.Println("Copied Array:", *copiedArray)
}

func someFunction() (int, int) {
	return 1, 2
}
