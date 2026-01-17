package main

import "fmt"

func palindrome(num int) string {
	var original, remainder int
	reverse := 0
	original = num

	for num != 0 {
		remainder = num % 10
		reverse = reverse*10 + remainder
		num = num / 10

	}

	if original == reverse {
		return fmt.Sprintf("%d is  palindrome", reverse)
	} else {
		return fmt.Sprintf("%d is not a palindrome", reverse)
	}

}

func main() {
	var num int
	fmt.Println("Enter a number:\n")
	fmt.Scan(&num)
	result := palindrome(num)
	fmt.Println(result)
}
