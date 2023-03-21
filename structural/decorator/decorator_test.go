package decorator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecorator(t *testing.T) {
	pizza := &VeggiePizza{}

	pizzaWithTomato := &TomatoTopping{
		pizza: pizza,
	}

	pizzaWithCheeseAndTomato := &CheeseTopping{
		pizza: pizzaWithTomato,
	}

	assert.Equal(t, 32, pizzaWithCheeseAndTomato.getPrice())
}
