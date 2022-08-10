package main

import "fmt"

func getUserChoice() string {
	var choice string
	fmt.Scanln(&choice)
	return choice
}

func dragonReplyToAnswer(ans string) {
	if ans == "42" {
		fmt.Println("You are right!")
	} else if ans == "I don't know." {
		fmt.Println("You are wrong!")
	} else {
		fmt.Println("You are so afraid!")
	}
}

func dragonReactToAction(action string) {
	switch action {
	case "Answer":
		var ans = getUserChoice()
		dragonReplyToAnswer(ans)
	case "Trick":
		dragonDenyAnswer()
	case "Run":
		dragonBecomeAngry()
	}
}

func dragonDenyAnswer() {
	fmt.Println("I like what you're trying to do, but i can't help those who are too lazy to think!")
}

func dragonBecomeAngry() {
	fmt.Println("Arr!! How dare you run away!?")
}

func main() {
	fmt.Println("What is the meaning of life?")
	var action = getUserChoice()
	dragonReactToAction(action)
}
