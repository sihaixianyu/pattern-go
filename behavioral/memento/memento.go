package memento

type State string

type memento struct {
	state State
}

func (m *memento) getSvaedMemento() State {
	return m.state
}

type Originator struct {
	state State
}

func (o *Originator) CreateMemento() *memento {
	return &memento{
		state: o.state,
	}
}

func (o *Originator) RestoreMemento(m *memento) {
	o.state = m.getSvaedMemento()
}

func (o *Originator) State() State {
	return o.state
}

func (o *Originator) SetState(s State) {
	o.state = s
}

type Caretaker struct {
	mementos []*memento
}

func (c *Caretaker) AddMemento(m *memento) {
	c.mementos = append(c.mementos, m)
}

func (c *Caretaker) GetByIdx(idx int) *memento {
	return c.mementos[idx]
}
