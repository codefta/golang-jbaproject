package main

import "fmt"

func main() {
	var s []int
	s = []int{0, 1, 3}
	s = []int{120, 56}
	s = []int{43, 42, 12, 12}
	s[0] = 20
	fmt.Println(s[0])

	// initialize slice
	var slice []int
	slice = make([]int, 6, 9)
	var length = len(slice)
	var capacity = cap(slice)
	fmt.Println("slice", slice)
	fmt.Println("length", length)
	fmt.Println("capacity", capacity)

	// multidimensional slice
	var mds [][]int         // define 2d slice
	mds = make([][]int, 10) // initialize 1st slice
	for i := range mds {
		mds[i] = make([]int, 10) // initialize 2nd slice
	}
	fmt.Println(mds)

	// slice like a pointer
	var sliceSource = []int{12, 23, 34}
	var sn = sliceSource

	sn[0] = 0
	sn[1] = 11
	fmt.Println(sliceSource)
	fmt.Println(sn)
}
