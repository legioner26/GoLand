package main

import (
	"fmt"
)

func main() {
	min := 100000
	max := 999999
	count := 0
	var num int

	for num = min; num <= max; num++ {
		a := num % 1000
		b := (num - a) / 1000
		if a == b {
			count++
		}
	}
	fmt.Println("Количество зеркальных билетов от 100000 до 999999:")
	fmt.Println(count)
}
