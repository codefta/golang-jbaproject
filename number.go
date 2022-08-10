package main

import "fmt"

func main() {
	// Integer
	var number uint8 = 100 // int
	number2 := 100         // int
	fmt.Println(number)
	fmt.Println(number2)

	// floating point
	var floatNumber float32 = 3.4 // float
	fmt.Println(floatNumber)

	// boolean
	var b bool = true
	fmt.Println(b)
	var b1 = true
	b2 := !b1
	fmt.Println(b2)
	b3 := b2 || b1
	fmt.Println(b3)
	b4 := b2 && b3
	fmt.Println(b4)
	b5 := b1 && b3
	fmt.Println(b5)

	// rune
	var runeNumber rune = 5 // int32, also using for UTF-8
	fmt.Println(runeNumber)
	// byte
	var byteNumber byte = 3 // uint8, also using for ASCII
	fmt.Println(byteNumber)

}
