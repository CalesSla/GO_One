package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

// func main() {
// 	var productNames [4]string = [4]string{"Laptop"}
// 	prices := [4]float64{19.99, 29.99, 39.99, 49.99}
// 	fmt.Println(prices)

// 	productNames[2] = "Desktop"

// 	fmt.Println(productNames)

// 	fmt.Println(prices[2])

// 	featuredPrices := prices[1:]
// 	featuredPrices[0] = 100.00

// 	highlightedPrices := featuredPrices[:1]
// 	fmt.Println(featuredPrices)
// 	fmt.Println(prices)
// 	fmt.Println(len(featuredPrices), cap(featuredPrices))
// 	fmt.Println(highlightedPrices)

// 	highlightedPrices = highlightedPrices[:3]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// }

func main() {
	prices := []float64{19.99, 29.99}
	fmt.Println(prices[0:1])
}
