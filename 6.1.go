package main

import (
	"fmt"
)

func main() {
	fmt.Println("Цикл")
	var num int
	fmt.Println("Введите целое число")
	fmt.Scan(&num)
	for i := 0; i < num; i++ {
		fmt.Println(i)
	}
}
