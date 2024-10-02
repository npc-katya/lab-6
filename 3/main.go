package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	/*
			3.	Применение select для управления каналами:
		•	Создайте две горутины, одна из которых будет генерировать случайные числа, а другая — отправлять сообщения об их чётности/нечётности.
		•	Используйте конструкцию select для приёма данных из обоих каналов и вывода результатов в консоль.
		•	Продемонстрируйте, как select управляет многоканальными операциями.
	*/

	c1 := make(chan int)    // создание канала
	c2 := make(chan string) // создание канала

	go randx(c1)

	go r(c1, c2)

	/*
		for i := 0; i < 10; i++ {
			result1 := <-c1
			result2 := <-c2
			fmt.Println(result1)
			fmt.Println(result2)
		}
	*/

	for i := 0; i < 10; i++ {
		select {
		case result := <-c2:
			fmt.Println(result) // вывод результата проверки

			/*
				case result := <-c1:
					fmt.Println(result) // вывод результата проверки
			*/

		case <-time.After(5 * time.Second):
			fmt.Println("программа остановлена...")
			return // завершение программы через 5 секунд

		}
	}

	time.Sleep(1 * time.Second)

}

func randx(c1 chan<- int) {
	var x int
	for i := 0; i < 10; i++ {
		x = rand.Intn(100)
		c1 <- x
		//fmt.Println("aaa", x)
		//time.Sleep(time.Second * 1) // Задержка в 1 секунду
	}
	close(c1)
}

func r(c1 <-chan int, c2 chan<- string) {
	str := ""
	for num := range c1 {
		if num%2 == 0 {
			str = strconv.Itoa(num) + " чётное"
			c2 <- str
			//fmt.Println("aaa", str)
		} else {
			str = strconv.Itoa(num) + " нечётное"
			c2 <- str
			//fmt.Println("aaa", str)
		}
		//time.Sleep(time.Second * 1) // Задержка в 1 секунду
	}
	close(c2)
}
