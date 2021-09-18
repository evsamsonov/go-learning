package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	irisFile, err := os.Open("machine-learning-with-go/data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	irisDataframe := dataframe.ReadCSV(irisFile)

	for _, colName := range irisDataframe.Names() {
		if colName == "species" {
			continue
		}

		vals := make(plotter.Values, irisDataframe.Nrow())
		for i, v := range irisDataframe.Col(colName).Float() {
			vals[i] = v
		}

		irisPlot := plot.New()
		irisPlot.Title.Text = fmt.Sprintf("Histogram of a %s", colName)

		hist, err := plotter.NewHist(vals, 16)
		if err != nil {
			log.Fatal(hist)
		}
		hist.Normalize(1)
		irisPlot.Add(hist)

		if err := irisPlot.Save(4*vg.Inch, 4*vg.Inch, "machine-learning-with-go/2_histogram/"+colName+"_hist.png"); err != nil {
			log.Fatal(colName, err)
		}
	}
}
