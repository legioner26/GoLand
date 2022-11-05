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

	if one > 0 || two > 0 || three > 0 {
		fmt.Println("Одно или несколько из введенных чисел положительное")
	} else {
		fmt.Println("Положительных чисел нет")
	}
}
