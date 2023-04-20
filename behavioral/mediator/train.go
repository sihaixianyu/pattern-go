package mediator

import "fmt"

type Train interface {
	arrive()
	depart()
	permitArrival()
}

type PassengerTrain struct {
	mediator Mediator
}

func (g *PassengerTrain) arrive() {
	if !g.mediator.canArrive(g) {
		fmt.Println("PassengerTrain: Arrival blocked , waiting...")
		return
	}

	fmt.Println("PassengerTrain: Arrived.")
}

func (g *PassengerTrain) depart() {
	fmt.Println("PassengerTrain: Leaving...")
	g.mediator.notifyDeparture()
}

func (g *PassengerTrain) permitArrival() {
	fmt.Println("PassengerTrain: Arrival permmited, arriving...")
	g.arrive()
}

type FreightTrain struct {
	mediator Mediator
}

func (g *FreightTrain) arrive() {
	if !g.mediator.canArrive(g) {
		fmt.Println("FreightTrain: Arrival blocked , waiting...")
		return
	}

	fmt.Println("FreightTrain: Arrived.")
}

func (g *FreightTrain) depart() {
	fmt.Println("FreightTrain: Leaving...")
	g.mediator.notifyDeparture()
}

func (g *FreightTrain) permitArrival() {
	fmt.Println("FreightTrain: Arrival permmited, arriving...")
	g.arrive()
}
