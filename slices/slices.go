package slices

import (
	"slices"
)

func sliceMethods() {
	slice := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8, 9, 10}


	// Unpractices Conditional statements
	if slices.Equal(slice, slice2) {
		println("The slices are equal")
	} else {
		println("The slices are not equal")

	} 
}


