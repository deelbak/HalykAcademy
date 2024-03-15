package main

import (
	"fmt"
	"sync"
)

var counter int = 0

func work(number int, ch chan bool, mutex *sync.Mutex) {
	mutex.Lock()
	counter = 0
	for k := 0; k < 5; k++ {
		counter++
		fmt.Println("Goroutine: ", number, " - ", counter)
	}
	mutex.Unlock()
	ch <- true
}
func main() {
	ch := make(chan bool)
	var mutex sync.Mutex
	for i := 0; i < 5; i++ {
		go work(i, ch, &mutex)
	}
	for i := 0; i < 5; i++ {
		<-ch
	}
	fmt.Println("End")
}
