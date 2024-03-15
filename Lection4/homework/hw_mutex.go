package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	data  int
	mutex sync.Mutex
)

func readData() {
	mutex.Lock()         // Блокируем мьютекс перед чтением данных
	defer mutex.Unlock() // Разблокируем мьютекс после завершения чтения
	fmt.Println("Read data:", data)
}

func writeData(value int) {
	mutex.Lock()         // Блокируем мьютекс перед записью данных
	defer mutex.Unlock() // Разблокируем мьютекс после завершения записи
	fmt.Println("Write data:", value)
	data = value
}

func main() {
	// Запускаем несколько горутин для чтения данных
	for i := 0; i < 5; i++ {
		go readData()
	}

	// Ждем некоторое время перед записью данных
	time.Sleep(1 * time.Second)

	// Запускаем горутину для записи данных
	go writeData(42)

	// Ждем завершения работы горутин
	time.Sleep(1 * time.Second)
}
