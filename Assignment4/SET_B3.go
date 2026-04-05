package main

import "fmt"

// Method to copy array elements
func copyArray(src [5]int) [5]int {
	var dest [5]int

	for i := 0; i < len(src); i++ {
		dest[i] = src[i]
	}

	return dest
}

func main() {
	// Source array
	arr1 := [5]int{10, 20, 30, 40, 50}

	// Call method
	arr2 := copyArray(arr1)

	// Print arrays
	fmt.Println("Source Array:", arr1)
	fmt.Println("Copied Array:", arr2)
}
