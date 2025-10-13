package main

import "fmt"

func main() {

	// fmt.Print("Hello")
	// fmt.Print(" World")
	// fmt.Print(12, 456)

	// fmt.Println("Hello")
	// fmt.Println("World")
	// fmt.Println(12, 456)

	// name := "John"
	// age := 25
	// fmt.Printf("Name: %s, Age: %d\n", name, age)
	// fmt.Printf("Binary: %b, Hex: %X\n", age, age)

	// s := fmt.Sprint("Hello", "World!", 123, 456)
	// fmt.Print(s)

	// s = fmt.Sprintln("Hello ", "World!", 123, 456)
	// fmt.Print(s)
	// fmt.Print(s)

	// sf := fmt.Sprintf("Name: %s, Age %d", name, age)
	// fmt.Println(sf)
	// fmt.Println(sf)

	// var name string
	// var age int

	// fmt.Print("Enter your name and age: ")
	// fmt.Scanln(&name, &age)
	// fmt.Printf("Name: %s, Age: %d\n", name, age)

	// fmt.Scanf("%s %d", &name, &age)
	// fmt.Printf("Name: %s, Age: %d\n", name, age)

	var age int = 15

	err := checkAge(&age)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(err)
}

func checkAge(age *int) error {
	if *age < 18 {
		return fmt.Errorf("Age %d is below the allowed limit", *age)
	}

	return nil
}
