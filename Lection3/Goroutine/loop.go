package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
		fmt.Println("it is", i)

	}

	time.Sleep(1 * time.Second)
	fmt.Println("time")
}
