package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilder(t *testing.T) {
	normal, err := GetBuilder("normal")
	if err != nil {
		panic(err.Error())
	}
	igloo, err := GetBuilder("igloo")
	if err != nil {
		panic(err.Error())
	}

	director := NewDirector(normal)
	house := director.BuildHouse()

	assert.Equal(t, "wooden window", house.window)
	assert.Equal(t, "wooden door", house.door)

	director.SetBuilder(igloo)
	house = director.BuildHouse()

	assert.Equal(t, "snow window", house.window)
	assert.Equal(t, "snow door", house.door)
}
