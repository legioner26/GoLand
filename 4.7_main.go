package main

import "fmt"

func main() {
	balMax := 275
	fmt.Println("Балы ЕГЭ.")
	fmt.Println("Введите результат первого экзамена:")
	var resultOne int
	fmt.Scan(&resultOne)
	fmt.Println("Введите результат второго экзамена:")
	var resultTwo int
	fmt.Scan(&resultTwo)
	fmt.Println("Введите результат третьего экзамена:")
	var resultThree int
	fmt.Scan(&resultThree)
	fmt.Println("Сумма проходных баллов:", balMax)
	resultMax := resultOne + resultTwo + resultThree
	fmt.Println("Количество набранных баллов:", resultMax)
	if resultMax >= balMax {
		fmt.Println("Вы поступили")
	} else {
		fmt.Println("Вы не поступили")
	}
}
