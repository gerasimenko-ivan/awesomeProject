package main

import "fmt"

func main() {
	type DayOfWeek int
	const (
		SUN DayOfWeek = iota // SUN = 0; iota - is a function to generate set of constants
		MON
		TUE
		WED
		THU
		FRI
		SAT
	)

	var dayOfWeek DayOfWeek = THU
	fmt.Printf("day='%v'", dayOfWeek) // day='4'

	// it is more readable in code:
	if dayOfWeek == SAT || dayOfWeek == SUN {
		fmt.Printf("Don't go to work!")
	} else {
		fmt.Printf("Go to work!")
	}
}
