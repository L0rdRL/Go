package main

import "fmt"

func main() {
	age := 15

	// Basic if
	if age < 25 {
		fmt.Println("You are too young (full)")
	}

	// Short  syntax
	if isChild := isChildren(age); isChild == true {
		fmt.Println("You are very young (short)")
		fmt.Println(isChild)
	}

	// If ...else
	age = 40
	if age < 18 {
		fmt.Println("You are too young")
	} else {
		fmt.Println("You are an adult")
	}

	// &&
	if age >= 7 && age <= 18 {
		fmt.Println("You are in school")
	}

	// ||
	if age == 14 || age == 20 || age == 40 {
		fmt.Println("You have to get new passport")
	}

	// Short suyntax
	if !isChildren(age) {
		fmt.Println("You are an adult")
	}
}

func isChildren(age int) bool {
	return age < 18
}
