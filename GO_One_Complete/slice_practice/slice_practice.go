package main

import (
	"fmt"
)

type product struct {
	title string
	id    int
	price float64
}

func main() {
	hobbies := [3]string{"reading", "traveling", "swimming"}
	fmt.Println(hobbies)

	fmt.Println(hobbies[0])
	second_third := hobbies[1:3]
	fmt.Println(second_third)

	fs1 := hobbies[:2]
	fmt.Println(fs1)

	fs2 := hobbies[0:2]
	fmt.Println(fs2)

	sl := fs1[1:3]
	fmt.Println(sl)

	goals := []string{"Go", "Python"}
	fmt.Println(goals)

	goals[1] = "Java"
	fmt.Println(goals)

	goals = append(goals, "JavaScript")
	fmt.Println(goals)

	products := []product{
		{title: "Laptop", id: 1, price: 999.99},
		{title: "Smartphone", id: 2, price: 499.99},
	}
	fmt.Println(products)

	products = append(products, product{title: "Tablet", id: 3, price: 299.99})
	fmt.Println(products)
}
