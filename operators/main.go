package main

import (
	"fmt"
)

func main() {
	var x, y = 35, 7

	fmt.Printf("x + y = %d\n", x+y)
	fmt.Printf("x - y = %d\n", x-y)
	fmt.Printf("x * y = %d\n", x*y)
	fmt.Printf("x / y = %d\n", x/y)
	fmt.Printf("x mod y = %d\n", x%y)

	x++
	fmt.Printf("x++ = %d\n", x)

	y--
	fmt.Printf("y-- = %d\n", y)
	one()
	two()
	three()
	four()
}

func one() {
	var x, y = 15, 25
	x = y
	fmt.Println(x)

	x = 15
	x += y
	fmt.Println(x)

	x = 50
	x -= y
	fmt.Println(x)

	x = 2
	x *= y
	fmt.Println(x)

	x = 100
	x /= y
	fmt.Println(x)

	x = 40
	x %= y
	fmt.Println(x)
}

func two() {
	var x, y = 15, 25

	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x < y)
	fmt.Println(x <= y)
	fmt.Println(x > y)
	fmt.Println(x >= y)
}

func three() {

	var x, y, z = 10, 20, 30
	fmt.Println(x < y && x > z)
	fmt.Println(x < y || x > z)
	fmt.Println(!(x == y && x > z))
}

func four() {

	var x int = 9
	var y int = 65
	var z int

	z = x & y
	fmt.Println(z)

	z = x | y
	fmt.Println(z)

	z = x ^ y
	fmt.Println(z)

	z = x << 1
	fmt.Println(z)

	z = x >> 1
	fmt.Println(z)
}
