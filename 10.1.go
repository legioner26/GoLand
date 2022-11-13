package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Разложение 'экспоненты' в ряд Тейлора")

	fmt.Println("Введите значение x")
	var x float64
	_, _ = fmt.Scan(&x)

	fmt.Println("Введите степень, равную количеству знаков после запятой")

	var n float64
	_, _ = fmt.Scan(&n)

	var epsilon float64
	epsilon = 1 / math.Pow(10, n)

	result := 0.0
	prevResult := 0.0
	fact := 1
	k := 0

	for {
		if k > 0 {
			fact *= k
		}

		result += math.Pow(x, float64(k)) / float64(fact)
		if math.Abs(result-prevResult) < epsilon {

			fmt.Println(result)
			break
		}
		k++
		prevResult = result

	}
}
