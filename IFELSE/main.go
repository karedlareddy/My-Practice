package main

import "fmt"


func main(){
	var s ="japan"
	x := true
	if x {
		fmt.Println(s)
	}
	one()
	two()
	three()

}

func one(){
	x :=100

	if x == 50{
		fmt.Println("japan")
	} else {
		fmt.Println("india")
	}
}

func two(){
	x :=100

	if x == 50{
		fmt.Println("japan")
	} else if x ==100 {
		fmt.Println("America")
	} else {
		fmt.Println("canada")
	}
}

func three(){
	if x :=100; x ==100 {
		fmt.Println("germany")
	}
}