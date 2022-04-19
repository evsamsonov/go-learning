package main

import (
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
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

	irisPlot := plot.New()
	irisPlot.Title.Text = "Box plots"
	irisPlot.Y.Label.Text = "Values"

	for i, colName := range irisDataframe.Names() {
		if colName == "species" {
			continue
		}

		vals := make(plotter.Values, irisDataframe.Nrow())
		for i, v := range irisDataframe.Col(colName).Float() {
			vals[i] = v
		}

		boxPlot, err := plotter.NewBoxPlot(vg.Points(50), float64(i), vals)
		if err != nil {
			log.Fatal(err)
		}
		irisPlot.Add(boxPlot)
	}

	irisPlot.NominalX("sepal_length", "sepal_width", "petal_length", "petal_width")

	if err := irisPlot.Save(6*vg.Inch, 8*vg.Inch, "machine-learning-with-go/2_box_plot/boxplots.png"); err != nil {
		log.Fatal(err)
	}
}
