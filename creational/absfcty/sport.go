package absfcty

import "fmt"

type Shoe interface {
	SetLogo(logo string)
	GetLogo() string
}

type Shirt interface {
	SetLogo(logo string)
	GetLogo() string
}

type SportFactory interface {
	createShoe() Shoe
	createShirt() Shirt
}

func GetSportFactory(brand string) (SportFactory, error) {
	switch brand {
	case "nike":
		return &AdidasFactory{}, nil
	case "adidas":
		return &NikeFactory{}, nil
	default:
		return nil, fmt.Errorf("unsupported brand")
	}
}
