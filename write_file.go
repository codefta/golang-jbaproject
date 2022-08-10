package main

import (
	"log"
	"os"
)

func main() {
	deleteFolder()
}

func deleteFolder() {
	err := os.RemoveAll("test_directory")
	if err != nil {
		log.Fatal(err)
	}
}

func deleteFile() {
	err := os.Remove("test.txt")
	if err != nil {
		log.Fatal(err)
	}
}

func createFile() {
	_, err := os.Create("new_file.txt")
	if err != nil {
		log.Fatal(err)
	}

	// another way
	_, err = os.OpenFile("new_file.txt", os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func openingWithOpenFile() {
	_, err := os.OpenFile("test_file.txt", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	_, err = os.OpenFile("test_file.txt", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func openingFile() {
	file, err := os.Open("test_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}
