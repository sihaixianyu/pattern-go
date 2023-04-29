package visitor

import "fmt"

const Pi float32 = 3.14

type Visitor interface {
	visitSquare(*Square)
	visitCircle(*Circle)
	visitRectangle(*Rectangle)
}

type AreaCalculator struct {
	area float32
}

func (a *AreaCalculator) visitSquare(s *Square) {
	a.area = float32(s.side * s.side)
	fmt.Println("Area of square: ", a.area)
}

func (a *AreaCalculator) visitCircle(c *Circle) {
	a.area = float32(c.radius) * Pi
	fmt.Println("Area of circle: ", a.area)
}

func (a *AreaCalculator) visitRectangle(r *Rectangle) {
	a.area = float32(r.breadth * r.length)
	fmt.Println("Area of rectangle: ", a.area)
}

type MiddleCoordinates struct{}

func (a *MiddleCoordinates) visitSquare(s *Square) {
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *MiddleCoordinates) visitCircle(c *Circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}
func (a *MiddleCoordinates) visitRectangle(t *Rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}
