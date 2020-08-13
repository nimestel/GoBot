package main

import "fmt"

func main() {
	var test bool

	text := "Hello world!"
	fmt.Println(text, test)
	fmt.Println(text + "2")

	text = text + "123"
	fmt.Println(text)

	result := isAllowedGroups(2, 10)
	fmt.Println("isAllowedGroups", result)
}

func isAllowedGroups(groupCount int, subscriberCount int) bool {
	residue := subscriberCount % groupCount

	fmt.Println(residue)

	return residue == 0
}
