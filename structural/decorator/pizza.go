package decorator

type Pizza interface {
	getPrice() int
}

type VeggiePizza struct {
}

func (p *VeggiePizza) getPrice() int {
	return 15
}
