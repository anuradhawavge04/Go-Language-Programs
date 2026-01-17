package main

import "fmt"

func try(b int) (x, y int) {
	x = b + 4
	y = b - 4
	return
}

func main() {
	x, y := try(12)
	fmt.Println(x, y)
}
