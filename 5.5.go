package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Решение квадратного уравнения")
	var a, b, c, disc, x, x1, x2 float64

	fmt.Println("Введите первое число")
	fmt.Scan(&a)
	fmt.Println("Введите второе число")
	fmt.Scan(&b)
	fmt.Println("Введите третье число")
	fmt.Scan(&c)
	disc = math.Pow(b, 2) - 4*a*c

	if disc < 0 {
		fmt.Println("Корней нет")
	}

	if disc == 0 {
		x = -b / (2 * a)
		fmt.Println("Корень один:", x)
	}

	if disc > 0 {
		x1 = (-b + math.Sqrt(disc)) / (2 * a)
		x2 = (-b - math.Sqrt(disc)) / (2 * a)
		fmt.Println("Два корня:", x1, x2)
	}
}
