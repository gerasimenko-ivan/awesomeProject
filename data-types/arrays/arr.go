package main

import "fmt"

func main() {
	var x = [6]int{0, 1, 2, 3, 4, 5}
	l := x[:2]
	h := x[2:]
	fmt.Printf("%v\n", l)
	fmt.Printf("%v\n", h)
	fmt.Println("hello world")
}

func arrDo() {
	// how to pass array to func?
	// what is :2 or 2: in request to elements of array
	// lets do it as func
}
