package main

import (
	"flag"
	"math"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/go-echarts/go-echarts/charts"
	"github.com/sirupsen/logrus"
)

var (
	filename    = "output.html"
	numElements = 50
	src         = rand.NewSource(time.Now().Unix())
	rng         = rand.New(src)
	port        = ":8080"
	wg          sync.WaitGroup
)

func main() {
	flag.StringVar(&filename, "o", filename, "Output filename (html)")
	flag.IntVar(&numElements, "c", numElements, "Number of data points")
	flag.Parse()

	startWebServer()

	hostname, _ := os.Hostname()
	logrus.Infof("Listing on http://%v%v", hostname, port)
	wg.Wait()
}

func startWebServer() {
	wg.Add(1)
	go func() {

		http.HandleFunc("/", renderPage)
		http.ListenAndServe(port, nil)
		wg.Done()
	}()
}

func createChart1() *charts.Line {
	x1Values, y1Values := getData1()

	line := charts.NewLine()
	line.AddXAxis(x1Values)
	line.AddYAxis("Sin(x)", y1Values, charts.LineOpts{Smooth: true})
	line.Title = "My Chart 1"
	return line
}
func createChart2() *charts.Line {
	x1Values, y1Values := getData2()

	line := charts.NewLine()
	line.AddXAxis(x1Values)
	line.AddYAxis("Random Points 1", y1Values, charts.LineOpts{Smooth: true})

	_, y1Values = getData2()
	line.AddYAxis("Random Points 2", y1Values, charts.LineOpts{Step: true})

	line.Title = "My Chart 2"
	return line
}
func renderPage(w http.ResponseWriter, r *http.Request) {
	chart1 := createChart1()
	chart2 := createChart2()

	page := charts.NewPage()
	page.Add(chart1)
	page.Add(chart2)

	err := page.Render(w)
	if err != nil {
		logrus.Errorf("Unable to render graph: %v", err)
		return
	}
}

func getData2() (x, y []float64) {
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

func getData1() (x, y []float64) {
	x = make([]float64, numElements)
	for i := range x {
		x[i] = float64(i)
	}

	y = make([]float64, numElements)
	for i := range y {
		y[i] = 10 * math.Sin(x[i])
	}

	return
}
