package visitor

import (
	"fmt"
	"testing"
)

func TestVisitor(t *testing.T) {
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{length: 2, breadth: 3}

	areaCalculator := &AreaCalculator{}
	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)

	fmt.Println()
	
	middleCoordinates := &MiddleCoordinates{}
	square.accept(middleCoordinates)
	circle.accept(middleCoordinates)
	rectangle.accept(middleCoordinates)
}
