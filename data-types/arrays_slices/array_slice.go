package main

import "fmt"

func main() {
	// ARRAYS
	// - fixed length

	var array0 = [2]string{} // length 2 / initialized by type zeros '' (e.g. empty string in this case)
	printStrLen2ArrByPointer(&array0)

	var array1 = [...]int{0, 1, 2, 3} // length 4, cause 4 elements / array cause [...]
	printIntLen4ArrByValue(array1)    //fmt.Printf("arr: %v\n", array1)

	// iterate through
	for i, el := range array0 {
		fmt.Printf("i=%v el='%v'\n", i, el)
	}

	// SLICES

}

// how to pass array to func - BY POINTER
func printStrLen2ArrByPointer(arr *[2]string) {
	fmt.Printf("arr: %v, length: %v\n", arr, len(arr))
}

// how to pass array to func - BY VALUE
func printIntLen4ArrByValue(arr [4]int) {
	fmt.Printf("arr: %v, length: %v\n", arr, len(arr))
}
