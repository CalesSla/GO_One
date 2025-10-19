package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age,omitempty"`
	Email   string   `xml:"email"`
	Address Address  `xml:"address"`
}

type Address struct {
	City  string `xml:"city"`
	State string `xml:"state"`
}

func main() {

	person := Person{
		Name: "John",
		// Age:   30,
		// City:  "New York",
		Email: "example@email.com",
		Address: Address{
			City:  "New York",
			State: "NY",
		},
	}

	// xmlData, err := xml.Marshal(person)
	// if err != nil {
	// 	log.Fatalln("Error Marshalling data into XML:", err)
	// }

	// // fmt.Println("XML Data: ", xmlData)
	// fmt.Println(string(xmlData))
	// fmt.Println()

	xmlData1, err := xml.MarshalIndent(person, "", "  ")
	if err != nil {
		log.Fatalln("Error Marshalling data into XML:", err)
	}

	// fmt.Println("XML Data: ", xmlData1)
	fmt.Println(string(xmlData1))

	// xmlRaw := `
	// <person>
	//   <name>Jane</name>
	//   <age>25</age>
	//   <city>Los Angeles</city>
	//   <email>example@email.com</email>
	// </person>
	// `
	xmlRaw := `
	<person>
		<name>Jane</name>
		<age>25</age>
		<email>example@mail.com</email>
		<address>
			<city>Los Angeles</city>
			<state>CA</state>
		</address>
	</person>
	`

	var personxml Person
	err = xml.Unmarshal([]byte(xmlRaw), &personxml)
	if err != nil {
		log.Fatalln("Error Unmarshalling XML data:", err)
	}

	fmt.Println(personxml)
	fmt.Println("Local String: ", personxml.XMLName.Local)
	fmt.Println("Namespace: ", personxml.XMLName.Space)

	book := Book{
		ISBN:       "123-456-789",
		Title:      "Go Programming",
		Author:     "John Doe",
		Pseudo:     "Gopher",
		PseudoAttr: "GoLang",
	}

	fmt.Println()

	xmlDataAttr, err := xml.MarshalIndent(book, "", "  ")
	if err != nil {
		log.Fatalln("Error Marshalling book data into XML:", err)
	}

	fmt.Println(string(xmlDataAttr))

}

type Book struct {
	XMLName    xml.Name `xml:"book"`
	ISBN       string   `xml:"isbn,attr"`
	Title      string   `xml:"title,attr"`
	Author     string   `xml:"author,attr"`
	Pseudo     string   `xml:"pseudo"`
	PseudoAttr string   `xml:"pseudo_attr,attr"`
}
