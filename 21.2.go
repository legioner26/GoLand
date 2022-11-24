package main

import (
	"fmt"
)

func matfunc(x int, y int, A func(int, int) int) int {
	defer fmt.Println("Результат:")
	return A(x, y)
}

func main() {
	one := matfunc(23, 3, func(a int, b int) int { return a - b })
	fmt.Printf("первая функция вернула - %d \n", one)

	two := matfunc(12, 6, func(a int, b int) int { return a + b })
	fmt.Printf("вторая функция вернула - %d \n", two)

	three := matfunc(3, 7, func(a int, b int) int { return a * b })
	fmt.Printf("третья функция вернула - %d \n", three)

}
