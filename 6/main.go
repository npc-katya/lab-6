package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

// worker функция, которая будет обрабатывать задачи
func worker(id int, jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		// Реверсируем строку
		reversed := reverseString(job)
		fmt.Printf("Worker %d processed: %s -> %s\n", id, job, reversed)
		results <- reversed
	}
}

// функция для реверсирования строки
func reverseString(s string) string {

	// переворачивание текста
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	/*
			6.	Создание пула воркеров:
		•	Реализуйте пул воркеров, обрабатывающих задачи (например, чтение строк из файла и их реверсирование).
		•	Количество воркеров задаётся пользователем.
		•	Распределение задач и сбор результатов осуществляется через каналы.
		•	Выведите результаты работы воркеров в итоговый файл или в консоль.
	*/

	var numWorkers int
	fmt.Print("введите количество воркеров: ")
	fmt.Scan(&numWorkers)

	jobs := make(chan string, 100)
	results := make(chan string, 100)

	var wg sync.WaitGroup

	// запуск воркеров
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// чтение строки из файла
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		jobs <- scanner.Text() // отправление строки в канал задач
	}
	close(jobs) // закрытие канала задач после завершения чтения

	// ожидание завершения обработки задач
	go func() {
		wg.Wait()
		close(results) // закрытие канала результатов после завершения всех воркеров
	}()

	// сохранение результатов в файл
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("ошибка при создании файла:", err)
		return
	}
	defer outputFile.Close()

	for result := range results {
		outputFile.WriteString(result + "\n") // запись результатов в файл
	}

	fmt.Println("обработка завершена. Результаты записаны в output.txt.")

}
