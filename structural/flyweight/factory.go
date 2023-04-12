package flyweight

import "fmt"

const (
	TerroristDressType        = "tDress"
	CounterTerroristDressType = "ctDress"
)

var dressFactorySingleton = &DressFactory{
	dressMap: make(map[string]Dress),
}

type DressFactory struct {
	dressMap map[string]Dress
}

func GetDressFactorySingleton() *DressFactory {
	return dressFactorySingleton
}

func (d *DressFactory) getDressByType(dressType string) (Dress, error) {
	if dress, ok := d.dressMap[dressType]; ok {
		return dress, nil
	}

	if dressType == TerroristDressType {
		dress := NewTerroristDress()
		d.dressMap[TerroristDressType] = dress
		return dress, nil
	}

	if dressType == CounterTerroristDressType {
		dress := NewCounterTerroristDress()
		d.dressMap[CounterTerroristDressType] = dress
		return dress, nil
	}

	return nil, fmt.Errorf("Unknown DressType %v", dressType)
}
