package main

import "fmt"

type book struct {
	BookId int
	Title  string
	Author string
	Price  float32
}

func main() {
	var books [3]book
	for i := 0; i < len(books); i++ {
		fmt.Println("\nEnter details of book", i+1)

		fmt.Printf("Book ID:")
		fmt.Scanln(&books[i].BookId)

		fmt.Printf("Title:")
		fmt.Scanln(&books[i].Title)

		fmt.Printf("Author:")
		fmt.Scanln(&books[i].Author)

		fmt.Printf("Price:")
		fmt.Scanln(&books[i].Price)
	}

	fmt.Println("\n-------------- Book Details--------------")

	for i := 0; i < len(books); i++ {
		fmt.Println("\nBook ID:", books[i].BookId)
		fmt.Println("Title:", books[i].Title)
		fmt.Println("Author:", books[i].Author)
		fmt.Println("Price:", books[i].Price)
	}
}
