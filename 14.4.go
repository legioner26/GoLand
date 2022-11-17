package main

import (
	"fmt"
)

var (
	oneGlob   = 100
	twoGlob   = 200
	threeGlob = 300
)

func One(a int) int {
	a = a + oneGlob
	return a
}

func Two(a int) int {
	a = a + twoGlob
	return a
}

func Three(a int) int {
	a = a + threeGlob
	return a
}
func main() {

	fmt.Println("Введите число")
	var number int
	fmt.Scan(&number)

	fmt.Println("Первая функция:", One(number))
	fmt.Println("Вторая функция:", Two(number))
	fmt.Println("Третья функция:", Three(number))
	fmt.Println("Поочередный вызов функций:", One(Two(Three(number))))
}
