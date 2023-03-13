package bridge

import "fmt"

type Computer interface {
	Print()
	SetPrinter(Printer)
}

type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request from mac...")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(printer Printer) {
	m.printer = printer
}

type Windows struct {
	printer Printer
}

func (m *Windows) Print() {
	fmt.Println("Print request from windows...")
	m.printer.PrintFile()
}

func (m *Windows) SetPrinter(printer Printer) {
	m.printer = printer
}
