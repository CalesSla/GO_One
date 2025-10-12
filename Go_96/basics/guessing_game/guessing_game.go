package guessing_game

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	target := random.Intn(100) + 1

	fmt.Println("Welcome to the Guessing Game")
	fmt.Println("I have selected a number between 1 and 100. Can you guess it?")
	fmt.Println("Can you guess the number?")

	var guess int
	for {
		fmt.Println("Enter your guess:")
		fmt.Scanln(&guess)

		if guess == target {
			fmt.Println("Congratulations! You guessed the correct number:", target)
			break
		} else if guess < target {
			fmt.Println("Too low! Try again.")
		} else {
			fmt.Println("Too high! Try again.")
		}
	}

}
