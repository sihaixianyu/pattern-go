package flyweight

import "fmt"

const (
	TerroristDressType        = "tDress"
	CounterTerroristDressType = "ctDress"
)

var (
	dressFactorySingleton = &DressFactory{
		dressMap: make(map[string]Dress),
	}
)

type DressFactory struct {
	dressMap map[string]Dress
}

func GetDressFactorySingleton() *DressFactory {
	return dressFactorySingleton
}

func (d *DressFactory) getDressByType(dressType string) (Dress, error) {
	if d.dressMap != nil {
		return d.dressMap[dressType], nil
	}

	if dressType == TerroristDressType {
		return NewTerroristDress(), nil
	}

	if dressType == CounterTerroristDressType {
		return NewCounterTerroristDress(), nil
	}

	return nil, fmt.Errorf("Unknown DressType %v", dressType)
}
