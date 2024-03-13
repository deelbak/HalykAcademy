package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			z := x
			x = y
			y = z + y
		case <-quit:
			return
		}
	}
}

func main() {
	//ch1 := make(chan string)
	//ch2 := make(chan string)
	//
	//go func() {
	//	time.Sleep(2 * time.Second)
	//	ch1 <- "one"
	//}()
	//
	//go func() {
	//	time.Sleep(1 * time.Second)
	//	ch2 <- "two"
	//}()
	//
	//// Select выбирает первый готовый канал для чтения
	//select {
	//case msg1 := <-ch1:
	//	fmt.Println("Received:", msg1)
	//case msg2 := <-ch2:
	//	fmt.Println("Received:", msg2)
	//}

	fmt.Println("Start")
	c := make(chan int)
	q := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			fmt.Print(<-c, " ")
		}
		q <- 0
	}()
	fibonacci(c, q)
	fmt.Println("")
	fmt.Println("Stop")
}
