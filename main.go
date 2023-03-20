package main

import "fmt"

func Cat() string {
	word := "Miao"
	return fmt.Sprintf("%s~ ~ ~ ~", word)
	// return "Wang ~ ~"
}

func main() {
	saying := Cat()
	fmt.Print(saying)
}
