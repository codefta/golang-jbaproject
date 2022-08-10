package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var myFirstString string
	myFirstString = "I am programmer"
	fmt.Println(myFirstString)

	// Special Characters
	var iamSpecial string = "Hello\n\t"
	fmt.Println(iamSpecial)
	var iamNotSpecial string = `Hello\n\t`
	fmt.Println(iamNotSpecial)

	// Concatenation
	discover := "hello"
	discover = discover + " there"

	// UTF-8 from the box
	var russian = "Привет, Мир!"
	fmt.Println(russian)
	korean := "안녕하세요 월드입니다!"
	fmt.Println(korean)
	var emoji string = "😤"
	fmt.Println(emoji)

	asciiString := "ABCDE"
	utf8String := "Привет"
	fmt.Println(len(asciiString))
	fmt.Println(len(utf8String))
	fmt.Println(utf8.RuneCountInString(asciiString))
	fmt.Println(utf8.RuneCountInString(utf8String))

	fmt.Println(len(emoji))
	fmt.Println(utf8.RuneCountInString(emoji))
}
