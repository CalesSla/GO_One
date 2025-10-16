package main

import (
	"fmt"
	"strconv"
)

func main() {

	numStr := "12345"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Converted number:", num)
	fmt.Println("Incremented number:", num+1)

	numistr, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Parsed int64 number:", numistr)
	fmt.Println("Incremented int64 number:", numistr+1)

	floatStr := "3.14"
	floatval, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Parsed float64 number:", floatval)
	fmt.Println("Incremented float64 number:", floatval+1)
	fmt.Printf("%.2f\n", floatval)

	binaryStr := "1010"
	decimal, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Binary to decimal:", decimal)

	hexStr := "FF"
	decimalHex, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Hex to decimal:", decimalHex)

	invalidNumStr := "12A45"
	invalidNum, err := strconv.Atoi(invalidNumStr)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Converted number:", invalidNum)
	}

}
