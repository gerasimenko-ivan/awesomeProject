package main

import "fmt"

func main() {
	var x = [6]int{0, 1, 2, 3, 4, 5}
	l := x[:2]
	h := x[2:]
	fmt.Println("%v", l)
	fmt.Println("%v", h)
	fmt.Println("hello world")
}
