package main

import (
	"fmt"
)

func main() {
	age := 32

	var agePointer *int
	agePointer = &age

	fmt.Println("Age: ", *agePointer)

	getAdultYears(agePointer)
	fmt.Println("Adult Years: ", age)
}

func getAdultYears(age *int) int {
	*age = *age - 18
}
