package memento

import (
	"fmt"
	"testing"
)

func TestMemento(t *testing.T) {
	caretaker := &Caretaker{
		mementos: make([]*memento, 0),
	}

	originator := &Originator{
		state: "A",
	}

	fmt.Printf("Originator Current State: %s\n", originator.State())
	caretaker.AddMemento(originator.CreateMemento())

	originator.SetState("B")
	fmt.Printf("Originator Current State: %s\n", originator.State())
	caretaker.AddMemento(originator.CreateMemento())

	originator.SetState("C")
	fmt.Printf("Originator Current State: %s\n", originator.State())
	caretaker.AddMemento(originator.CreateMemento())

	originator.RestoreMemento(caretaker.GetByIdx(1))
	fmt.Printf("Restored to State: %s\n", originator.State())

	originator.RestoreMemento(caretaker.GetByIdx(0))
	fmt.Printf("Restored to State: %s\n", originator.State())
}
