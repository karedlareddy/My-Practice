package main

import "fmt"

// Global variable declaration and addittion

var (
	m int
	n int
)

func main() {
	var x int = 1 // integer data type
	var y int     // integer data type
	fmt.Println(x)
	fmt.Println(y)

	var a, b, c = 5.25, 5.35, 5.45 // multiple float32
	fmt.Println(a, b, c)

	city := "Hyderabad" // string varibale declaration
	country := "India"  // variable names are case sensitive
	fmt.Println(city)
	fmt.Println(country) // variable names are case sensitive

	food, drink, price := "Pizza", "Pepsi", 125
	fmt.Println(food, drink, price)
	m, n = 2, 3
	fmt.Println(m, n)
}

