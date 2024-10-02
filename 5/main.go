package main

import (
	"fmt"
	"sync"
)

// структура для запроса
type Request struct {
	Operation string
	A         float64
	B         float64
}

// структура для ответа
type Response struct {
	Result float64
	Error  error
}

// каналы для запросов и ответов
var requestChan = make(chan Request)
var responseChan = make(chan Response)

// функция калькулятора
func calculator() {
	for req := range requestChan {
		var result float64
		var err error

		switch req.Operation {
		case "+":
			result = req.A + req.B
		case "-":
			result = req.A - req.B
		case "*":
			result = req.A * req.B
		case "/":
			if req.B != 0 {
				result = req.A / req.B
			} else {
				err = fmt.Errorf("деление на ноль")
			}
		default:
			err = fmt.Errorf("неизвестная операция: %s", req.Operation)
		}

		// передача результата в канал
		responseChan <- Response{Result: result, Error: err}
	}
}

func main() {

	/*
			Разработка многопоточного калькулятора:
		•	Напишите многопоточный калькулятор, который одновременно может обрабатывать запросы на выполнение простых операций (+, -, *, /).
		•	Используйте каналы для отправки запросов и возврата результатов.
		•	Организуйте взаимодействие между клиентскими запросами и серверной частью калькулятора с помощью горутин.

	*/

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		calculator()
	}()

	// операции
	requests := []Request{
		{Operation: "+", A: 15, B: 4},
		{Operation: "-", A: 25, B: 15},
		{Operation: "*", A: 13, B: 5},
		{Operation: "/", A: 44, B: 12},
		{Operation: "/", A: 6, B: 0},
	}

	// вывод результата
	for _, req := range requests {
		requestChan <- req

		response := <-responseChan
		if response.Error != nil {
			fmt.Printf("ошибка: %s\n", response.Error)
		} else {
			fmt.Printf("%.2f %s %.2f = %.2f\n", req.A, req.Operation, req.B, response.Result)
		}
	}

	// закрытие каналов
	close(requestChan)
	wg.Wait()
	close(responseChan)
}
