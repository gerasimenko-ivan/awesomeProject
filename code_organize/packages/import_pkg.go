package main

// Programs start running in package main.
// So how do I have so many main packages??? :)

import ( // "factored" import style (imports in parenthesis) - it's GOOD!
	"fmt"
	"math/rand"
)

// multiple import statements - BAD STYLE
//import "fmt"
//import "math/rand"

func main() {
	// the package (e.g. "rand") name is the same as the last element of the import path, by convention
	fmt.Println("Random number is", rand.Intn(10)) // why returns the same number???
}
