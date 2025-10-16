package main

import (
	"fmt"
	"strconv"
)

func main() {

	str := "Hello World"
	fmt.Println(len(str))

	str1 := "Hello"
	str2 := "World"
	result := str1 + " " + str2
	fmt.Println(result)

	fmt.Println(str[0])
	fmt.Println(str[1:7])

	num := 18
	str3 := strconv.Itoa(num)
	fmt.Println(str3)
}
