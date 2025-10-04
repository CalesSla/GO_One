package main

import "fmt"

func main() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate (%): ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	printStats(ebt, profit, ratio)

}

func getUserInput(infoText string) (result float64) {

	fmt.Print(infoText)
	fmt.Scan(&result)

	return result
}

func calculateFinancials(
	revenue float64,
	expenses float64,
	taxRate float64,
) (
	ebt float64,
	profit float64,
	ratio float64,
) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = profit / revenue

	return ebt, profit, ratio
}

func printStats(
	ebt float64,
	profit float64,
	ratio float64,
) {
	fmt.Printf("EBT: %.1f\n", ebt)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio: %.1f%%\n", ratio)
}
