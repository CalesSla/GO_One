package main

import "fmt"

func main() {

	message := "Hello, Go!"
	message1 := "Hello, \tGo!"
	message2 := "Hello \rGo!"
	rawMessage := `Hello\nGo`

	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	fmt.Println("Length of message variable is: ", len(message))

	fmt.Println("The first character in message var is:", message[0])
	fmt.Println("The first character in message var is:", string(message[0]))

	greeting := "Hello "
	name := "Alice"

	fmt.Println(greeting + name)
}
