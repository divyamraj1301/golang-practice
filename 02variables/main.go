package main

import "fmt"

// jwtToken := 300000  // this will throw error as walrus operator (:=) is allowed onyl inside main function
var jwtToken int = 300000 // this syntax has to be followed

const LoginToken string = "random value" //  first letter capital - public variable

func main() {
	var username string = "Divyam"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n ", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n ", isLoggedIn)

	var smallVal int = 256
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n ", smallVal)

	var smallFloat float32 = 255.464773975553353535
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n ", smallFloat)

	var largeFloat float64 = 255.464773975553353535
	fmt.Println(largeFloat)
	fmt.Printf("Variable is of type: %T \n ", largeFloat)

	// default values and aliases
	var anotherVariable int // Default value is zero, so it will be set to 0
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var stringVariable string // Default value is empty string
	fmt.Println(stringVariable)
	fmt.Printf("Variable is of type: %T \n", stringVariable)

	// implicit type
	var website = "divyamraj1301.netlify"
	fmt.Println(website)

	// no var style
	numberOfUser := 300000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
