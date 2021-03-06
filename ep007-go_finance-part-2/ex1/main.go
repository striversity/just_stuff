package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/piquette/finance-go/datetime"

	"github.com/leekchan/accounting"
	"github.com/piquette/finance-go/chart"
	"github.com/sirupsen/logrus"
)

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		logrus.Fatalf("No argument provided, exected at least one stock symbol. Example: %v cldr goog aapl intc amd ...", os.Args[0])
	}

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
