package main

import (
	"fmt"
	"sync"
)

type SharedData struct {
	mu    sync.Mutex
	value int
}

func main() {
	// Создаем экземпляр SharedData со значением 42
	data := SharedData{value: 42}

	// Теперь вы можете работать с данными и использовать мьютекс при необходимости
	// Например, читаем значение из одной горутины:
	go func() {
		data.mu.Lock()         // Захватываем мьютекс для чтения данных
		defer data.mu.Unlock() // Гарантируем, что мьютекс будет разблокирован после чтения

		fmt.Println("Read Value:", data.value)
	}()

	// Модифицируем значение из другой горутины:
	go func() {
		data.mu.Lock()         // Захватываем мьютекс для записи данных
		defer data.mu.Unlock() // Гарантируем, что мьютекс будет разблокирован после записи

		data.value = 100
		fmt.Println("Updated Value:", data.value)
	}()

	// Ждем завершения горутин
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		data.mu.Lock()
		defer data.mu.Unlock()
		fmt.Println("Final Read Value:", data.value)
	}()
	go func() {
		defer wg.Done()
		data.mu.Lock()
		defer data.mu.Unlock()
		fmt.Println("Final Updated Value:", data.value)
	}()

	// Ждем завершения всех горутин
	wg.Wait()
}
