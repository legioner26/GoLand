package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 10

func main() {
	array := generate()
	fmt.Println("Array:", array)

	even, uneven := searchEU(array)
	fmt.Println("Четных:", even)
	fmt.Println("Нечетных:", uneven)
}

func generate() (arr [n]int) {
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < n; i++ {
		arr[i] = generator.Intn(100)
	}

	return
}

func searchEU(arr [n]int) (even int, uneven int) {

	for i := 0; i < n; i++ {
		if arr[i]%2 == 0 {
			even++
		} else {
			uneven++
		}
	}

	return
}
