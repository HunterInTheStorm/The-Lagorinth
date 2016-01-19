/*
	Here we create a randomly generated maze/labyrinth using Prim's algorithm.
	Description of said algorithm can be found on https://en.wikipedia.org/wiki/Prim%27s_algorithm
	We also define structure Point which is used in the maze generation as well tracking the position
	of various objects/structures(npc, traps, the playeble character ect)
	Not only do we generate the labyrith but we also run chechs,
	for where to place truesures(in dead-ends of the maze) and traps(on crossroads) 
*/

package main


//The structure has 3 fields:
//	integers for X and Y coordinates
//	field for parent Point
type Point struct{
	x,y int
	parent *Point
}

//This function returns the opposite of a given Point given its Parent Point
//The returned Point is used for our maze generation sequence
//refer to Prim's algorithm for more details
func (point Point) Opposite() Point {
	return Point
}


//The stucture has two fields:
//integers width and heigth of the labyrinth/2D array
//2D array of strings(of characters) which will represent the generated maze
type Labyrinth struct{
	width, heigt int
	//rng int
	labyrinth [][]string
}


//The main algorithm used to generate the maze
//"0" for wall cells
//" " for empty cells(a.k.a path)
func (lab Labyrinth) Prim() {

}


//This function will make all outher cells of the maze "wall" cells
//unless the algorithm is improved enough so that this step becomes obsolete
func(lab Labyrinth) AddBorder() {

}


//all neighbours(left, right, top, bottom) of a given Point will be passed to AddNeighbour
func(lab Labyrinth) Neighbours(point Point, frontier []Point) {

}


//adds the Neighbours af a give point to the frontier list which is used in the maze generation algorithm
func(lab Labyrinth) AddNeighbours(x int, y int, parent Point, frontier) {

}


//25% to place a treasure at a dead-end in the maze
func(lab Labyrinth) IsTreasure(x int, y int) bool {
	return true
}

//Determines if a given point has 3 neighnours that are "wall" cells
func(lab Labyrinth) IsDeadEnd(situation [4]string) bool {
	return true
}

//places a "T" for treasure in the 2d array at x and y coordinates
func(lab Labyrinth) CreateTreasure(x int, y int) {

}

//Determines if a given point is a crossroad, a point that has 1 or 0 neighbours that are "wall" cells
func(lab Labyrinth) IsCrossRoad(situation [4]string) bool {
	return true
}

//at a given crossroad randoms whethere the tile will be a trap
func(lab Labyrinth) IsTrap(x int, y int,) bool {
	return true
}

//creates a trap at give coordinates
func(lab Labyrinth) CreateTrap(x int, y int) {

}