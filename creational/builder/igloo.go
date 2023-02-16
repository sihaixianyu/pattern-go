package builder

type IglooBuilder struct {
	window string
	door   string
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (b *IglooBuilder) setWindow() {
	b.window = "snow window"
}

func (b *IglooBuilder) setDoor() {
	b.door = "snow door"
}

func (b *IglooBuilder) getHouse() House {
	return House{
		window: b.window,
		door:   b.door,
	}
}
