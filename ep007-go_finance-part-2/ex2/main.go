package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/go-echarts/go-echarts/charts"
	"github.com/piquette/finance-go/datetime"

	"github.com/leekchan/accounting"
	"github.com/piquette/finance-go/chart"
	"github.com/sirupsen/logrus"
)

const (
	addr = ":8080"
)

var (
	src = rand.NewSource(time.Now().Unix())
	rng = rand.New(src)
)

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		logrus.Fatalf("No argument provided, exected at least one stock symbol. Example: %v cldr goog aapl intc amd ...", os.Args[0])
	}

	startWebServer()
}

func startWebServer() {
	http.HandleFunc("/", graphHandler)
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		logrus.Fatalf("Unable to start critical service: %v", err)
	}
}

func graphHandler(w http.ResponseWriter, r *http.Request) {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: "Stocks"})

	x, y := getData()

	line.AddXAxis(x)
	line.AddYAxis("APPL", y)

	line.Render(w)
}

func getData() ([]string, []float64) {
	count := 356
	x := make([]string, count)
	y := make([]float64, count)

	t := time.Now()

	for i := 0; i < count; i++ {
		x[i] = t.Format("2006-01-02")
		y[i] = rng.Float64()*600.0 - 300.0
		t = t.AddDate(0, 0, -1)
	}

	return x, y
}

func foo() {
	cf := accounting.Accounting{Symbol: "$", Precision: 2}
	smbls := flag.Args()
	_ = cf
	_ = smbls

	now := time.Now()
	yearAgo := now.AddDate(-1, 0, 0)

	p := &chart.Params{
		Symbol:   "aapl",
		Start:    datetime.New(&yearAgo),
		End:      datetime.New(&now),
		Interval: datetime.OneDay,
	}

	iter := chart.Get(p)
	count := iter.Count()

	fmt.Printf("We got %v data point\n", count)

	for iter.Next() {
		d := iter.Bar()
		price, _ := d.Close.Round(2).Float64()
		date := time.Unix(int64(d.Timestamp), 0).Format("2006-01-02")

		fmt.Printf("%v - $%v\n", date, price)
	}
}
