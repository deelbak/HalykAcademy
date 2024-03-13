package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Goroutine")
	}()
	time.Sleep(10 * time.Second)
	fmt.Println("TimeSecond")
}
