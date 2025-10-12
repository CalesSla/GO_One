package main

func main() {

	age := 25
	if age >= 18 {
		println("You are an adult.")
	}

	temperature := 25
	if temperature > 30 {
		println("It's a hot day.")
	} else {
		println("It's a cold day.")
	}

	score := 90
	if score >= 90 {
		println("Grade: A")
	} else if score >= 80 {
		println("Grade: B")
	} else if score >= 70 {
		println("Grade: C")
	} else {
		println("Grade: F")
	}
}
