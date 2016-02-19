//Package Point is used for tracking the coordinates of everything in the game
package Point

type Point struct{
	X, Y int
	Parent *Point
}


//Opposite return the opposite point of a point given a parent point.
//*****
//*pco*
//*****
//Knowing the coordinates of p(parent point) and c(current point) we can calculate the opposite point's coordinates.
func (point Point) Opposite() Point {
	return Point{2 * point.X - point.Parent.X, 2 * point.Y - point.Parent.Y, &point}
}

//LineEquationRegardsToX is used to calculate the Y of a line equation.
//Given we have the coordinates of 2 points we can calculate the Y given X for any point belonging to the line described by the equation
func LineEquationRegardsToX(x int, x0 int, y0 int, x1 int, y1 int) int{
	 y := float32((x - x0)*(y1 - y0)/(x1 - x0) + y0)
	 return int(y + 0.5)
}

//LineEquationRegardsToY is used to calculate the X of a line equation.
//Given we have the coordinates of 2 points we can calculate the X given Y for any point belonging to the line described by the equation
func LineEquationRegardsToY(y int, x0 int, y0 int, x1 int, y1 int) int{
	 x := float32((y - y0)*(x1 - x0)/(y1 - y0) + x0)
	 return int(x + 0.5)
}
