package main

import "fmt"

func main() {
	fmt.Println("Defer in golang")
	defer fmt.Println("Deferred statement 1")
	defer fmt.Println("Deferred statement 2")
	defer fmt.Println("Deferred statement 3")

	fmt.Println("Non deferred statement")

	myDefer()

	// defer statements are executed at the very last - auto placed at this position
	// sequence -> 4, 3, 2, 1, 0, Deferred statement 3, Deferred statement 2, Deferred statement 1
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
