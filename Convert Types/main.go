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
}
