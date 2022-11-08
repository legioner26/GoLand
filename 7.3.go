package main

import (
	"fmt"
)

func main() {
	fmt.Println("Вывод ёлочки")
	fmt.Println("Введите высоту ёлочки")

	var height int
	fmt.Scan(&height)
	height++
	for i := 2; i <= height; i++ {
		for space := height; space > i; space-- {
			fmt.Print(" ")
		}
		for star := 2; star < (i*2 - 1); star++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}
