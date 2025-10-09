package arithmetic_operations

import "fmt"

func main() {
	var a, b int = 10, 20
	var result int

	result = a - b
	fmt.Printf("Subtraction: %d - %d = %d\n", a, b, result)

	result = a + b
	fmt.Printf("Addition: %d + %d = %d\n", a, b, result)

	result = a * b
	fmt.Printf("Multiplication: %d * %d = %d\n", a, b, result)

	result = b / a
	fmt.Printf("Division: %d / %d = %d\n", b, a, result)

	result = a % b
	fmt.Printf("Modulus: %d %% %d = %d\n", a, b, result)

	const p float64 = 22 / 7.0
	fmt.Printf("Constant value of p: %f\n", p)

	var maxInt int64 = 232147483647283483647
	fmt.Printf("Max Int64 value: %d\n", maxInt)
}
