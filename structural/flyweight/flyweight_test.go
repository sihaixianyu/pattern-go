package flyweight

import (
	"fmt"
	"testing"
)

func TestFlyweight(t *testing.T) {
	game := NewGame()

	//Add Terrorist
	game.AddTerrorist(TerroristDressType)
	game.AddTerrorist(TerroristDressType)
	game.AddTerrorist(TerroristDressType)
	game.AddTerrorist(TerroristDressType)

	//Add CounterTerrorist
	game.AddCounterTerrorist(CounterTerroristDressType)
	game.AddCounterTerrorist(CounterTerroristDressType)
	game.AddCounterTerrorist(CounterTerroristDressType)

	for dressType, dress := range GetDressFactorySingleton().dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.GetColor())
	}
}
