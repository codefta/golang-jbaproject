package main

import "fmt"

func main() {
	// Printf
	fmt.Printf("kitty\nkitty2\n")
	s := "Golang"
	a := "Young"
	// string %s
	fmt.Printf("We are %s hackers, we are so %s", s, a)
	// string with width
	fmt.Printf("|%8s|", "pikachu")
	// raw string
	fmt.Printf("%%")
	// integer %d
	fmt.Printf("%d", 36)
	// floating point %f
	fmt.Printf("%f", 2.1)
	fmt.Printf("%.1f", 20.33) // floating point with precision
	// rune %c
	fmt.Printf("%c", 'c')
	// bool %t
	fmt.Printf("%t", true)

	// ordering
	x := "First Variable"
	y := "Second Variable"
	fmt.Printf("%[1]s | %[1]s \n", x)
	fmt.Printf("%[2]s | %[1]s \n", x, y)

	// unicode %v
	fmt.Printf("%v", '🪲')
	fmt.Printf("%v", []int{1, 2, 3, 4, 5, 6})
	fmt.Printf("%v", [3]int{3, 9, 8})
	fmt.Printf("%v", 1+5i)

	// Sprintf
	first := fmt.Sprintf("%d", 500) // variable of first has value 500
	binaryValue := fmt.Sprintf("%b", 500)
	fmt.Println(first, binaryValue)

	// Sprintln
	stringExample := fmt.Sprintln("if you try to print s, ",
		"it will automatically print a new line at the end of the string")
	fmt.Print(stringExample, "NEWLINE\n")

	// Multiline string
	fmt.Printf(`%s
	and
	the
	brave
	new\n
	world\n`, "Go")
}
