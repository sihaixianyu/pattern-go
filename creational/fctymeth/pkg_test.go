package fctymeth

import (
	"fmt"
	"testing"
)

func TestGunFactory(t *testing.T) {
	ak47, _ := GetGun("ak47")
	musket, _ := GetGun("musket")

	fmt.Println(ak47)
	fmt.Println(musket)
}
