package bridge

import (
	"fmt"
	"testing"
)

func TestCase1(t *testing.T) {
	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}

	macComputer := &Mac{}
	
	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &Windows{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}
