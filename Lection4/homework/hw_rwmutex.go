package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	data int
	rw   sync.RWMutex
)

func readData() {
	rw.RLock()         // Блокируем мьютекс для чтения
	defer rw.RUnlock() // Разблокируем мьютекс после завершения чтения
	fmt.Println("Read data:", data)
}

func writeData(value int) {
	rw.Lock()         // Блокируем мьютекс для записи
	defer rw.Unlock() // Разблокируем мьютекс после завершения записи
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
