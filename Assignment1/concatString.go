package main

import "fmt"

func main() {
	var str1, str2 string
	fmt.Println("Enter first string:")
	fmt.Scanf("%s\n", &str1)
	fmt.Println("Enter Second String:")
	fmt.Scanf("%s\n", &str2)

	//pointers to string
	p1 := &str1
	p2 := &str2

	result := *p1 + *p2

	fmt.Println("Concatnated String:", result)
}
