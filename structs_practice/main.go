package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

type outputtable interface {
	saver
	Display()
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func main() {
	printSomething(1)
	printSomething(3.14)
	printSomething("Hello, Go!")

	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}

	outputData(userNote)

	if err != nil {
		return
	}
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func printSomething(value any) {
	switch value.(type) {
	case int:
		fmt.Println("It's an integer.")
	case float64:
		fmt.Println("It's a float.")
	case string:
		fmt.Println(value)
	default:
		fmt.Println("Unknown type.")
	}
	fmt.Println(value)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving note failed.")
		return err
	}
	fmt.Println("Saving the note succeeded!")
	return nil
}

func getUserInput(prompt string) string {
	fmt.Printf("%v", prompt)
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content
}
