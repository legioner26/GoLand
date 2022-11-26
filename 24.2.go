package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 10

func main() {
	array := generate()
	fmt.Println("Array", array)

	SortReverse := func(arr [n]int) (output [n]int) {
		inputLen := len(array)
		for a := 0; a < (inputLen); a++ {

			for b := (a + 1); b < inputLen; b++ {
				if array[a] > array[b] {
					c := array[a]
					array[a] = array[b]
					array[b] = c
				}
			}
			d := inputLen - a - 1

			output[d] = array[a]
		}

		return
	}
	fmt.Println("Сортировка и реверс", SortReverse(array))

}

func generate() (arr [n]int) {
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < n; i++ {
		arr[i] = generator.Intn(100)
	}

	return
}
