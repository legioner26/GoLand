package main

import (
	"fmt"
	"math"
)

func main() {

	var (
		a           uint8
		b           uint16
		countUint8  int
		countUint16 int
	)

	fmt.Println("Количество переполнений от 0 до MaxUint32")

	for i := 0; i < math.MaxUint32; i++ {
		if a == math.MaxUint8 {
			countUint8++
		}
		if b == math.MaxUint16 {
			countUint16++
		}
		a++
		b++
	}
	fmt.Printf("Uint8 переполнился %d раз \n", countUint8)
	fmt.Printf("Uint16 переполнился %d раз \n", countUint16)
}
