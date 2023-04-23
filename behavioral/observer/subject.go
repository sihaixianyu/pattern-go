package observer

import "fmt"

type Subject interface {
	register(Observer)
	deregister(Observer)
	notifyAll()
}

type Item struct {
	observers map[string]Observer
	name      string
	inStock   bool
}

func NewItem(name string) *Item {
	return &Item{
		name:      name,
		observers: make(map[string]Observer),
	}
}

func (i *Item) Register(o Observer) {
	i.observers[o.getId()] = o
}

func (i *Item) Deregister(o Observer) {
	delete(i.observers, o.getId())
}

func (i *Item) UpdateAvailability() {
	fmt.Println("I'm in stock, you guys can buy me now!")
	i.inStock = true
	i.notifyAll()
}

func (i *Item) notifyAll() {
	for _, v := range i.observers {
		v.update(i.name)
	}
}
