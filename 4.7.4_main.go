package main

import (
	"fmt"
)

func main() {
	fmt.Println("Три числа")
	var one, two, three int
	num := 5
	count := 0
	fmt.Println("Введите первое число:")
	fmt.Scan(&one)
	if one >= num {
		count++
	}
	fmt.Println("Введите второе число:")
	fmt.Scan(&two)
	if two >= num {
		count++
	}
	fmt.Println("Введите третье число:")
	fmt.Scan(&three)
	if three >= num {
		count++
	}
	if count > 0 {
		fmt.Println("Среди введённых чисел", count, "больше или равны ", num, ".")
	} else {
		fmt.Println("Среди введённых чисел нет числа больше пяти.")
	}
}
