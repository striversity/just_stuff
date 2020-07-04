package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/leekchan/accounting"
	"github.com/piquette/finance-go/quote"
	"github.com/sirupsen/logrus"
)

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		logrus.Fatalf("No argument provided, exected at least one stock symbol. Example: %v cldr goog aapl intc amd ...", os.Args[0])
	}

	cf := accounting.Accounting{Symbol: "$", Precision: 2}
	smbls := flag.Args()

	iter := quote.List(smbls)

	for iter.Next() {
		q := iter.Quote()
		fmt.Printf("------- %v -------\n", q.ShortName)
		fmt.Printf("Current Price: %v\n", cf.FormatMoney(q.Ask))
		fmt.Printf("52wk High: %v\n", cf.FormatMoney(q.FiftyTwoWeekHigh))
		fmt.Printf("52wk Low: %v\n", cf.FormatMoney(q.FiftyTwoWeekLow))
		fmt.Printf("-----------------\n")
	}
}
