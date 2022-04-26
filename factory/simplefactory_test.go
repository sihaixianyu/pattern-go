package factory

import (
	"testing"
)

func TestNewAnimal(t *testing.T) {
	var a Animal

	a = NewAnimal(CatType)
	a.Eat()
	a.Yowl()

	a = NewAnimal(DogType)
	a.Eat()
	a.Yowl()

	a = NewAnimal(DuckType)
	a.Eat()
	a.Yowl()
}
