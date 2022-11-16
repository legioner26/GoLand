package main

import (
	"fmt"
)

func sum(a, b int) {
	if a < b {
		c := a
		a = b
		b = c
	}
	var sum int

	for i := b; i < a; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
func main() {

	sum(5, 9)
	sum(11, 4)
}
