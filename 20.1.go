package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func determinant(matrix mat.Matrix) {
	//Определитель матрицы
	determinant := mat.Det(matrix)
	fmt.Println(determinant)
}
func main() {
	row1 := []float64{1, 2, 5}
	row2 := []float64{7, 3, 4}
	row3 := []float64{6, 9, 0}

	matrix := mat.NewDense(3, 3, nil)
	matrix.SetRow(0, row1)
	matrix.SetRow(1, row2)
	matrix.SetRow(2, row3)
	fmt.Printf("A matrix:\n%v\n\n", mat.Formatted(matrix, mat.Prefix(""), mat.Excerpt(0)))
	determinant(matrix)

}
