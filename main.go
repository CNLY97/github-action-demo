package main

import "fmt"

func Cat() string {
	var password = "123456"
	fmt.Print(password)
	word := "Miao"
	return fmt.Sprintf("%s~ ~ ~ ~", word)
	// return "Wang ~ ~"
}

func main() {
	saying := Cat()
	fmt.Print(saying)
}
