package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the Rating for Our Pizza:")

	// comma ok || error ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for Rating, ", input)
	fmt.Printf("Type of Rating is %T", input)
}