package main

import "fmt"

// Structure
type Student struct {
	rollno     int
	name       string
	percentage float64
}

// Named slice type
type StudentList []Student

// Method on named type
func (s StudentList) showDescending() {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i].percentage < s[j].percentage {
				s[i], s[j] = s[j], s[i]
			}
		}
	}

	fmt.Println("\nStudents in Descending Order:")
	for _, v := range s {
		fmt.Println(v.rollno, v.name, v.percentage)
	}
}

func main() {
	var n int
	fmt.Print("Enter number of students: ")
	fmt.Scan(&n)

	// Use named type
	students := make(StudentList, n)

	for i := 0; i < n; i++ {
		fmt.Println("\nEnter details of student", i+1)

		fmt.Print("Roll No: ")
		fmt.Scan(&students[i].rollno)

		fmt.Print("Name: ")
		fmt.Scan(&students[i].name)

		fmt.Print("Percentage: ")
		fmt.Scan(&students[i].percentage)
	}

	// Call method
	students.showDescending()
}
