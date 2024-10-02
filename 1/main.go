package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
			1.	Создание и запуск горутин:
		•	Напишите программу, которая параллельно выполняет три функции (например, расчёт факториала, генерация случайных чисел и вычисление суммы числового ряда).
		•	Каждая функция должна выполняться в своей горутине.
		•	Добавьте использование time.Sleep() для имитации задержек и продемонстрируйте параллельное выполнение.
	*/

	n := 2
	array := make([]int, n)

	// заполнение массива
	for i := 0; i < n; i++ {
		array[i] = rand.Intn(100)
	}

	// n := 5

	go factorial(2)
	go randx()
	go sum(array)
	time.Sleep(1 * time.Second)
}

func sum(array []int) (x int) {
	for i := 0; i < len(array); i++ {
		x += array[i]
	}

	fmt.Println("сумма числового ряда", x)
	return x
}

func randx() (x int) {
	x = rand.Intn(100)
	fmt.Println("генерация случайных чисел", x)
	return x
}

// функция для вычисления факториала
func factorial(x int) (y int) {

	y = 1
	for i := 1; i <= x; i++ {
		y = y * i
	}

	fmt.Println("факториал", y)

	return y
}
