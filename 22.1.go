package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 10

func main() {
	array := generate(100)
	fmt.Println("сгенерирован массив", array)
	fmt.Println("Введите число для поиска")
	var number int
	fmt.Scan(&number)

	countResult := getCount(array, number)
	if countResult > 0 {
		fmt.Println("Остаток чисел в массиве после найденной", countResult)
	} else {
		fmt.Println("Число не найдено", countResult)
	}

}

func generate(count int) (arr [n]int) {
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < n; i++ {
		arr[i] = generator.Intn(count)
	}
	return
}

func getCount(arr [n]int, value int) (count int) {
	count = 0
	for i := 0; i < n; i++ {
		if arr[i] == value {
			count = n - (i + 1)
			break
		}
	}

	return
}
