package main

import (
	"fmt"
)

func Reverse(input [10]int) (output [10]int) {
	inputLen := len(input)

	for i, n := range input {
		j := inputLen - i - 1

		output[j] = n
	}

	return
}
func main() {

	fmt.Println("Введите 10 чисел")

	var array [10]int
	for i := 0; i < len(array); i++ {
		fmt.Scan(&array[i])
	}

	fmt.Printf("Введеный массив: %d\n", array)
	fmt.Printf("Обратный порядок: %d\n", Reverse(array))
}
