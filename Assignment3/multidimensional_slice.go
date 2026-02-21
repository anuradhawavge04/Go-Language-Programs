package main

import "fmt"

func main() {
	//creating multi dimensional slice
	s1 := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	fmt.Println("S1:", s1) //accesing slice

	s2 := [][]string{
		[]string{"A1", "A2", "A3"},
		[]string{"B1", "B2", "B3"},
	}
	fmt.Println("\nS2", s2) //accessing second slice

	fmt.Println("\n Multidimensional Slice s2:")
	for i := 0; i < len(s2); i++ {
		fmt.Println(s2[i])
	}

	fmt.Println("\n Printing element slice s2 like Matrix:")
	n := len(s2)
	m := len(s2[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(s2[i][j] + " ")
		}
		fmt.Print("\n")
	}

}
