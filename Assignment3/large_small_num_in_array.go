package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 40, 30, 90, 89}
	if len(arr) == 0 {
		fmt.Println("Array is empty.")
		return
	}
	min := arr[0]
	max := arr[0]

	for _, v := range arr[1:] {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
