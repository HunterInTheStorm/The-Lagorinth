/*
	Here we create a randomly generated maze/labyrinth using Prim's algorithm.
	Description of said algorithm can be found on https://en.wikipedia.org/wiki/Prim%27s_algorithm
	We also define structure Point which is used in the maze generation as well tracking the position
	of various objects/structures(npc, traps, the playeble character ect)
	Not only do we generate the labyrith but we also run chechs,
	for where to place truesures(in dead-ends of the maze) and traps(on crossroads) 
*/

package Labyrinth


import "math/rand"

var wall string = "0"
var pass string = " "
var treasure string = "$"
var trap string = "*"
var pointMap = map[Point]bool{}

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
	return Point{2 * point.x - point.parent.x, 2 * point.y - point.parent.y, &point}
}


//The stucture has two fields:
//integers width and height of the labyrinth/2D array
//2D array of strings(of characters) which will represent the generated maze
type Labyrinth struct{
	width, height int
	//rng int
	labyrinth [40][40]string
}


//The main algorithm used to generate the maze
//"0" for wall cells
//" " for empty cells(a.k.a path)
func (lab *Labyrinth) Prim(seed int64) {
	frontier := make([]Point, 0, 40)
	rand.Seed(seed)
	start := Point{rand.Intn(lab.width - 1) + 1, rand.Intn(lab.width - 1) + 1, nil}
	lab.labyrinth[start.x][start.y] = trap
	lab.Neighbours(&start, &frontier)
	for {
		randomPoint := rand.Intn(len(frontier))
		current := frontier[randomPoint]
		frontier = append(frontier[:randomPoint], frontier[randomPoint +1:]...)

		opposite := current.Opposite()
		last := opposite
		
		if lab.labyrinth[opposite.x][opposite.y] == wall {
			lab.labyrinth[current.x][current.y] = pass

			if opposite.x != 0 && opposite.x != lab.width - 1 && opposite.y != 0 && opposite.y != lab.height - 1 {
				lab.labyrinth[opposite.x][opposite.y] = pass
			}

			lab.Neighbours(&opposite, &frontier)
		}
		if len(frontier) == 0 {
			lab.labyrinth[last.x][last.y] = "E"
			break
		}
	}
}

//all neighbours(left, right, top, bottom) of a given Point will be passed to AddNeighbour
func(lab *Labyrinth) Neighbours(point *Point, frontier *[]Point) {
	lab.AddNeighbour(point.x + 1, point.y  	 , point, frontier)
	lab.AddNeighbour(point.x - 1, point.y 	 , point, frontier)
	lab.AddNeighbour(point.x    , point.y + 1,point, frontier)
	lab.AddNeighbour(point.x    , point.y - 1,point, frontier)
}


//adds the Neighbours af a give point to the frontier list which is used in the maze generation algorithm
func(lab *Labyrinth) AddNeighbour(x int, y int, parent *Point, frontier *[]Point) {
	if !pointMap[Point{x, y, parent}] {
		if (x > 0 && x < lab.width - 1) && (y > 0 && y < lab.height - 1) {
			pointToBeAdd := Point{x, y, parent}
			*frontier = append(*frontier, pointToBeAdd)
			pointMap[pointToBeAdd] = true
		}
	}
}

//This function will make all outher cells of the maze "wall" cells
//unless the algorithm is improved enough so that this step becomes obsolete
func(lab *Labyrinth) AddBorder() {

}

//Determines if a given point has 3 neighnours that are "wall" cells
func(lab *Labyrinth) IsDeadEnd(situation [4]string) bool {
	return true
}

// IsTreasure 25% to place a treasure at a dead-end in the maze
func(lab *Labyrinth) IsTreasure(x int, y int) bool {
	return true
}

//places a "T" for treasure in the 2d array at x and y coordinates
func(lab *Labyrinth) CreateTreasure(x int, y int) {

}

//Determines if a given point is a crossroad, a point that has 1 or 0 neighbours that are "wall" cells
func(lab *Labyrinth) IsCrossRoad(situation [4]string) bool {
	return true
}

//at a given crossroad randoms whethere the tile will be a trap
func(lab *Labyrinth) IsTrap(x int, y int,) bool {
	return true
}

//creates a trap at give coordinates
func(lab *Labyrinth) CreateTrap(x int, y int) {

}