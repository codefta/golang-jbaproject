package main

import "fmt"

func main() {
	// If - else
	number := 4
	if number%2 == 0 {
		fmt.Println("The number", number, "is even")
	} else {
		fmt.Println("The number", number, "is odd")
	}
	// If-else-if
	var score int = 100
	if score > 90 {
		fmt.Println("Your grade is A.Congratulations!")
	} else if score > 80 {
		fmt.Println("Your grade is B. Good enough.")
	} else if score > 70 {
		fmt.Println("Your grade is C. Could've done better!")
	} else if score > 60 {
		fmt.Println("Your grade is D. Study more next time!")
	} else {
		fmt.Println("Your grade is F. Terrible! you failed the test!")
	}

	// Switch
	var selection int = 2
	switch selection {
	case 1:
		fmt.Println("Starting new game!")
	case 2:
		fmt.Println("Loading a saved game.")
	case 3:
		fmt.Println("Quit the game.")
	default:
		fmt.Println("Invalid selection. Try again.")
	}

	var switchNumber int = 10
	switch {
	case switchNumber%2 == 0:
		fmt.Println("Switch number is even")
	case number%2 != 0:
		fmt.Println("Switch number is odd")
	}
}
