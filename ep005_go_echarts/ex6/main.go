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
func createChart3() *charts.Line {
	x1Values, y1Values := getData2()

	line := charts.NewLine()
	line.AddXAxis(x1Values)
	line.AddYAxis("Random Points 1", y1Values, charts.LineOpts{Smooth: true})

	line.Title = "Chart 3"
	line.SetSeriesOptions(
		charts.MLNameTypeItem{Name: "Average", Type: "average"},
		charts.MLStyleOpts{Label: charts.LabelTextOpts{Show: true, Formatter: "{a}: {b}"}},
	)
	return line
}
func createBar() *charts.Bar {
	months := []string{"Oct 2019", "Nov 2019", "Dec 2019", "Jan 2020", "Feb 2020", "Mar 2020"}
	numItems := len(months)
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "6mnts Sales"}, charts.ToolboxOpts{Show: true})
	bar.AddXAxis(months)
	bar.AddYAxis("Estimated", randInt(numItems))
	bar.AddYAxis("Actual", randInt(numItems))
	return bar
}
func randInt(c int) []int {
	r := make([]int, c)
	for i := 0; i < c; i++ {
		r[i] = rng.Intn(50)
	}
	return r
}

func renderPage(w http.ResponseWriter, r *http.Request) {
	chart1 := createChart1()
	chart2 := createChart2()
	chart3 := createChart3()
	bar := createBar()

	page := charts.NewPage()
	page.Add(chart1)
	page.Add(chart2)
	page.Add(chart3)
	page.Add(bar)

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
