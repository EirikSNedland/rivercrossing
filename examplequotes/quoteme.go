package examplequotes

import (
	"fmt"
	"rsc.io/quote"
)

func PrintExQuotes() {
	fmt.Println(quote.Hello())
	fmt.Println(quote.Go())
	fmt.Println(quote.Opt())
	fmt.Println(quote.Glass())
}
