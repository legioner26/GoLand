package main

import (
	"fmt"
)

func main() {
	var one, two, three int

	fmt.Println("Введите первое число")
	fmt.Scan(&one)
	fmt.Println("Введите второе число")
	fmt.Scan(&two)
	fmt.Println("Введите третье число")
	fmt.Scan(&three)

	if one == two || one == three || two == three {
		fmt.Println("Среди введенных чисел есть два или более одинаковых числа!")
	} else {
		fmt.Println("Нет одинаковых чисел")
	}
}
