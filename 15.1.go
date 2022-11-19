package main

import (
	"fmt"
)

func main() {

	fmt.Println("Введите 10 чисел")
	var array [10]int
	for i := 0; i < len(array); i++ {
		fmt.Scan(&array[i])
	}

	var countEven int
	var countOdd int
	for _, val := range array {
		if val%2 == 0 {
			countEven++
		} else {
			countOdd++
		}
	}
	fmt.Printf("Из введённых чисел %d : чётных — %d, нечётных — %d", array, countEven, countOdd)
}
