package main

import (
	"errors"
	"fmt"
	"os"
)

const dataFile = "data.txt"

func writeToFile(ebt, profit, ratio float64) {
	data := fmt.Sprintf("%v,%v,%v", ebt, profit, ratio)
	fmt.Println("Writing to file:", data)
	err := os.WriteFile(dataFile, []byte(data), 0644)
	if err != nil {
		panic("Unable to write data to file.")
	}
}

func readFromFile(filePath string) (revenue, expenses, taxRate float64) {
	fmt.Println("Reading from file...")
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic("Unable to read data from file.")
	}
	_, err = fmt.Sscanf(string(content), "%f,%f,%f", &revenue, &expenses, &taxRate)
	if err != nil {
		panic("Unable to parse data from file.")
	}

	fmt.Println("Revenue: ", revenue)
	fmt.Println("Expenses: ", expenses)
	fmt.Println("Tax Rate: ", taxRate)

	return revenue, expenses, taxRate
}

func main() {
	var revenue, expenses, taxRate float64

	_, err := os.Stat(dataFile)
	if os.IsNotExist(err) {
		revenue = getUserInput("Revenue: ")
		expenses = getUserInput("Expenses: ")
		taxRate = getUserInput("Tax Rate (%): ")
	} else {
		revenue, expenses, taxRate = readFromFile(dataFile)
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	writeToFile(ebt, profit, ratio)

	printStats(ebt, profit, ratio)

}

func getUserInput(infoText string) (result float64) {

	fmt.Print(infoText)
	fmt.Scan(&result)
	err := validateInput(result)
	if err != nil {
		return getUserInput(infoText)
	}
	return result
}

func validateInput(input float64) (err error) {
	if input <= 0 {
		err = errors.New("Invalid amount. Please enter a positive number.\n")
		fmt.Print(err)
		return err
	}
	return nil
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
