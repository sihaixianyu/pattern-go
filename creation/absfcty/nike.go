package absfcty

type NikeFactory struct{}

func (f *NikeFactory) createShoe() Shoe {
	return &NikeShoe{}
}

func (f *NikeFactory) createShirt() Shirt {
	return &NikeShirt{}
}

type NikeShoe struct {
	logo string
}

func (shoe *NikeShoe) SetLogo(logo string) {
	shoe.logo = logo
}

func (shoe *NikeShoe) GetLogo() string {
	return shoe.logo
}

type NikeShirt struct {
	logo string
}

func (shirt *NikeShirt) SetLogo(logo string) {
	shirt.logo = logo
}

func (shirt *NikeShirt) GetLogo() string {
	return shirt.logo
}
