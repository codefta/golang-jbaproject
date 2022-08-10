package main

import "fmt"

func main() {
	var s = "some string variable"
	var p = &s // referencing to pointer s
	fmt.Println(s)
	fmt.Println(p) // just print address of reference variable

	// create pointer variable
	var a *string
	var b = "my string"
	a = &b
	fmt.Println(a)

	// other way to create pointer
	var c = new(string)
	var d = "my string 2"
	*c = d
	fmt.Println(*c) // print value of pointer

	var e **string
	e = new(*string)
	*e = new(string)
	fmt.Println(**e)
	**e = "is this even possible?"
	fmt.Println(**e)

	var f = "yes, it is possible"
	var p1 = &f
	var p2 = &p1
	fmt.Println(*p2 == p1)
	fmt.Println(**p2 == f)
}
