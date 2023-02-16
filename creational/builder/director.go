package builder

type Director struct {
	builder HouseBuilder
}

func NewDirector(builder HouseBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) SetBuilder(builder HouseBuilder) {
	d.builder = builder
}

func (d *Director) BuildHouse() House {
	d.builder.setWindow()
	d.builder.setDoor()

	return d.builder.getHouse()
}
