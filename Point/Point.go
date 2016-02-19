package Point

type Point struct{
	X, Y int
	Parent *Point
}

//This function returns the opposite of a given Point given its Parent Point
//The returned Point is used for our maze generation sequence
//refer to Prim's algorithm for more details
func (point Point) Opposite() Point {
	return Point{2 * point.X - point.Parent.X, 2 * point.Y - point.Parent.Y, &point}
}

func LineEquationRegardsToX(x int, x0 int, y0 int, x1 int, y1 int) int{
	 y := float32((x - x0)*(y1 - y0)/(x1 - x0) + y0)
	 return int(y + 0.5)
}

func LineEquationRegardsToY(y int, x0 int, y0 int, x1 int, y1 int) int{
	 x := float32((y - y0)*(x1 - x0)/(y1 - y0) + x0)
	 return int(x + 0.5)
}
