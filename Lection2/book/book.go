package book

import "fmt"

type Book struct {
	Author string
	Name   string
}

func (b Book) String() string {
	return fmt.Sprintf("Book: %s - %s", b.Author, b.Name)
}
