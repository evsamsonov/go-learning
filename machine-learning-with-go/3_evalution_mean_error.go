package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"

	"gonum.org/v1/gonum/stat"
)

func main() {
	file, err := os.Open("data/continuous_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileReader := csv.NewReader(file)

	var observed []float64
	var predicted []float64

	var i int
	for {
		record, err := fileReader.Read()
		if err == io.EOF {
			break
		}
		if i == 0 {
			// Skip headers
			i++
			continue
		}

		observedVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}
		predictedVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)
		i++
	}

	var meanSquaredError, meanAbsoluteError float64
	var meanObserved float64
	for i, val := range observed {
		meanSquaredError += math.Pow(val-predicted[i], 2) / float64(len(observed))
		meanAbsoluteError += math.Abs(val-predicted[i]) / float64(len(observed))
		meanObserved += val / float64(len(observed))
	}

	rSquared := stat.RSquaredFrom(observed, predicted, nil)

	fmt.Printf("Mean observed: %.2f\n", meanObserved)
	fmt.Printf("Mean squared error (MSE): %.2f\n", meanSquaredError)
	fmt.Printf("Mean absolute error (MAE): %.2f\n", meanAbsoluteError)
	fmt.Printf("R ^ 2: %.2f\n", rSquared)
}
