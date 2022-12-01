package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Введите число или 'стоп' чтобы завершить программу")
	var content string
	for {
		fmt.Scan(&content)
		if content == "стоп" {
			break
		}

		num, err := strconv.Atoi(content)
		if err != nil {
			fmt.Println("Ошибка:", err)
		}

		square := square(num)
		mult := multiplication(square)
		fmt.Println("Произведение:", <-mult)
	}
}

func square(num int) chan int {
	fChan := make(chan int)
	go func() {
		fChan <- num * num
	}()
	return fChan
}

func multiplication(fChan chan int) chan int {
	sChan := make(chan int)
	fChanRes := <-fChan
	fmt.Println("Квадрат:", fChanRes)
	go func() {
		sChan <- fChanRes * 2
	}()
	return sChan
}
