package main

import (
	"flag"
	"os"

	"github.com/wcharczuk/go-chart/drawing"

	"github.com/sirupsen/logrus"
	chart "github.com/wcharczuk/go-chart" // expose package "chart"
)

var (
	filename = "output.png"
)

func main() {
	flag.StringVar(&filename, "o", filename, "Output filename (PNG)")
	flag.Parse()

	xValues := []float64{0.0, 1.0, 2.0, 3.0, 4.0, 5.0}
	yValues := []float64{1.0, 1.0, 2.0, 3.0, 5.0, 8.0}

	graph := chart.Chart{
		Title:      "Sample Chart",
		TitleStyle: getTitleStyle(),
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: xValues,
				YValues: yValues,
			},
		},
	}

	f, err := os.Create(filename)
	if err != nil {
		logrus.Errorf("Failed to create file: %v: %v", filename, err)
		return
	}

	defer f.Close()

	err = graph.Render(chart.PNG, f)
	if err != nil {
		logrus.Errorf("Unable to render graph: %v", err)
		return
	}
}

func getTitleStyle() chart.Style {
	return chart.Style{
		Show:      true,
		FontColor: drawing.ColorBlue,
	}
}
