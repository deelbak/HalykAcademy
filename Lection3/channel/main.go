package main

import "fmt"

func greet(c chan int) {
	for i := 0; i < 5; i++ {
		data := <-c
		fmt.Println(data)
	}

}
func main() {
	fmt.Println("Start")
	//c := make(chan int, 3)
	////fmt.Printf("type of chan is %T\n", c)
	////fmt.Printf("value of chan is %v\n", c)
	////go func() {
	////	c <- "adil"
	////}()
	////data := <-c
	//
	////fmt.Println(data)
	//go greet(c)
	//c <- 8
	//c <- 7
	//c <- 16
	////c <- 15
	ch := make(chan int, 3)

	// Отправка данных в канал без блокировки
	ch <- 1
	ch <- 2
	ch <- 3

	// Прием данных из канала без блокировки
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println("Stop")
}
