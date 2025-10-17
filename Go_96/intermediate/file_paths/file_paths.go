package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	relativePath := "./data/file.txt"
	absolutePath := "/home/user/docs/file.txt"

	joinedPath := filepath.Join("home", "Documents", "downloads", "file.zip")
	fmt.Println(joinedPath)

	normalizedPath := filepath.Clean("./data/../data/file.txt")
	fmt.Println(normalizedPath)

	dir, file := filepath.Split("/home/user/docs/file.txt")
	fmt.Println("Directory:", dir)
	fmt.Println("File:", file)
	fmt.Println("Base:", filepath.Base("/home/user/docs/file.txt"))

	fmt.Println(filepath.IsAbs(relativePath))
	fmt.Println(filepath.IsAbs(absolutePath))

	fmt.Println(filepath.Ext(file))
	fmt.Println(strings.TrimSuffix(file, filepath.Ext(file)))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		fmt.Println("Error calculating relative path:", err)
		return
	}
	fmt.Println("Relative path:", rel)

	rel1, err := filepath.Rel("a/c", "a/b/t/file")
	if err != nil {
		fmt.Println("Error calculating relative path:", err)
		return
	}
	fmt.Println("Relative path:", rel1)

	absPth, err := filepath.Abs(relativePath)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
		return
	} else {
		fmt.Println("Absolute path:", absPth)
	}

}
