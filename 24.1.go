package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 10

func main() {
	array := generate()
	fmt.Println("массив:", array)
	fmt.Println("сортирующую массив:", Sort(array))
}

func generate() (arr [n]int) {
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < n; i++ {
		arr[i] = generator.Intn(100)
	}

	return
}

func Sort(array [n]int) [n]int {

	for a := 0; a < (len(array)); a++ {

		for b := (a + 1); b < len(array); b++ {
			if array[a] > array[b] {
				c := array[a]
				array[a] = array[b]
				array[b] = c
			}
		}

	}

	return array
}
