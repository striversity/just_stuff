package main

import (
	"fmt"

	"github.com/piquette/finance-go/quote"
)

func main() {
	smbl := "cldr"

	q, _ := quote.Get(smbl)

	fmt.Printf("------- %v -------\n", q.ShortName)
	fmt.Printf("Current Price: $%v\n", q.Ask)
	fmt.Printf("52wk High: $%v\n", q.FiftyTwoWeekHigh)
	fmt.Printf("52wk Low: $%v\n", q.FiftyTwoWeekLow)
}
