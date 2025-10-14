package main

import "fmt"

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func main() {
	rect := Rectangle{length: 10, width: 9}
	area := rect.Area()
	fmt.Println("Area: ", area)
}
