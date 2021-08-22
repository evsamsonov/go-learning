package main

import (
	"fmt"
	"github.com/kniren/gota/dataframe"
	"log"
	"os"
)

func main() {
	irisFile, err := os.Open("machine-learning-with-go/data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	irisDF := dataframe.ReadCSV(irisFile)
	fmt.Println(irisDF)

	// Фильтрация по species
	filteredIrisDF := irisDF.Filter(dataframe.F{
		Colname:    "species",
		Comparator: "==",
		Comparando: "Iris-versicolor",
	})

	// Выбор только определенных полей
	fmt.Println(filteredIrisDF.Select([]string{"sepal_width", "species"}))

	// Только 3 первых ряда
	fmt.Println(filteredIrisDF.Select([]string{"sepal_width", "species"}).Subset([]int{0, 1, 2}))
}
