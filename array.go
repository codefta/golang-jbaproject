package main

import "fmt"

func main() {
	// create fixed-array with default value
	var array [4]int
	fmt.Println(array)

	const s1 = 4
	var a [s1]int
	fmt.Println(a)

	var assignArr = [4]int{10, 2, -2, 42}
	fmt.Println(assignArr)

	var assignHalfArr = [4]int{10, 2}
	fmt.Println(assignHalfArr)

	var assignWithIndex = [4]int{
		0: 10,
		3: 42,
	}
	fmt.Println(assignWithIndex)

	// get value by index
	fmt.Println(assignArr[3])
	assignHalfArr[2] = 3
	fmt.Println(assignHalfArr[2])

	// array assignment
	var sourceArr = [4]int{4, 3, 2, 1}
	var targetArr [4]int

	targetArr = sourceArr
	fmt.Println(targetArr)

	var multiArr [3][4]int
	fmt.Println(multiArr)

	var multiArray = [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(multiArray)

	var multiArrayWithIndex = [3][4]int{
		1: {
			2: 5,
			3: 6,
		},
	}
	fmt.Println(multiArrayWithIndex)
}
