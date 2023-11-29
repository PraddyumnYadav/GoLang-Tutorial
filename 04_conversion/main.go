package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("(=======================* Welcome to our Pizza App *=======================)")
	fmt.Println("Please Rate our Pizza app between 1 and 5 stars")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for Rating,", input)
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to Your Rating:", numRating+1)
	}

	// End the Programm
	fmt.Println("---------------------------------- The End ----------------------------------")
}
