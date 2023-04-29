package visitor

type Shape interface {
	getType() string
	accept(Visitor)
}

type Square struct {
	side int
}

func (s *Square) getType() string {
	return "Squre"
}

func (s *Square) accept(v Visitor) {
	v.visitSquare(s)
}

type Circle struct {
	radius int
}

func (c *Circle) getType() string {
	return "Circle"
}

func (c *Circle) accept(v Visitor) {
	v.visitCircle(c)
}

type Rectangle struct {
	length  int
	breadth int
}

func (t *Rectangle) getType() string {
	return "rectangle"
}

func (t *Rectangle) accept(v Visitor) {
	v.visitRectangle(t)
}
