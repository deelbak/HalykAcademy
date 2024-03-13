package main

import (
	"awesomeProject/twosum"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func calc(x int, y int) int {
	return x + y
}
func main() {
	defer fmt.Println("It is defer")
	fmt.Println(calc(5, 9))
	data := "https://web.telegram.org/k/#@adlbk"
	err := validation.Validate(data,
		validation.Required,
		validation.Length(5, 100),
		is.URL,
	)
	fmt.Println(err)
	fmt.Println(twosum.Two())
}
