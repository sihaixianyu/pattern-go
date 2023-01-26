package fctymeth

import "fmt"

type Ak47 struct {
	gun
}

func newAk47() Gun {
	return &Ak47{
		gun: gun{
			name:  "AK47",
			power: 4,
		},
	}
}

func (g *Ak47) String() string {
	return fmt.Sprintf("name: %s, power: %d", g.name, g.power)
}
