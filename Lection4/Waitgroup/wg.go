package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		wg      sync.WaitGroup
		datamap = make(map[int]string)
	)
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(j int) {
			datamap[j] = fmt.Sprintf("test %v", j)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("Done.")
	for k, v := range datamap {
		fmt.Println(fmt.Sprintf("for loop: key: %d, value: %s", k, v))
	}
}
