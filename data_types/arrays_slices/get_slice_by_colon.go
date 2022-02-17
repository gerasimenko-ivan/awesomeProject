package main

import "fmt"

func main() {
	// Need to know difference:
	//var x = [6]int{0, 1, 2, 3, 4, 5} // ARRAY - fixed length // or [...]int{0}
	var x = []int{0, 1, 2, 3, 4, 5}

	l := arrDoXAfterColon(&x, 2) // Cannot use '&x' (type *[6]int) as the type *[]int
	h := arrDoXBeforeColon(&x, 2)
	fmt.Printf("%v\n", l)
	fmt.Printf("%v\n", h)
	fmt.Println("hello world")
}

// how to pass array/slice to func?
//  - by pointer - V
//  - by value - it is "expensive"
func arrDoXAfterColon(a *[]int, x int) []int { // pass slice by pointer
	// what is :2 or 2: in request to elements of array
	// lets do it as func
	return (*a)[:x] // returns slice up to X element, not including the X-element
}

func arrDoXBeforeColon(a *[]int, x int) []int { // pass slice by pointer
	return (*a)[x:] // returns slice from the X-th element, including the X-element
}
