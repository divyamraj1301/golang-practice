package main

import (
	"fmt"
)

func main() {
	fmt.Println("Concept of Slices")

	// var fruitList = []string{"Apple", "Banana", "Guava"}
	// fmt.Printf("Type of data: %T\n", fruitList)

	// fruitList = append(fruitList, "Mango", "Banana")
	// fmt.Println(fruitList)

	// fruitList = append(fruitList[1:3]) // return elemet from 1st index to 2nd index, range last value is non-inclusive
	// fmt.Println(fruitList)

	// fruitList = append(fruitList[:3]) // will start from 0th index
	// fmt.Println(fruitList)

	// highScores := make([]int, 4)

	// highScores[0] = 234
	// highScores[1] = 945
	// highScores[2] = 465
	// highScores[3] = 867
	// // highScores[4] = 777  // this method will go out of bound
	// highScores = append(highScores, 555, 666, 321) // will reallocate the memory and all the values will be accomadated

	// fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted((highScores)))
	// sort.Ints(highScores)
	// fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted((highScores)))

	// remove value from slice on basis of index

	var courses = []string{"React", "JavaScript", "Node.js", "C++", "GoLang"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
