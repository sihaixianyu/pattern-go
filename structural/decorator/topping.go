package decorator

type TomatoTopping struct {
	pizza Pizza
}

func (tt *TomatoTopping) getPrice() int {
	return tt.pizza.getPrice() + 7
}

type CheeseTopping struct {
	pizza Pizza
}

func (ct *CheeseTopping) getPrice() int {
	return ct.pizza.getPrice() + 10
}
