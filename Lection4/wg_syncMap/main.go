package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var m sync.Map
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(j int) {
			m.Store(j, fmt.Sprintf("test %v", j))
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("Done...")

	m.Range(func(key, value interface{}) bool {
		fmt.Println(fmt.Sprintf("for loop: key: %d, value: %s", key, value))
		return true
	})
}
