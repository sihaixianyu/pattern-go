package mediator

import (
	"testing"
)

func TestMediator(t *testing.T) {
	manager := NewStationManager()

	passengerTrain := &PassengerTrain{
		mediator: &manager,
	}
	freightTrain := &FreightTrain{
		mediator: &manager,
	}

	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()
}
