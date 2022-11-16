package main

import (
	"fmt"
	"math/rand"
	"time"
)

func isReturn(a int) bool {
	if a%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {

	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(100)
	result := isReturn(x)
	var text string
	if result {
		text = "четное"
	} else {
		text = "нечетное"
	}
	fmt.Printf("Число %d является %s\n", x, text)
}
