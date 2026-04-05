package main

import "fmt"

// Define structure
type Author struct {
	name  string
	age   int
	books int
}

// Method with struct receiver
func (a Author) show() {
	fmt.Println("Author Name:", a.name)
	fmt.Println("Age:", a.age)
	fmt.Println("Number of Books:", a.books)
}

func main() {
	// Create object
	author1 := Author{
		name:  "John Doe",
		age:   40,
		books: 5,
	}

	// Call method
	author1.show()
}
