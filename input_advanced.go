package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func newReader() {
	reader := bufio.NewReader(os.Stdin)

	b, err := reader.ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	s, err := reader.ReadString('d')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
}

func newScanner() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}

func wordScanner() {
	wordScanner := bufio.NewScanner(os.Stdin)
	wordScanner.Split(bufio.ScanWords)

	for wordScanner.Scan() {
		fmt.Println(wordScanner.Text())
	}
}

func splitBool(data []byte, atEOF bool) (advance int, token []byte, err error) {
	advance, token, err = bufio.ScanWords(data, atEOF)
	if err == nil && token != nil {
		_, err = strconv.ParseBool(string(token))
	}
	return
}

func boolScanner() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(splitBool)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func command() {
	var command string
	for command != "exit" {
		fmt.Print("Enter a command and data: > ")
		reader := bufio.NewReader(os.Stdin)
		var err error
		command, err = reader.ReadString(' ')
		if err != nil {
			log.Fatal(err)
		}
		command = strings.ReplaceAll(command, " ", "")
		if command != "exit" {
			data, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s %s <<< this is replace by \"%s\" >>>\n", command, data, command)
		}

		fmt.Println("[Info] Bye!")
	}
}

func splitCommand(data []byte, atEOF bool) (advance int, token []byte, err error) {
	advance, token, err = bufio.ScanWords(data, atEOF)
	if err == nil && token != nil {
		switch {
		case bytes.Equal(token, []byte{'e', 'x', 'i', 't'}):
			return 0, []byte{'[', 'I', 'n', 'f', 'o', ']', ',', ' ', 'B', 'y', 'e', '!'}, bufio.ErrFinalToken
		case bytes.Contains(data, []byte{' '}):
			strCommand := fmt.Sprintf("<<< this line replace by \"%s\" >>>", token)
			strCommandByte := []byte(strCommand)
			strCommandFinal := append(data, strCommandByte...)

			return 0, strCommandFinal, nil
		}
	}
	return
}

func splitCommand2(data []byte, atEOF bool) (advance int, token []byte, err error) {
	advance, token, err = bufio.ScanWords(data, atEOF)
	if err == nil && token != nil {
		return 0, token, err
	}
	return 0, token, bufio.ErrFinalToken
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a command and data: ")
	for scanner.Scan() {
		text := scanner.Text()
		if text == "exit" {
			break
		}
		fmt.Println(text)
		fmt.Print("Enter a command and data: ")
	}
	fmt.Print("[Info] Bye!")
}
