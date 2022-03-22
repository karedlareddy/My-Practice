package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	strVar := "100"
	fmt.Println(reflect.TypeOf(strVar))
	intVar, err := strconv.Atoi(strVar)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))
	two()
	three()
}

func two() {
	strVar := "100"

	intVar, err := strconv.ParseInt(strVar, 0, 8)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	intVar, err = strconv.ParseInt(strVar, 0, 16)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	intVar, err = strconv.ParseInt(strVar, 0, 32)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	intVar, err = strconv.ParseInt(strVar, 0, 64)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))
}

func three() {
	strVar := "100"
	intValue := 0
	_, err := fmt.Sscan(strVar, &intValue)
	fmt.Println(intValue, err, reflect.TypeOf(intValue))
}


