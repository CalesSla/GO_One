package main

import (
	"fmt"
	"os"
)

func main() {

	defer fmt.Println("Deferred: This will run before exiting main")

	fmt.Println("Starting the main function")

	os.Exit(1)

	fmt.Println("This will not be printed")

}
