package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Введите два числа")

	var (
		a int16
		b int16
	)
	fmt.Scan(&a)
	fmt.Scan(&b)

	result := int32(a) * int32(b)

	switch {

	case result > 0 && result < math.MaxUint8:
		fmt.Printf("%d  %d результат Uint8 \n", a, b)
	case result > math.MinInt8 && result < math.MaxInt8:
		fmt.Printf("%d  %d результат Int8 \n", a, b)
	case result > 0 && result < math.MaxUint16:
		fmt.Printf("%d  %d результат Uint16 \n", a, b)
	case result > math.MinInt16 && result < math.MaxInt16:
		fmt.Printf("%d  %d результат Int16 \n", a, b)
	case result > 0 && uint32(result) < math.MaxUint32:
		fmt.Printf("%d  %d результат Uint32 \n", a, b)
	case result > math.MinInt32 && result < math.MaxInt32:
		fmt.Printf("%d  %d результат Int32 \n", a, b)
	}
}
