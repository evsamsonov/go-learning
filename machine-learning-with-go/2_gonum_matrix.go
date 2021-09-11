package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
)

func main() {
	components := []float64{1.2, -5.7, -2.4, 7.3}
	matrix := mat.NewDense(2, 2, components)
	formattedMatrix := mat.Formatted(matrix, mat.Prefix(" "))
	fmt.Printf("Matrix:\n %v\n\n", formattedMatrix)

	fmt.Printf("Second value: %v\n\n", matrix.At(0, 1))

	fmt.Printf("First column: %v\n\n", mat.Col(nil, 0, matrix))

	fmt.Printf("Second row: %v\n\n", mat.Row(nil, 0, matrix))

	matrix.Set(0, 1, 11.2)
	fmt.Printf("Matrix:\n %v\n\n", mat.Formatted(matrix, mat.Prefix(" ")))

	matrix.SetCol(1, []float64{10.1, 0.13})
	fmt.Printf("Matrix:\n %v\n\n", mat.Formatted(matrix, mat.Prefix(" ")))

	// Matrix operations
	a := mat.NewDense(3, 3, []float64{
		1, 2, 3,
		0, 4, 5,
		0, 0, 6,
	})
	b := mat.NewDense(3, 3, []float64{
		8, 9, 10,
		1, 4, 2,
		9, 0, 2,
	})

	d := mat.NewDense(3, 3, nil)
	d.Add(a, b)
	fd := mat.Formatted(d, mat.Prefix(" "))
	fmt.Printf("a + b =\n %v\n\n", fd)

	c := mat.NewDense(3, 2, []float64{
		3, 2,
		1, 4,
		0, 8,
	})

	f := mat.NewDense(3, 2, nil)
	f.Mul(a, c)
	ff := mat.Formatted(f, mat.Prefix(" "))
	fmt.Printf("a * c =\n %v\n\n", ff)

	g := mat.NewDense(3, 3, nil)
	g.Pow(a, 5)
	fg := mat.Formatted(g, mat.Prefix(" "))
	fmt.Printf("a ^ 5 =\n %v\n\n", fg)

	h := mat.NewDense(3, 3, nil)
	h.Apply(func(_, _ int, v float64) float64 {
		return math.Pow(v, 2)
	}, a)
	fh := mat.Formatted(h, mat.Prefix(" "))
	fmt.Printf("pow each value =\n %v\n\n", fh)
}
