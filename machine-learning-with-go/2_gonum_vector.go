package main

import (
	"fmt"
	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

func main() {
	// 2 варианта работы с векторами
	// floats удобнее использовать, когда требуются только вектора
	vector1 := []float64{0.1, 0.35, 0.12}
	vector2 := []float64{1.1, 1.02, 0.01}

	dotProduct := floats.Dot(vector1, vector2)
	fmt.Printf("Dot product: %0.2f\n", dotProduct)

	floats.Scale(2., vector1)
	fmt.Printf("Vector1 after scale: %+v\n", vector1)

	norm := floats.Norm(vector2, 2)
	fmt.Printf("norm: %0.2f\n", norm)

	// Пакет mat и blas64, когда требуется работа с матрицами
	vector3 := mat.NewVecDense(3, []float64{0.1, 0.35, 0.12})
	vector4 := mat.NewVecDense(3, []float64{1.1, 1.02, 0.01})

	dotProduct2 := mat.Dot(vector3, vector4)
	fmt.Printf("Dot product: %0.2f\n", dotProduct2)

	vector3.ScaleVec(2, vector3)
	fmt.Printf("Vector1 after scale: %+v\n", vector3)

	norm2 := blas64.Nrm2(vector4.RawVector())
	fmt.Printf("norm: %0.2f\n", norm2)
}
