package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/montanaflynn/stats"
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
	sepalLength := irisDataframe.Col("sepal_length")

	fmt.Printf("series = %v\n\n", sepalLength.Float())

	meanVal := stat.Mean(sepalLength.Float(), nil)
	fmt.Printf("mean = %f\n", meanVal)

	modeVal, modeCount := stat.Mode(sepalLength.Float(), nil)
	fmt.Printf("mode = %f, mode count = %f\n", modeVal, modeCount)

	medianVal, err := stats.Median(sepalLength.Float())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("median = %f\n", medianVal)
}
