package main

import "fmt"

func main() {
	defer hello()
}

func hello() {
	fmt.Println("Hello World")
}
