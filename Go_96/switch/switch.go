package main

import "fmt"

func main() {
	// fruit := "pineapple"
	// switch fruit {
	// case "apple":
	// 	fmt.Println("It's an apple")
	// case "banana":
	// 	fmt.Println("It's a banana")
	// case "orange":
	// 	fmt.Println("It's an orange")
	// default:
	// 	fmt.Println("Unknown fruit")
	// }

	// day := "a"
	// switch day {
	// case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
	// 	fmt.Println("It's a weekday")
	// case "Saturday", "Sunday":
	// 	fmt.Println("It's the weekend")
	// default:
	// 	fmt.Println("Unknown day")
	// }

	// number := 50
	// switch {
	// case number < 10:
	// 	fmt.Println("The number is less than 10")
	// case number >= 10 && number < 20:
	// 	fmt.Println("The number is between 10 and 19")
	// default:
	// 	fmt.Println("The number is 20 or greater")
	// }

	// num := 2

	// switch {
	// case num > 1:
	// 	fmt.Println("Greater than 1")
	// 	fallthrough
	// case num == 2:
	// 	fmt.Println("Equal to 2")
	// default:
	// 	fmt.Println("Now two")
	// }
	checkType(42)
	checkType("hello")
	checkType(true)
	checkType(3.14)
}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("It's an integer")
	case string:
		fmt.Println("It's a string")
	case bool:
		fmt.Println("It's a boolean")
	default:
		fmt.Println("Unknown type")
	}
}
