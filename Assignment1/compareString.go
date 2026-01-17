package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1, str2 string
	fmt.Println("Enter First String:")
	fmt.Scanf("%s\n", &str1)
	fmt.Println("Enter Second String:")
	fmt.Scanf("%s\n", &str2)

	fmt.Println(strings.Compare(str1, str2))
}
