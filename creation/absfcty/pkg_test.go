package absfcty

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	adidasFactory, _ := GetSportFactory("adidas")
	nikeFactory, _ := GetSportFactory("nike")

	nikeShoe := nikeFactory.createShoe()
	nikeShirt := nikeFactory.createShirt()

	nikeShoe.SetLogo("Nike Shoe")
	nikeShirt.SetLogo("Nike Shirt")

	assert.Equal(t, nikeShoe.GetLogo(), "Nike Shoe")
	assert.Equal(t, nikeShirt.GetLogo(), "Nike Shirt")

	adidasShoe := adidasFactory.createShoe()
	adidasShirt := adidasFactory.createShirt()

	adidasShoe.SetLogo("Adidas Shoe")
	adidasShirt.SetLogo("Adidas Shirt")

	assert.Equal(t, adidasShoe.GetLogo(), "Adidas Shoe")
	assert.Equal(t, adidasShirt.GetLogo(), "Adidas Shirt")
}
