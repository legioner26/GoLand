package main

import (
	"fmt"
)

func multiply(number int) (x int) {
	x = number * 2
	return x
}

func increment(number int) (y int) {
	y = number + 1
	return y
}
func main() {

	fmt.Println("Введите число")
	var n int
	fmt.Scan(&n)

	fmt.Printf("Умножаем число %d на 2: %d\n", n, multiply(n))
	fmt.Printf("Прибавляем к числу %d единицу: %d\n", n, increment(n))
}
