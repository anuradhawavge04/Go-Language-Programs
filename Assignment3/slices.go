package main

import (
	"fmt"
	"sort"
)

func main() {
	number := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 1, 2, 3, 4, 6, 78}

	countries := []string{"India", "Russia", "America", "China", "France"}
	fmt.Println("Numbers:", number)

	//creating slice from another slice

	belowTen := number[10:]
	fmt.Println("Numbers Below 10=", belowTen)

	//Iterating over a slice using forloop
	fmt.Println("Using for loop:")
	for i := 0; i < len(number); i++ {
		fmt.Println(i, number[i])
	}

	fmt.Println("\n Using range keyword:")
	for i, value := range number {
		fmt.Println(i, value)
	}

	//Appending one slice to another
	Name := []string{"A", "B", "C", "D", "E", "F"}
	AppendData := append(Name, countries...)

	fmt.Println("Names:", Name)
	fmt.Println("Countries:", countries)

	fmt.Println("After Appending Names and Countries:", AppendData)

	//Copying one slice to another
	copySlice := make([]string, len(countries))
	fmt.Println("Countries:", countries)
	fmt.Println("copySlice", copySlice)
	copy(copySlice, countries)
	fmt.Println("After copy countries=", countries)

	//sorting slice
	fmt.Println("Before Sorting:", number)
	sort.Ints(number)
	fmt.Println("After sorting:", number)
}
