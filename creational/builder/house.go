package builder

import "errors"

type House struct {
	window string
	door   string
}

type HouseBuilder interface {
	setWindow()
	setDoor()
	getHouse() House
}

func GetBuilder(builderType string) (HouseBuilder, error) {
	if builderType == "normal" {
		return newNormalBuilder(), nil
	}

	if builderType == "igloo" {
		return newIglooBuilder(), nil
	}

	return nil, errors.New("unsopported builder type")
}
