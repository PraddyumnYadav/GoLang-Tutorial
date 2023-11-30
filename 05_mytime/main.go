package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("(====================* Welcome to time study of Golang *====================)")
	fmt.Println("")

	// How to get current Time
	fmt.Println("Printing Current Time")
	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 03:04:05 Monday"))

	// Sparation
	fmt.Println("")

	// Creating Time
	fmt.Println("Creating Your Own Time")
	createDate := time.Date(2009, time.April, 12, 4, 24, 0, 0, time.UTC)
	fmt.Println(createDate.Format("01-02-2006 03:04:05 Monday"))
}
