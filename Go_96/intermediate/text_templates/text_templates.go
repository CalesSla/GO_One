package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {

	// tmpl, err := template.New("example").Parse("Welcome, {{.name}}! How are you doing?\n")
	// tmpl := template.Must(template.New("example").Parse("Welcome, {{.name}}! How are you doing?\n"))

	// data := map[string]any{
	// 	"name": "John",
	// }

	// err := tmpl.Execute(os.Stdout, data)
	// if err != nil {
	// 	panic(err)
	// }

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	templates := map[string]string{
		"welcome":      "Welcome, {{.name}}! We're glad you joined.\n",
		"notification": "{{.name}}, you have a new notification: {{.notification}}\n",
		"error":        "Oops, An error occurred: {{.errorMessage}}\n",
	}

	parsedTemplates := make(map[string]*template.Template)
	for name, tmpl := range templates {
		parsedTemplates[name] = template.Must(template.New(name).Parse(tmpl))
	}

	for {
		// Show menu
		fmt.Println("\nMenu:")
		fmt.Println("1. Join")
		fmt.Println("2. Get Notifications")
		fmt.Println("3. Get Error")
		fmt.Println("4. Exit")

		fmt.Print("Choose an option: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string]any
		var tmpl *template.Template

		switch choice {
		case "1":
			tmpl = parsedTemplates["welcome"]
			data = map[string]any{
				"name": name,
			}
		case "2":
			fmt.Println("Enter your notification message: ")
			notification, _ := reader.ReadString('\n')
			tmpl = parsedTemplates["notification"]
			data = map[string]any{
				"name":         name,
				"notification": strings.TrimSpace(notification),
			}
		case "3":
			fmt.Println("Enter your error message: ")
			errorMessage, _ := reader.ReadString('\n')
			errorMessage = strings.TrimSpace(errorMessage)
			tmpl = parsedTemplates["error"]
			data = map[string]any{
				"name":         name,
				"errorMessage": errorMessage,
			}
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
			continue
		}
		if tmpl != nil {
			err := tmpl.Execute(os.Stdout, data)
			if err != nil {
				fmt.Println("Error executing template:", err)
			}
		}
	}

}
