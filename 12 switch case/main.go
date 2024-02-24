package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Concept of Switch case")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is : ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you may open")
	case 2:
		fmt.Println("You may move 2 spots")
	case 3:
		fmt.Println("You may move 3 spots")
		// fallthrough  keyword allows to execute next statement  also when the condition matches current case
		fallthrough
	case 4:
		fmt.Println("You may move 4 spots")
		fallthrough
	case 5:
		fmt.Println("You may move 5 spots")
	case 6:
		fmt.Println("You may move 6 spots and roll the dice again")
	default:
		fmt.Println("Invalid Dice Value")
	}
}
