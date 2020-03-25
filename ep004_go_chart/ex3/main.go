package main

import (
	"flag"
	"math/rand"
	"os"
	"time"

	"github.com/wcharczuk/go-chart/drawing"

	"github.com/sirupsen/logrus"
	chart "github.com/wcharczuk/go-chart" // expose package "chart"
)

var (
	filename    = "output.png"
	numElements = 20
	src         = rand.NewSource(time.Now().Unix())
	rng         = rand.New(src)
)

func main() {
	flag.StringVar(&filename, "o", filename, "Output filename (PNG)")
	flag.IntVar(&numElements, "c", numElements, "Number of data points")
	flag.Parse()

	xValues, yValues := getData()

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

func getData() (x, y []float64) {
	x = make([]float64, numElements)
	for i := range x {
		x[i] = float64(i)
	}

	y = make([]float64, numElements)
	for i := range y {
		y[i] = 10 * rng.Float64()
	}

	return
}
