package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// val := rand.New(rand.NewSource(time.Now().Unix()))

	// // fmt.Println(rand.Intn(100) + 1)
	// fmt.Println(val.Intn(101))

	// fmt.Println(rand.Float64())
	fmt.Println("Welcome to the Dice Game!")
	for {

		fmt.Println("1. Roll Dice")
		fmt.Println("2. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("Invalid input. Please enter 1 or 2.")
			continue
		}
		if choice == 2 {
			fmt.Println("Thank you for playing. Goodbye!")
			break
		}

		die1 := rand.Intn(6) + 1
		die2 := rand.Intn(6) + 1

		fmt.Printf("You rolled a %d and a %d\n", die1, die2)
		fmt.Printf("Total: %d\n", die1+die2)

		fmt.Print("Do you want to roll again? (y/n): ")

		for {
			var rollAgain string
			_, err = fmt.Scan(&rollAgain)
			if err != nil || (rollAgain != "y" && rollAgain != "n") {
				fmt.Println("Invalid input. Please enter 'y' or 'n'.")
				continue
			}

			if rollAgain == "n" {
				fmt.Println("Thank you for playing. Goodbye!")
				return
			} else {
				break
			}
		}
	}
}
