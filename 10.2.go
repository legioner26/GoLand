package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Введите сумму")
	var money float64
	fmt.Scan(&money)

	fmt.Println("Введите ежемесячный процент")
	var percent int
	fmt.Scan(&percent)

	fmt.Println("Введите количество лет")
	var years int
	fmt.Scan(&years)

	months := years * 12
	moneyResult := money
	bankBenefit := 0.0

	for i := 0; i < months; i++ {

		monthlyBenefit := moneyResult * (float64(percent) / 100)

		intNumber, fractionalNumber := math.Modf(monthlyBenefit * 100) // возврат целого и дробной части числа.
		moneyResult += intNumber / 100
		bankBenefit += fractionalNumber / 100
	}
	moneyResult = math.Round(moneyResult*100) / 100

	fmt.Println("Сумма после начисления:", moneyResult)
	fmt.Println("Округленный доход: ", bankBenefit)
}
