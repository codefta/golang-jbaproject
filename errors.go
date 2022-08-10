package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Implementation error
func openFile() {
	file, err := os.Open("new_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}

// Implementation with errors.New
func divide(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return num1 / num2, nil
}

func divide2(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("cannot divide %.2f by %.2f", num1, num2)
	}
	return num1 / num2, nil
}

func customError() {
	_, err := divide(2.0, 0)
	if err != nil {
		log.Fatal(err)
	}

	_, err = divide2(3.0, 0)
	if err != nil {
		log.Fatal(err)
	}
}

func openFile2(filename string) error {
	if _, err := ioutil.ReadFile(filename); err != nil {
		return fmt.Errorf("error opening %s: %w", filename, err)
	}

	return nil
}

func main() {
	err := openFile2("new_file.txt")
	if err != nil {
		fmt.Printf("error running main.go %s \n", err.Error())

		unwrappedErr := errors.Unwrap(err)
		fmt.Printf("unwrapped error: %v \n", unwrappedErr)
	}
}
