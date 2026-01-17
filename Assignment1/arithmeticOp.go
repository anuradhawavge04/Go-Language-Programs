package main

import "fmt"

func main() {
	var num1, num2, choice int
	fmt.Println("Enter num1:")
	fmt.Scanf("%d\n", &num1)
	fmt.Println("Enter num2:")
	fmt.Scanf("%d\n", &num2)

	fmt.Println("Enter Your Choice:1.Addition\t 2.Substraction\t 3.Multiplication\t 4.Division\t 5.Modulo")
	fmt.Scanf("%d\n", &choice)
	switch choice {
	case 1:

		fmt.Println("Addition of num1 & num2 is:", num1+num2)

	case 2:
		fmt.Println("Substraction of num1 & num2 is:", num1-num2)

	case 3:
		fmt.Println("Multiplication of num1 & num2 is:", num1*num2)

	case 4:
		fmt.Println("Division of num1 & num2 is:", num1/num2)

	case 5:
		fmt.Println("Modulo of num1 & num2 is:", num1%num2)

	default:
		fmt.Println("Please Enter Valid Digit.....")
	}

}
