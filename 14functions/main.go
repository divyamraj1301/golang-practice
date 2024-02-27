package main

import "fmt"

func main() {
	fmt.Println("Functions in golang")
	greet()

	// greetTwo()

	result := sum(3, 5)
	fmt.Println("Sum is : ", result)

	multiSumResult, message := multiSum(2, 5, 8, 7, 3)
	fmt.Println("Multi Sum is : ", multiSumResult)
	fmt.Println("Message is : ", message)
}

// func greetTwo() {
// 	fmt.Println("greet Two called")
// }

func sum(num1 int, num2 int) int { // what kind of value you expect in the output
	return num1 + num2
}

func multiSum(inputs ...int) (int, string) {
	total := 0
	for _, val := range inputs {
		total += val
	}
	return total, "Multi Sum"
}

func greet() {
	fmt.Println("Greetings from GoLang")
}
