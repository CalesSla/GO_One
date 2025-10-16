package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	str := "Hello World"
	fmt.Println(len(str))

	str1 := "Hello"
	str2 := "World"
	result := str1 + " " + str2
	fmt.Println(result)

	fmt.Println(str[0])
	fmt.Println(str[1:7])

	num := 18
	str3 := strconv.Itoa(num)
	fmt.Println(str3)

	fruits := "apple, orage, banana"
	parts := strings.Split(fruits, ",")
	fmt.Println(parts)

	fmt.Printf("Type of parts: %T\n", parts)

	countries := []string{"Germany", "France", "Italy"}
	joined := strings.Join(countries, ", ")
	fmt.Println(joined)

	fmt.Println(strings.Contains(str, "World"))

	replaced := strings.Replace(str, "World", "Go", 2)
	fmt.Println(replaced)

	strwspace := "   Hello World   "
	fmt.Println("Before Trim:", strwspace)
	fmt.Println("After Trim:", strings.TrimSpace(strwspace))

	fmt.Println(strings.ToLower(strwspace))
	fmt.Println(strings.ToUpper(strwspace))

	fmt.Println(strings.Repeat("foo", 3))

	fmt.Println(strings.Count(str, "l"))

	fmt.Println(strings.HasPrefix("Hello", "he"))
	fmt.Println(strings.HasSuffix("Hello", "lo"))

	str5 := "Hello, 123 Go 11!"
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(str5, -1)
	fmt.Println(matches)

}
