package factory

import "fmt"

type Animal interface {
	Eat()
	Yowl()
}

type AnimalType int

const (
	CatType AnimalType = iota
	DogType
	DuckType
)

func NewAnimal(t AnimalType) Animal {
	switch t {
	case CatType:
		return newCat()
	case DogType:
		return newDog()
	default:
		return newDuck()
	}
}

type Cat struct{}

func newCat() Cat {
	return Cat{}
}

func (c Cat) Eat() {
	fmt.Println("I'm a cat, I eat cat food!")
}

func (c Cat) Yowl() {
	fmt.Println("Mew! Mew! Mew!")
}

type Dog struct{}

func newDog() Dog {
	return Dog{}
}

func (d Dog) Eat() {
	fmt.Println("I'm a dog, I eat anything!")
}

func (d Dog) Yowl() {
	fmt.Println("Woof! Woof! Woof!")
}

type Duck struct{}

func newDuck() Duck {
	return Duck{}
}

func (d Duck) Eat() {
	fmt.Println("I'm a duck, I eat vegetable!")
}

func (d Duck) Yowl() {
	fmt.Println("Quack! Quack! Quack!")
}

