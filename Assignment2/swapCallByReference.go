package main

import "fmt"

func main() {
	var a int = 1
	var b int = 2
	fmt.Printf("Before swap,value of a:%d  b:%d\n", a, b)
	swap(&a, &b)
	fmt.Printf("After swap,value of a:%d  b:%d\n", a, b)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
