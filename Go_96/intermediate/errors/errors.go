package main

import (
	"errors"
	"fmt"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	return 1, nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("Empty data")
	}
	return nil
}

func main() {
	// result, err := sqrt(16)
	// if err != nil {
	// 	fmt.Println("Error:", err.Error())
	// 	return
	// } else {
	// 	fmt.Println("Result:", result)
	// }
	// result1, err1 := sqrt(-16)
	// if err1 != nil {
	// 	fmt.Println("Error:", err1.Error())
	// 	return
	// } else {
	// 	fmt.Println("Result:", result1)
	// }

	// data := []byte{}
	// err := process(data)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println("Data processed successfully")

	// err1 := eprocess()
	// if err1 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }

	err := readData()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data read successfully")
}

type myError struct {
	message string
}

func (m *myError) Error() string {
	return fmt.Sprintf("Error: %s", m.message)
}

func eprocess() error {
	return &myError{"Custom error message"}
}

func readConfig() error {
	return errors.New("Config error")
}

func readData() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("readData: %w", err)
	}
	return nil
}
