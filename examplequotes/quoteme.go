package examplequote

import (
	"fmt"
	"rsc.io/quote"
)

func PrintExQuote() {
	fmt.Println(quote.Hello())
	fmt.Println(quote.Go())
	fmt.Println(quote.Opt())
	fmt.Println(quote.Glass())
}
