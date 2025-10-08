package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Slava"
	userNames[1] = "Ana"
	fmt.Println(userNames)

	courseRatings := make(floatMap, 3)
	courseRatings["Go"] = 4.5
	courseRatings["Python"] = 4.7
	fmt.Println(courseRatings)

	courseRatings.output()

	for index, value := range userNames {
		fmt.Println(index)
		fmt.Println(value)
	}

	for key, value := range courseRatings {
		fmt.Println(key, value)

	}
}
