package main

import "fmt"

func main() {
	// ARRAYS -------------------------------------------------------
	// - fixed length

	array0 := [2]string{} // length 2 / initialized by type zeros '' (e.g. empty string in this case)
	printStrLen2ArrByPointer(&array0)

	array1 := [...]int{0, 1, 2, 3} // length 4, cause 4 elements / array cause [...]
	printIntLen4ArrByValue(array1) //fmt.Printf("arr: %v\n", array1)

	// iterate through array
	fmt.Printf("-- for range\n")
	for i, el := range array0 {
		fmt.Printf("j=%v el='%v'\n", i, el)
	}

	fmt.Printf("-- for three-component\n")
	for i := 0; i < len(array1); i++ {
		fmt.Printf("j=%v el='%v'\n", i, array1[i])
	}

	fmt.Printf("-- for only-condition\n")
	i := 0
	for i < len(array1) {
		fmt.Printf("j=%v el='%v'\n", i, array1[i])
		i++
	}

	// SLICES -------------------------------------------------------
	// length - ???
	// capacity - ???

	slice0 := []int{1, 2, 3} // ? initialized directly, ? so it is a slice literal

	slice1 := make([]float32, 2, 4) // args: array, length, capacity

	slice2 := make([]string, 3) // args: array, length-capacity

	// iterate through slice
	fmt.Printf("-- for range\n")
	for i, el := range slice0 {
		fmt.Printf("j=%v el='%v'\n", i, el)
	}

	fmt.Printf("-- for three-component\n")
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("j=%v el='%v'\n", i, slice1[i])
	}

	fmt.Printf("-- for only-condition\n")
	j := 0
	for j < len(slice2) {
		fmt.Printf("j=%v el='%v'\n", j, slice2[2])
		j++
	}

	// let's figure out what is length & capacity by testing append()
	fmt.Printf("Check length VS capacity in slice\n")

	slice3 := make([]int, 2, 3)
	slice3[0] = 0
	slice3[1] = 1
	for i, el := range slice3 {
		fmt.Printf("i=%v, el=%v\n", i, el)
	}
}

// how to pass array to func - BY POINTER
func printStrLen2ArrByPointer(arr *[2]string) {
	fmt.Printf("arr: %v, length: %v\n", arr, len(arr))
}

// how to pass array to func - BY VALUE
func printIntLen4ArrByValue(arr [4]int) {
	fmt.Printf("arr: %v, length: %v\n", arr, len(arr))
}
