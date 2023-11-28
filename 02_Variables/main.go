package main

import "fmt"

// You are not allowed to use := syntax outside a function
const LoginToken string = "abcdefghijklmnopqrstuvwxyz" // Public(L is capital) Constant(const) Variable

func main() {
	fmt.Println("(===========================* Variables *===========================)")
	var username string = "praddyumn"
	fmt.Println("The Variable is:", username)
	fmt.Printf("Variable is of type: %T \n", username)

	fmt.Println("")

	var isLoggedIn bool = true
	fmt.Println("The Variable is:", isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	fmt.Println("")

	var age int = 14
	fmt.Println("The Variable is:", age)
	fmt.Printf("Variable is of type: %T \n", age)

	fmt.Println("")

	var height float32 = 5.6
	fmt.Println("The Variable is:", height)
	fmt.Printf("Variable is of type: %T \n", height)

	fmt.Println("")

	var smallVal uint8 = 255 // between 0 and 255
	fmt.Println("The Variable is:", smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	// default values and some aliases
	fmt.Println(" \n ")
	fmt.Println("(----------------------------* Default Values and some aliases *----------------------------)")
	var anotherVariable int
	fmt.Println("The Variable is:", anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type
	var website = "https://praddyumnyadav.netlify.app/"
	fmt.Println(website)

	// no var style
	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
