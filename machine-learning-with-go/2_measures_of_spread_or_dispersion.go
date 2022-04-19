package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"
	"log"
	"os"
)

func main() {
	irisFile, err := os.Open("machine-learning-with-go/data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	irisDataframe := dataframe.ReadCSV(irisFile)
	petalLengthCol := irisDataframe.Col("petal_length")

	fmt.Printf("series = %v\n\n", petalLengthCol.Float())

	fmt.Printf("max = %0.2f\n", petalLengthCol.Max())
	fmt.Printf("min = %0.2f\n", petalLengthCol.Min())
	fmt.Printf("range = %0.2f\n", petalLengthCol.Max()-petalLengthCol.Min())
	fmt.Printf("variance = %0.2f\n", stat.Variance(petalLengthCol.Float(), nil))
	fmt.Printf("standard deviation = %0.2f\n\n", stat.StdDev(petalLengthCol.Float(), nil))

	petalLengths := petalLengthCol.Float()
	inds := make([]int, len(petalLengthCol.Float())) // Сохраняем порядок элементов
	floats.Argsort(petalLengths, inds)

	fmt.Printf("quantile 0.25 = %0.2f\n", stat.Quantile(0.25, stat.Empirical, petalLengths, nil))
	fmt.Printf("quantile 0.50 (median) = %0.2f\n", stat.Quantile(0.50, stat.Empirical, petalLengths, nil))
	fmt.Printf("quantile 0.75 = %0.2f\n", stat.Quantile(0.75, stat.Empirical, petalLengths, nil))
}
