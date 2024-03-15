package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context, ch chan string) {
	fmt.Println("doSomething is sleeping...")
	time.Sleep(time.Second * 2)
	fmt.Println("doSomething Wake up!")
	ch <- "Did Something"
}
func main() {
	ch := make(chan string, 1)
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	go doSomething(ctxTimeout, ch)
	select {
	case <-ctxTimeout.Done():
		fmt.Printf("Context concelled: %v\n", ctxTimeout.Err())
	case result := <-ch:
		fmt.Printf("Received: %v\n", result)
	}

}
