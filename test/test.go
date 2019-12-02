package main

import (
	"fmt"

	Matrix "github.com/9init/matrix"
)

func main() {

	//first matrix
	m1 := Matrix.NewMatrix(2, 3)
	col1m1 := []float64{1, 2, 3}
	col2m1 := []float64{4, 5, 6}
	m1.Matrix[0] = col1m1
	m1.Matrix[1] = col2m1

	//second matrix
	m2 := Matrix.NewMatrix(3, 2)
	col1m2 := []float64{7, 8}
	col2m2 := []float64{9, 10}
	col3m2 := []float64{11, 12}
	m2.Matrix[0] = col1m2
	m2.Matrix[1] = col2m2
	m2.Matrix[2] = col3m2

	//dot product

	m3, err := m1.DotProduct(m2)
	if err != nil {
		panic(err)
	}
	fmt.Println(m3.Matrix)

}
