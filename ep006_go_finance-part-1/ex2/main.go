package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/piquette/finance-go/quote"
	"github.com/sirupsen/logrus"
)

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		logrus.Fatalf("No argument provided, exected one stock symbol. Example: %v CLDR", os.Args[0])
	}
	
	smbl := flag.Args()[0]

	q, _ := quote.Get(smbl)

	fmt.Printf("------- %v -------\n", q.ShortName)
	fmt.Printf("Current Price: $%v\n", q.Ask)
	fmt.Printf("52wk High: $%v\n", q.FiftyTwoWeekHigh)
	fmt.Printf("52wk Low: $%v\n", q.FiftyTwoWeekLow)
}
