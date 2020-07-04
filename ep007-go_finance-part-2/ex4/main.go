package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-echarts/go-echarts/charts"
	"github.com/piquette/finance-go/datetime"

	"github.com/piquette/finance-go/chart"
	"github.com/sirupsen/logrus"
)

const (
	addr = ":8080"
)

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		logrus.Fatalf("No argument provided, exected at least one stock symbol. Example: %v cldr goog aapl intc amd ...", os.Args[0])
	}
	smbls := flag.Args()
	_ = smbls

	startWebServer()
}

func startWebServer() {
	http.Handle("/", &graphHandler{"aapl"})

	err := http.ListenAndServe(addr, nil)

	if err != nil {
		logrus.Fatalf("Unable to start critical service: %v", err)
	}
}

type graphHandler struct {
	sym string
}

func (gh *graphHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: "Stocks"})

	x, y := getData(gh.sym)

	line.AddXAxis(x)
	line.AddYAxis(gh.sym, y)

	line.Render(w)
}

func getData(sym string) ([]string, []float64) {
	now := time.Now()
	yearAgo := now.AddDate(-1, 0, 0)

	p := &chart.Params{
		Symbol:   sym,
		Start:    datetime.New(&yearAgo),
		End:      datetime.New(&now),
		Interval: datetime.OneDay,
	}

	iter := chart.Get(p)
	count := iter.Count()

	fmt.Printf("We got %v data point\n", count)

	x := make([]string, count)
	y := make([]float64, count)

	var date string
	var price float64

	i := 0
	for iter.Next() {
		d := iter.Bar()
		price, _ = d.Close.Round(2).Float64()
		date = time.Unix(int64(d.Timestamp), 0).Format("2006-01-02")

		x[i] = date
		y[i] = price
		i++
	}

	return x, y
}
