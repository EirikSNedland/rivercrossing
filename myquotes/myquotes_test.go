package myquotes

import (
	"testing"
)

func TestMyQuote(t *testing.T) {
	wanted := string("Hello there")
	state := Hello()
	if state != wanted {
		t.Errorf("Feil, fikk %v, ønsket %v", state, wanted)
	}
}
