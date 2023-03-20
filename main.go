package main

import "fmt"

func Cat() string {
	var password = "123456"
	word := "Miao"
	return fmt.Sprintf("%s~ ~ ~ ~", word)
	// return "Wang ~ ~"
}

func main() {
	saying := Cat()
	fmt.Print(saying)
}
