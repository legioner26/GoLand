package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 12

func main() {
	var num int

	array := generate(5)
	fmt.Println("массив", array)

	fmt.Println("Введите число")
	fmt.Scan(&num)

	firstIndex := getFirstIndex(array, num)
	fmt.Println("Индекс:", firstIndex)
}
func generate(count int) (arr [n]int) {
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < n; i++ {
		arr[i] = generator.Intn(count)
	}
	return
}
func getFirstIndex(arr [n]int, value int) (index int) {
	index = -1
	min := 0
	max := n - 1

	for max >= min {
		middle := (max + min) / 2

		if arr[middle] == value {
			index = middle
			i := 0

			for arr[middle-i] == value {
				index = middle - i
				if i == middle {
					break
				} else {
					i++
				}
			}

			break
		} else if arr[middle] > value {
			max = middle - 1
		} else {
			min = middle + 1
		}
	}

	return
}
