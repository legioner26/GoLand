package main

import (
	"fmt"
)

func main() {
	var one, two, three = 0, 0, 0
	for i := 1; i <= 1000; i++ {
		if one < 10 {
			one, two, three = i, i, i
		} else if two < 100 {
			two, three = i, i
		} else if three < 1000 {
			three = i
		}

	}

	fmt.Println("Значения переменных по итогу работы цикла:")
	fmt.Println("Первая переменная:", one)
	fmt.Println("Вторая переменная:", two)
	fmt.Println("Третья переменная:", three)
}
