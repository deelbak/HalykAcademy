package main

import "fmt"

// Функция для слияния данных из нескольких каналов
func mergeChannels(channels ...<-chan int) <-chan int {
	// Создаем канал для объединенных данных
	merged := make(chan int)

	// Запускаем горутину для слияния данных
	go func() {
		defer close(merged) // Закрываем канал после завершения работы

		// Функция для чтения данных из канала и отправки их в объединенный канал
		for _, ch := range channels {
			for v := range ch {
				merged <- v
			}
		}
	}()

	return merged
}

func main() {
	// Создаем тестовые каналы
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	// Заполняем каналы тестовыми данными
	go func() {
		defer close(ch1)
		for i := 0; i < 3; i++ {
			ch1 <- i
		}
	}()
	go func() {
		defer close(ch2)
		for i := 3; i < 6; i++ {
			ch2 <- i
		}
	}()
	go func() {
		defer close(ch3)
		for i := 6; i < 9; i++ {
			ch3 <- i
		}
	}()

	// Вызываем функцию слияния для передачи тестовых каналов
	merged := mergeChannels(ch1, ch2, ch3)

	// Читаем данные из объединенного канала и выводим их
	for v := range merged {
		fmt.Println("Received:", v)
	}
}
