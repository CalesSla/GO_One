package main

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// err := os.Mkdir("subdir", 0755)
	// checkError(err)
	// // defer os.RemoveAll("subdir")

	// os.WriteFile("subdir/file", []byte(""), 0755)

	// err = os.MkdirAll("subdir/parent/child", 0755)
	// checkError(err)

	// checkError(os.MkdirAll("subdir/parent/child1", 0755))
	// checkError(os.MkdirAll("subdir/parent/child2", 0755))
	// checkError(os.MkdirAll("subdir/parent/child3", 0755))

	// os.WriteFile("subdir/parent/file", []byte("some data"), 0755)
	// os.WriteFile("subdir/parent/child1/file1", []byte("some data1"), 0755)

	result, err := os.ReadDir("subdir/parent")
	checkError(err)

	for _, entry := range result {
		fmt.Println(entry.Name(), entry.IsDir(), entry.Type())
	}

	checkError(os.Chdir("subdir/parent/child"))

	result, err = os.ReadDir(".")
	checkError(err)

	fmt.Println("Reading subdir/parent/child directory:")
	for _, entry := range result {
		fmt.Println(entry.Name(), entry.IsDir(), entry.Type())
	}

	checkError(os.Chdir("../../.."))

	dir, err := os.Getwd()
	checkError(err)
	fmt.Println("Current working directory:", dir)

}
