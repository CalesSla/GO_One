package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	City    string   `xml:"city"`
	Email   string   `xml:"email"`
}

func main() {

	person := Person{
		Name:  "John",
		Age:   30,
		City:  "New York",
		Email: "example@email.com",
	}

	xmlData, err := xml.Marshal(person)
	if err != nil {
		log.Fatalln("Error Marshalling data into XML:", err)
	}

	// fmt.Println("XML Data: ", xmlData)
	fmt.Println(string(xmlData))
	fmt.Println()

	xmlData1, err := xml.MarshalIndent(person, "", "  ")
	if err != nil {
		log.Fatalln("Error Marshalling data into XML:", err)
	}

	// fmt.Println("XML Data: ", xmlData1)
	fmt.Println(string(xmlData1))

}
