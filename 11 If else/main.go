package main

import "fmt"

func main() {
	fmt.Println("If else in GoLang")

	Number := 10
	var result string

	if Number < 10 {
		result = "Number less than 10"
	} else if Number > 10 {
		result = "Number greater than 10"
	} else {
		result = "Number is exactly equal to 10"
	}

	fmt.Println(result)

	if Number%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 5; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Printf("num is not less than 10")
	}

}
