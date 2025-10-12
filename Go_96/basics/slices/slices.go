package main

import (
	"fmt"
	"slices"
)

func main() {

	// var numbers []int
	// fmt.Println(numbers)

	// var numbers1 = []int{1, 2, 3}
	// fmt.Println(numbers1)

	// numbers2 := []int{4, 5, 6}
	// fmt.Println(numbers2)

	// slice := make([]int, 5)
	// fmt.Println(slice)

	a := [5]int{1, 2, 3, 4, 5}
	slice := a[1:4]
	fmt.Println(slice)

	slice = append(slice, 6, 7)
	fmt.Println(slice)

	sliceCopy := make([]int, len(slice))
	fmt.Println("Before copy, sliceCopy:", sliceCopy)
	copy(sliceCopy, slice)
	fmt.Println("After copy, sliceCopy:", sliceCopy)

	var nilSlice []int
	fmt.Println("nilSlice:", nilSlice)

	for i, v := range slice {
		fmt.Println("Index:", i, "Value:", v)
	}

	fmt.Println("element at index 3:", slice[3])

	// slice[3] = 50
	// fmt.Println("Element at index 3 after modification:", slice[3])

	if slices.Equal(slice, sliceCopy) {
		fmt.Println("The slices are equal.")
	} else {
		fmt.Println("The slices are not equal.")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2D Slice:", twoD)

	slice2 := slice[1:4]
	fmt.Println("slice2:", slice2)

	fmt.Println("Length of slice2:", len(slice2))
	fmt.Println("Capacity of slice2:", cap(slice2))
	fmt.Println("Length of slice:", len(slice))
	fmt.Println("Capacity of slice:", cap(slice))
}
