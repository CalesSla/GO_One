package functionsarevalues

import (
	"fmt"
)

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	moreNumbers := []int{5, 6, 7, 8, 9}
	fmt.Println(numbers)

	doubled := transformNumbers(&numbers, double)
	fmt.Println(doubled)

	tripled := transformNumbers(&numbers, triple)
	fmt.Println(tripled)

	transformFn1 := getTrasformerFunction(&numbers)
	transformFn2 := getTrasformerFunction(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transformFn1)
	moreTransformedNumbers := transformNumbers(&moreNumbers, transformFn2)
	fmt.Println(transformedNumbers)
	fmt.Println(moreTransformedNumbers)

}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func getTrasformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
