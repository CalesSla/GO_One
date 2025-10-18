package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName    string  `json:"first_name"`
	Age          int     `json:"age,omitempty"`
	EmailAddress string  `json:"email"`
	Address      Address `json:"address"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func main() {

	person := Person{FirstName: "John", Age: 30}
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println(jsonData)
	fmt.Println(string(jsonData))

	person1 := Person{
		FirstName:    "Jane",
		Age:          30,
		EmailAddress: "jane",
		Address: Address{
			City:  "New York",
			State: "NY",
		},
	}

	jsonData1, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println(string(jsonData1))

	jsonData2 := `{"full_name": "Mike", "age": 25, "email": "example@email.com"}`
	var employeeFromJson Employee
	err = json.Unmarshal([]byte(jsonData2), &employeeFromJson)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Println(employeeFromJson)
	fmt.Println("Mike's age increased by 5: ", employeeFromJson.Age+5)
	fmt.Println("Mike's email is: ", employeeFromJson.Email)

	listOfCityState := []Address{
		{City: "Los Angeles", State: "CA"},
		{City: "Chicago", State: "IL"},
		{City: "Houston", State: "TX"},
	}

	fmt.Println(listOfCityState)
	jsonList, err := json.Marshal(listOfCityState)
	if err != nil {
		fmt.Println("Error marshalling list to JSON:", err)
		return
	}

	fmt.Println(string(jsonList))

	// Handling unknown json structures

	jsonData3 := `{"name": "John", "age": 30, "address": {"city": "New York", "state": "NY"}}`

	var data map[string]any

	json.Unmarshal([]byte(jsonData3), &data)
	if err != nil {
		log.Fatalln("Error unmarshalling JSON:", err)
	}

	fmt.Println("Decoded/Unmarshalled JSON:", data)
	fmt.Println("Decoded/Unmarshalled JSON:", data["address"])
	fmt.Println("Decoded/Unmarshalled JSON:", data["name"])

}

type Employee struct {
	FullName string `json:"full_name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
}
