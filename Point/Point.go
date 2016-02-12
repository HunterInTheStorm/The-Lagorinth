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