package main

import "fmt"

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r *Rectangle) Scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

func main() {
	rect := Rectangle{length: 10, width: 9}
	area := rect.Area()
	fmt.Println("Area: ", area)

	rect.Scale(1.5)
	newArea := rect.Area()
	fmt.Println("New Area: ", newArea)

	num := MyInt(-5)
	fmt.Println("Is Positive: ", num.IsPositive())
	num2 := MyInt(10)
	fmt.Println("Is Positive: ", num2.IsPositive())

	fmt.Println(num.welcomeMessage())
}

type MyInt int

func (m MyInt) IsPositive() bool {
	return m > 0
}

func (MyInt) welcomeMessage() string {
	return "Welcome to Go methods!"
}
