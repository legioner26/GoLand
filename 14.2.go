package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Generate() (int, int) {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(10)
	y := rand.Intn(10)
	return x, y
}

func Transformation(x int, y int) (int, int) {
	x = 2*x + 10
	y = -3*y - 5
	return x, y
}

func main() {

	x1, y1 := Transformation(Generate())
	x2, y2 := Transformation(Generate())
	x3, y3 := Transformation(Generate())
	fmt.Printf("Координаты первой точки: %d %d\n", x1, y1)
	fmt.Printf("Координаты второй точки: %d %d\n", x2, y2)
	fmt.Printf("Координаты третьей точки: %d %d\n", x3, y3)

}
