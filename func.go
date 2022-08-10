package main

import "fmt"

// using parameter
func hello(message string) string {
	return "Hello " + message + "!"
}

// using variadic parameter
func multiply(nums ...int) int {
	total := 1
	for _, num := range nums {
		total *= num
	}
	return total
}

// using pointer parameter
func realSwap(x *int, y *int) {
	*x, *y = *y, *x
}

func fakeSwap(x, y int) {
	x, y = y, x
}

func main() {
	greeting := hello("Fathi")
	fmt.Println(greeting)
	fmt.Println(hello("Azrania"))

	fmt.Println(multiply())
	fmt.Println(multiply(1, 2, 3))

	var num1 = 42
	var num2 = 27
	fmt.Println("Before swapping values: x = ", num1, "and y =", num2)
	realSwap(&num1, &num2)
	fmt.Println("After swapping values: x = ", num1, "and y =", num2)

	fmt.Println("Before swapping values: x = ", num1, "and y =", num2)
	fakeSwap(num1, num2)
	fmt.Println("After swapping values: x = ", num1, "and y =", num2)
}
