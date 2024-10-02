package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done() // уменьшение счетчика wg при завершении горутины

	for i := 0; i < 1000; i++ {
		mutex.Lock()   // блокировка мьютекса
		counter++      // увеличение счетчика
		mutex.Unlock() // освобождение мьютекса
	}
}

func main() {
	/*
			4.	Синхронизация с помощью мьютексов:
		•	Реализуйте программу, в которой несколько горутин увеличивают общую переменную-счётчик.
		•	Используйте мьютексы (sync.Mutex) для предотвращения гонки данных.
		•	Включите и выключите мьютексы, чтобы увидеть разницу в работе программы.
	*/

	fmt.Printf("изначальный контейнер: %d\n", counter)

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment(&wg) // запуск 5 горутин
	}

	wg.Wait() // ожидание завершения всех горутин

	fmt.Printf("финальный контейнер: %d\n", counter)
}
