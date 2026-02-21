package main

import "fmt"

type Employee struct {
	eno    int
	ename  string
	salary float64
}

func main() {
	var n int
	fmt.Printf("Enter number of employee:")
	fmt.Scan(&n)

	emp := make([]Employee, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Enter employee details:%d\n", i+1)

		fmt.Print("Enter Employee Number:")
		fmt.Scan(&emp[i].eno)

		fmt.Print("Enter Employee Name:")
		fmt.Scan(&emp[i].ename)

		fmt.Print("Enter employee Salary:")
		fmt.Scan(&emp[i].salary)

	}

	maxSalary := emp[0].salary
	for i := 1; i < n; i++ {
		if emp[i].salary > maxSalary {
			maxSalary = emp[i].salary
		}
	}

	fmt.Printf("\n Employees With Maximum Salary:\n", maxSalary)
	for i := 0; i < n; i++ {
		if emp[i].salary == maxSalary {
			fmt.Printf("\n Employee No:%d\n", emp[i].eno)
			fmt.Printf("\n Employee Name:%s\n", emp[i].ename)
			fmt.Printf("\n Employee Salary:%.2f\n", emp[i].salary)
		}
	}

}
