package main

import (
	"fmt"
)

func matFunc(x int, y int, A func(int, int) int) int {
	defer fmt.Println("Результат:")
	return A(x, y)
}

func main() {
	one := matFunc(23, 3, func(a int, b int) int { return a - b })
	fmt.Printf("первая функция вернула - %d \n", one)

	two := matFunc(12, 6, func(a int, b int) int { return a + b })
	fmt.Printf("вторая функция вернула - %d \n", two)

	three := matFunc(3, 7, func(a int, b int) int { return a * b })
	fmt.Printf("третья функция вернула - %d \n", three)

}
