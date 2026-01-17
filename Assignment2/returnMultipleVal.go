package main

import "fmt"

func main() {
	var v1, v2, v3 int
	fmt.Println("Enter 3 values:")
	fmt.Scanf("%d%d%d", &v1, &v2, &v3)
	v1, v2, v3 = myfunction(v1, v2)

	fmt.Println("Result is:%d", v1)
	fmt.Println("Result is:%d", v2)
	fmt.Println("Result is:%d", v3)
}

func myfunction(a, b int) (int, int, int) {
	return a - b, a * b, a + b
}
