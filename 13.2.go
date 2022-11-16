package main

import (
	"fmt"
)

func changePlace(a, b *int) {
	c := *a
	*a = *b
	*b = c
}

func main() {

	a := 10
	b := 5
	fmt.Println(a, b)
	changePlace(&a, &b)
	fmt.Println(a, b)
}
