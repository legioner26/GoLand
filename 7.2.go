package main

import (
	"fmt"
)

func main() {

	fmt.Println("Введите ширину")
	var width int
	fmt.Scan(&width)

	fmt.Println("Введите высоту")
	var height int
	fmt.Scan(&height)

	for i := 0; i < height; i++ {
		if i%2 != 0 {
			for j := 0; j <= width; j = j + 2 {
				fmt.Print("*")
				fmt.Print(" ")
			}
		} else {
			for j := 2; j <= width; j = j + 2 {
				fmt.Print(" ")
				fmt.Print("*")
			}
		}
		fmt.Print("\n")
	}
}
