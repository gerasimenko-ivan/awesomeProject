package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// the package (e.g. "rand") name is the same as the last element of the import path, by convention
	fmt.Println("Random number is", rand.Intn(10)) // why returns the same number?
}
