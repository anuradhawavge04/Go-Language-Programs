package main

import "fmt"

func main() {
	var num int = 9
	fmt.Println("Table of :", num)
	for i := 1; i <= 10; i++ {
		fmt.Println(num * i)
	}
}
