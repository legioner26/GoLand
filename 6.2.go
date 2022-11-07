package main

import (
	"fmt"
)

func main() {
	fmt.Println("Сумма двух чисел по единице")

	var one, Two, summ int

	fmt.Println("Введите первое число:")
	fmt.Scan(&one)
	fmt.Println("Введите второе число:")
	fmt.Scan(&Two)

	summ = one + Two

	for i := one; i < summ; i++ {
		one++
	}

	fmt.Println("Сумма двух чисел, сложенная с помощью цикла, равна ", one)

}
