package main

import (
	"fmt"
)

func Reverse(input [10]int) [10]int {
	inputLen := len(input)
	var output [10]int

	for i, n := range input {
		j := inputLen - i - 1

		output[j] = n
	}

	return output
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
