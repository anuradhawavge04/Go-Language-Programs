package main

import "fmt"

// Define structure
type Student struct {
	name  string
	age   int
	marks int
}

// Method with pointer receiver
func (s *Student) show() {
	fmt.Println("Student Name:", s.name)
	fmt.Println("Age:", s.age)
	fmt.Println("Marks:", s.marks)
}

func main() {
	// Create object
	stu := Student{
		name:  "Anuradha",
		age:   20,
		marks: 85,
	}

	// Call method
	stu.show()
}
