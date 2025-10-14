package main

import (
	"fmt"
	"os"

	"github.com/robfig/cron"
)

func main() {

	// p := Person{
	// 	firstName: "John",
	// 	lastName:  "Doe",
	// 	age:       30,
	// 	address: Address{
	// 		city:    "New York",
	// 		country: "USA",
	// 	},
	// 	PhoneHomeCell: PhoneHomeCell{
	// 		home: "123-456-7890",
	// 		cell: "098-765-4321",
	// 	},
	// }

	// p1 := Person{
	// 	firstName: "Jane",
	// 	age:       25,
	// }

	// p1.address.city = "Los Angeles"
	// p1.address.country = "USA"

	// fmt.Println(p.firstName)
	// fmt.Println(p1.firstName)

	// fmt.Println(p.lastName)
	// fmt.Println(p1.lastName)

	// fmt.Println(p.fullName())

	// fmt.Println(p.cell)
	// fmt.Println(p.home)

	// fmt.Println(p == p1)

	// user := struct {
	// 	username string
	// 	email    string
	// }{
	// 	username: "user123",
	// 	email:    "pseudoemail@example.com",
	// }

	// fmt.Println(user.username)

	// fmt.Println("Before incrementing age:", p1.age)
	// p1.incrementAgeBy1()
	// fmt.Println("After incrementing age:", p1.age)

	// fmt.Println(p.address.city)
	// fmt.Println(p1.address.city)

	// fmt.Println(p.address.country)
	// fmt.Println(p1.address.country)

	c := cron.New()

	err := c.AddFunc("@every 0h1m", printHello)
	if err != nil {
		os.Exit(1)
	}
	c.Start()

	select {}

}

func printHello() {
	fmt.Println("Hello")
}

// type Person struct {
// 	firstName string
// 	lastName  string
// 	age       int
// 	address   Address
// 	PhoneHomeCell
// }

// type Address struct {
// 	city    string
// 	country string
// }

// type PhoneHomeCell struct {
// 	home string
// 	cell string
// }

// func (p Person) fullName() string {
// 	return p.firstName + " " + p.lastName
// }

// func (p *Person) incrementAgeBy1() {
// 	p.age++
// }
