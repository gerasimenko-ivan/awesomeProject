package main

import (
	"fmt"
	"testing"
)

func TestArrDoXAfterColon(t *testing.T) {
	var arrToProcess = []int{0, 1, 2, 3, 4}
	var arrExpect = []int{0, 1, 2}
	arrXAfterColon := arrDoXAfterColon(&arrToProcess, 3)
	if !testEq(arrExpect, arrXAfterColon, true) {
		t.Error()
	}
}

func testEq(a, b []int, doPrintDiff bool) bool {
	if len(a) != len(b) {
		if doPrintDiff {
			fmt.Printf("length is different: %v and %v", len(a), len(b))
		}
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			if doPrintDiff {
				fmt.Printf("Element [%v] is different: %v and %v", i, a[i], b[i])
			}
			return false
		}
	}
	return true
}
