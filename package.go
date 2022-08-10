package main

import (
	"fmt"
	. "math" // using . for using package
	"math/rand"
	_ "os" // using _ for unused package
)

func main() {
	fmt.Println("Test package")
	Abs(16)
	rand.Float64()
}
