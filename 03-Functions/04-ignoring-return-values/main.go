package main

import "fmt"

/*
ASSIGNMENT
In this code snippet, we only need our customer's first name. Ignore the last name so that the code compiles.
*/

func main() {
	firstName, _ := getNames()
	fmt.Println("Welcome to Textio,", firstName, "!");
}

// don't edit below this line

func getNames() (string, string) {
	return "John", "Doe"
}
