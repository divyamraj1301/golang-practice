package main

import "fmt"

func main() {
	fmt.Println("Concept of loops")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("Index is %v and value is %v\n", index, day)
	}

	arbitraryValue := 1
	for arbitraryValue < 10 {
		if arbitraryValue == 2 {
			goto dev
		}
		if arbitraryValue == 5 {
			arbitraryValue++
			continue
		}
		fmt.Println("Value is : ", arbitraryValue)
		arbitraryValue++
	}
dev:
	fmt.Println("Value of dev is printed")

}
