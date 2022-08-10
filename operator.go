package main

import "fmt"

func main() {
	// Addition
	result := 1 + 5
	fmt.Println(result)

	result2 := 1234 + 4321
	fmt.Println(result2)

	result3 := -1000 + 2000
	fmt.Println(result3)

	result4 := 10.1 + 1.234
	fmt.Println(result4)
	result5 := -5.6789 + 19.23
	fmt.Println(result5)

	// unary
	x := 3.3
	y := +x
	fmt.Println(y)

	x = -30
	y = +x
	fmt.Println(y)

	x = 5
	y = -x
	fmt.Println(y)

	x = -20
	y = -x
	fmt.Println(y)

	// Multipication, divide, and modulo
	x = 30
	y = 14
	fmt.Println(x / y)
	fmt.Println(float32(x) / float32(y))

	n := 100
	m := 15
	fmt.Println(n % m)

	// Assignment operator
	d := 23
	d += 15
	fmt.Println(d)
	d -= 3
	fmt.Println(d)
	d *= 5
	fmt.Println(d)
	d /= 3
	fmt.Println(d)
	d %= 2
	fmt.Println(d)

	// Increment & decrement
	numExample := 0
	numExample++
	fmt.Println(numExample)
	numExample--
	fmt.Println(numExample)

}
