package main

import "fmt"

func main() {
	// for
	// manual
	var s = []int{0, 10, -3, 5, 99}
	for index := 0; index < len(s); index++ {
		fmt.Println(s[index])
	}
	// range
	for index := range s {
		fmt.Println(index)
	}

	for index, element := range s {
		fmt.Println(element, index)
	}

	for _, element := range s {
		fmt.Println(element)
	}

	for index := range s {
		s[index] = 10
	}
	fmt.Println(s)

	// copy slice
	var sliceToCopy = []int{12, 23, 34}
	var sliceCopied = make([]int, len(s))
	var nCopied = copy(sliceCopied, s)

	sliceCopied[0] = 0
	sliceCopied[1] = 11

	fmt.Println(nCopied)
	fmt.Println(sliceToCopy)
	fmt.Println(sliceCopied)
	// copy just copy to available space
	var sliceToCopy2 = []int{12, 23, 34}
	var sliceCopied2 []int
	var nCopied2 = copy(sliceCopied2, sliceToCopy2)
	fmt.Println(nCopied2)
	fmt.Println(sliceCopied2)

	// append slice
	var sliceToAppend = []int{12, 23, 34}
	sliceToAppend = append(s, 45)
	sliceToAppend = append(s, 56, 67)
	fmt.Println(sliceToAppend)

	var s1 = []int{12, 23, 34}
	var s2 = []int{45, 56, 67}
	var s3 = append(s1, s2...)
	fmt.Println(s3)

	var sliceNilToAppend []int
	sliceNilToAppend = append(sliceNilToAppend, 10)
	fmt.Println(sliceNilToAppend)

	var a = []int{1, 2, 4, 3, 6}
	var b = []int{-1, 9, -90}
	var ab = make([]int, 0, len(a)+len(b))
	ab = append(ab, a...)
	ab = append(ab, b...)
	fmt.Println(cap(a), cap(b), cap(ab))

	// append affect to add capacity
	var countries = make([]string, 2000)
	countries = append(countries, "Indonesia")
	fmt.Println("cap:", cap(countries))

	// instead, refer to use for range
	var countries2 = make([]string, 2000)
	for i := range countries2 {
		countries2[i] = "Indonesia"
	}
	fmt.Println("cap:", cap(countries2))
}
