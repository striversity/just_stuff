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
		logrus.Fatalf("No argument provided, exected one stock symbol. Example: %v CLDR", os.Args[0])
	}

	cf := accounting.Accounting{Symbol: "$", Precision: 2}
	smbl := flag.Args()[0]

	q, _ := quote.Get(smbl)

	fmt.Printf("------- %v -------\n", q.ShortName)
	fmt.Printf("Current Price: %v\n", cf.FormatMoney(q.Ask))
	fmt.Printf("52wk High: %v\n", cf.FormatMoney(q.FiftyTwoWeekHigh))
	fmt.Printf("52wk Low: %v\n", cf.FormatMoney(q.FiftyTwoWeekLow))
}
