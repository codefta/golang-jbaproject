package main

import "fmt"

func main() {
	var i int
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for i = 0; i < 10; i++ {
		fmt.Println(i)
	}

	var s int
	var t int
	for ; t < 50; t++ {
		if s > 100 {
			break
		}
		if t%2 == 0 {
			continue
		}
		s += t
	}
	fmt.Println(s)

	var n int
	fmt.Scan(&n)

	if n < 1 {
		n = 1
	}
	for i := n; i > 1; i-- {
		n = n * (i - 1)
		fmt.Println(n)
	}

	fmt.Println(n)

	arr := []int{1, -2, 3, -4, 5}

	var result int
	if len(arr) == 0 {
		result = -1
	}

	result = 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > 0 {
			result += arr[i]
		}
	}

	fmt.Println(result)
}
