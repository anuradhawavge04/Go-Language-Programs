package main

import "fmt"

type Student struct {
	roll_no   int
	stud_name string
	mark1     float64
	mark2     float64
	mark3     float64
	total     float64
	average   float64
}

func main() {
	var n int
	fmt.Print("Enter number of students:")
	fmt.Scan(&n)

	students := make([]Student, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Enter students detail:%d\n", i+1)

		fmt.Print("Roll No:\n")
		fmt.Scan(&students[i].roll_no)

		fmt.Print("Student Name:\n")
		fmt.Scan(&students[i].stud_name)

		fmt.Print("Student Mark1:\n")
		fmt.Scan(&students[i].mark1)

		fmt.Print("Student Mark2:\n")
		fmt.Scan(&students[i].mark2)

		fmt.Print("Student Mark3:\n")
		fmt.Scan(&students[i].mark3)

		students[i].total = students[i].mark1 + students[i].mark2 + students[i].mark3
		students[i].average = students[i].total / 3

	}

	fmt.Println("\n Student Detail with Total and Average")
	for i := 0; i < n; i++ {
		fmt.Printf("\n Roll No:%d\n", students[i].roll_no)
		fmt.Printf("Student Name:%s\n", students[i].stud_name)
		fmt.Printf("Total:%.2f\n", students[i].total)
		fmt.Print("Average:%.2f\n", students[i].average)
	}
}
