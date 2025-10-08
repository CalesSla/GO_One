package maps

import "fmt"

func main() {
	websites := map[string]string{
		"google":        "https://google.com",
		"facebook":      "https://facebook.com",
		"stackoverflow": "https://stackoverflow.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["google"])
	websites["linkedin"] = "https://linkedin.com"
	fmt.Println(websites)

	delete(websites, "facebook")
	fmt.Println("Latest: ", websites)
}
