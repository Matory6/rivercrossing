package Rivercrossing

import "testing"

// Test funksjon
func TestViewState(t *testing.T) {
	iWantThis := "---\\ \\_Fox_/ ____________________/---"
	testThis := ViewState()
	if testThis != iWantThis {
		t.Errorf("Feil, fikk %q, ønsket %q.", testThis, iWantThis)
	}
}