package main

import (
	"fmt"
)

func main() {
	one := []int{1, 2, 5, 6}
	two := []int{3, 4, 7, 8, 9}
	fmt.Println("Массив а: ", one)
	fmt.Println("Массив b: ", two)

	var array = append(one, two...)
	fmt.Println("Объединенный массив: ", array)

	for a := 0; a < (len(array)); a++ {

		for b := (a + 1); b < len(array); b++ {
			if array[a] > array[b] {
				c := array[a]
				array[a] = array[b]
				array[b] = c
			}
		}

	}
	fmt.Println("Объединенный отсортированный массив: ", array)
}
