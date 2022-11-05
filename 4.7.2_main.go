package main

import (
	"fmt"
)

func main() {
	intMax := 5
	quantity := 0
	fmt.Println("Три числа.")
	fmt.Println("Введите первое число:")
	var resultOne int
	fmt.Scan(&resultOne)
	if resultOne > intMax {
		quantity++
	}
	fmt.Println("Введите второе число:")
	var resultTwo int
	fmt.Scan(&resultTwo)
	if resultTwo > intMax {
		quantity++
	}
	fmt.Println("Введите третье число:")
	var resultThree int
	fmt.Scan(&resultThree)
	if resultThree > intMax {
		quantity++
	}
	if quantity > 0 {
		fmt.Println("Среди введённых чисел есть число больше ", intMax, ".")
	} else {
		fmt.Println("Среди введённых чисел нет чисел больше ", intMax, ".")
	}

}
