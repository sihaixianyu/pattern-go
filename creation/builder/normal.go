package builder

type NormalBuilder struct {
	window string
	door   string
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) setWindow() {
	b.window = "wooden window"
}

func (b *NormalBuilder) setDoor() {
	b.door = "wooden door"
}

func (b *NormalBuilder) getHouse() House {
	return House{
		window: b.window,
		door:   b.door,
	}
}
