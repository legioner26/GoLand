package main

import (
	"fmt"
)

func main() {
	array := []int{1, 2, 6, 3, 5, 4}
	fmt.Println("Исходный массив: ", array)

	for a := 0; a < (len(array)); a++ {

		for b := (a + 1); b < len(array); b++ {
			if array[a] > array[b] {
				c := array[a]
				array[a] = array[b]
				array[b] = c
			}
		}

	}

	fmt.Println("Массив после сортировки пузырьком: ", array)

}
