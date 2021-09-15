package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"log"
)

func main() {
	a := mat.NewDense(3, 3, []float64{
		1, 2, 3,
		0, 4, 5,
		0, 0, 6,
	})

	ft := mat.Formatted(a.T(), mat.Prefix(" "))
	fmt.Printf("a^T = \n%+v\n\n", ft)

	deta := mat.Det(a)
	fmt.Printf("det(a) = \n%.2f\n\n", deta)

	inversedA := mat.NewDense(3, 3, nil)
	if err := inversedA.Inverse(a); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("inversedA = \n%v\n\n", mat.Formatted(inversedA, mat.Prefix(" ")))
}
