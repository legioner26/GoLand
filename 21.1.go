package main

import (
	"fmt"
)

func matFunc(x int16, y uint8, z float32) float32 {
	return 2*float32(x) + float32(y)*float32(y) - 3/z
}

func main() {
	fmt.Println(matFunc(1, 2, 3))
}
