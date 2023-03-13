package bridge

import "fmt"

type Printer interface {
	PrintFile()
}

type Epson struct{}

func (e *Epson) PrintFile() {
	fmt.Println("Print by a EPSON printer...")
}

type Hp struct{}

func (p *Hp) PrintFile() {
	fmt.Println("Printing by a HP Printer...")
}
