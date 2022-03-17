package main

import (
	"fmt"
	"reflect"
)

var s = "Japan"

func main() {
	var i int = 10
	var s string = "canada"

	fmt.Println(i)
	fmt.Println(s)
	add()
	mul()
	set()
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(s))
	brace()
	lol()
	dec()
}

func add() {
	i := 10
	s := 15
	result := i + s

	fmt.Println(result)
}

func mul() {
	var fname, lname string = "John", "Doe"
	m, n, o := 1, 2, 3
	item, price := "Mobile", 2000
	fmt.Println(fname + lname)
	fmt.Println(m + n + o)
	fmt.Println(item, "-", price)
}

func set() {
	name := "Bhargav Reddy"
	fmt.Println(reflect.TypeOf(name))
}

func brace() {
	fmt.Println(s)
	x := true
	if x {
		y := 1
		if x != false {
			fmt.Println(s)
			fmt.Println(x)
			fmt.Println(y)
		}
	}
	fmt.Println(x)
}

func lol() {
	var quantity float32 = 30
	fmt.Println(quantity)

	var price int16 = 12
	fmt.Println(price)

	var product string = "fan"
	fmt.Println(product)

	var inStock bool = true
	fmt.Println(inStock)
}

var (
	product  = "Mobile"
	quantity = 50
	price    = 120
	inStock  = true
)

func dec() {
	fmt.Println(product)
	fmt.Println(quantity)
	fmt.Println(price)
	fmt.Println(inStock)
}
