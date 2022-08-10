package main

import "fmt"

func main() {
	var firstNum int = 1
	fmt.Println(firstNum)

	var secondNum = 20
	fmt.Println(secondNum)

	thirdNum := 30
	fmt.Println(thirdNum)

	fourthNum := 30
	fourthNum = 31

	fourthNum, fifthNum := 32, 33

	fmt.Println(fourthNum, fifthNum)

	// Declaring multiple variables
	// var python, java, kotlin bool

	// var (
	//	Go         bool
	//	Dart       bool
	//	Rust       bool
	//	TypeScript bool
	// )

	const pi = 3.14
	const hubble int = 77

	const (
		hello = "Hello"
		e     = 2.718
	)

	fmt.Println(pi, hubble, hello, e)

	// iota for increment number

	// const (
	// 	A = iota
	//	B = iota
	//	C = iota
	// )

	const (
		A = iota
		B
		C
	)

	fmt.Println(A, B, C)

	const (
		Read = 1 << iota
		Write
		Execute
	)

	fmt.Println(Read, Write, Execute)

}
