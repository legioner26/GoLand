package main

import (
	"fmt"
)

func main() {
	var X, Y int
	fmt.Println("Введите число X")
	fmt.Scan(&X)

	fmt.Println("Введите число Y")
	fmt.Scan(&Y)

	if X == 0 && Y == 0 {
		fmt.Println("X и Y на пересечении осей")
	}
	if X > 0 && Y > 0 {
		fmt.Println("Точка принадлежит I четверти")
	}
	if X < 0 && Y > 0 {
		fmt.Println("Точка принадлежит II четверти")
	}
	if X < 0 && Y < 0 {
		fmt.Println("Точка принадлежит III четверти")
	}
	if X > 0 && Y < 0 {
		fmt.Println("Точка принадлежит IV четверти")
	}
}
