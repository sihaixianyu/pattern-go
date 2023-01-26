package fctymeth

import "fmt"

type Musket struct {
	gun
}

func newMusket() Gun {
	return &Musket{
		gun: gun{
			name:  "Musket",
			power: 1,
		},
	}
}

func (g *Musket) String() string {
	return fmt.Sprintf("name: %v power: %v", g.name, g.power)
}
