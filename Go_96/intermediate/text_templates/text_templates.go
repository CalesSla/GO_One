package main

import (
	"os"
	"text/template"
)

func main() {

	// tmpl, err := template.New("example").Parse("Welcome, {{.name}}! How are you doing?\n")
	tmpl := template.Must(template.New("example").Parse("Welcome, {{.name}}! How are you doing?\n"))

	data := map[string]any{
		"name": "John",
	}

	err := tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}
