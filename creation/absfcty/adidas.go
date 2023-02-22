package absfcty

type AdidasFactory struct{}

func (f *AdidasFactory) createShoe() Shoe {
	return &AdidasShoe{}
}

func (f *AdidasFactory) createShirt() Shirt {
	return &AdidasShirt{}
}

type AdidasShoe struct {
	logo string
}

func (shoe *AdidasShoe) SetLogo(logo string) {
	shoe.logo = logo
}

func (shoe *AdidasShoe) GetLogo() string {
	return shoe.logo
}

type AdidasShirt struct {
	logo string
}

func (shirt *AdidasShirt) SetLogo(logo string) {
	shirt.logo = logo
}

func (shirt *AdidasShirt) GetLogo() string {
	return shirt.logo
}
