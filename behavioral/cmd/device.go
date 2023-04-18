package cmd

import "fmt"

type Device interface {
	On()
	Off()
}

type TV struct {
	isRunning bool
}

func (t *TV) On() {
	t.isRunning = true
	fmt.Println("TV is on!")
}

func (t *TV) Off() {
	t.isRunning = false
	fmt.Println("TV is off!")
}
