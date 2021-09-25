package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat/distuv"

	"gonum.org/v1/gonum/stat"
)

func main() {
	// Данные после некоторых изменений (альтернативная гипотеза)
	observed := []float64{
		260, // Не занимаются регулярно спортом
		135, // Периодически
		105, // Регулярно
	}
	totalObserved := 500.

	// Данные до некоторых изменений (нулевая гипотеза)
	expected := []float64{
		totalObserved * 0.60,
		totalObserved * 0.25,
		totalObserved * 0.15,
	}

	chiSquare := stat.ChiSquare(observed, expected)
	fmt.Printf("Chi-square: %f\n", chiSquare)

	// K - кол-во возможных категорий - 1
	chiDist := distuv.ChiSquared{K: 2, Src: nil}
	pValue := chiDist.Prob(chiSquare)

	// p-value (0.000058) показывает, вероятность того, что отклонение было получено чисто случайно
	// Если мы используем порог в 5% (обычно), то можем отклонить нулевую гипотезу в принять альтернативную
	// то есть на результат повлияли некоторые изменения
	fmt.Printf("P-value: %f", pValue)
}
