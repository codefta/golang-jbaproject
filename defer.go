package main

import (
	"fmt"
)

func main() {
	
}

func scopedDefer() {
	n := 0
	defer func() { fmt.Println("n =", n, "- first deferred print") }()
	{
		defer func() { fmt.Println("n =", n, "- second deferred print") }()
		n++
	}
	n++
}

func greeting() {
	defer fmt.Println("Printed after Hello, JB Academy!")
	fmt.Println("Hello, JB Academy!")
}

func deferMultipleFunc() {
	defer fmt.Println("Printed after the deferMultipleFunc() function is completed.")

	greeting()

	fmt.Println("Printed after calling the greeting() function.")
}

func multipleDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func singleDefer() {
	defer fmt.Println("Printed second! ⓶")
	fmt.Println("Printed first! 1️⃣")
}
