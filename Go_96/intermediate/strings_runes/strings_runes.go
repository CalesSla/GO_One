package main

import (
	"fmt"
	"unicode/utf8"
)

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

	str1 := "Apple"
	str := "apple"
	str2 := "Banana"
	str3 := "Apple"
	fmt.Println(str1 < str2)
	fmt.Println(str3 < str1)
	fmt.Println(str > str1)

	for _, char := range message {
		// fmt.Printf("Character at index %d is %c\n", i, char)
		fmt.Printf("%v\n", char)
	}

	fmt.Println("Rune count: ", utf8.RuneCountInString(greeting))

	greetingWithName := greeting + name
	fmt.Println(greetingWithName)

	var ch rune = 'a'
	var strCh string = "a"
	fmt.Println(ch)
	fmt.Println(strCh)

	fmt.Printf("%c\n", ch)
	fmt.Printf("%v\n", strCh)

	cstr := string(ch)
	fmt.Println(cstr)

	fmt.Printf("Type of ch is %T\n", ch)
	const NIHONGO = "日本語"
	fmt.Println(NIHONGO)

	jhello := "こんにちは"
	for _, runeValue := range jhello {
		fmt.Printf("%c\n", runeValue)
	}

	// Smile
	smile := '😊'
	fmt.Println(smile)
}
