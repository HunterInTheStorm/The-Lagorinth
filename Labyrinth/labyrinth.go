/*
	Here we create a randomly generated maze/labyrinth using Prim's algorithm.
	Description of said algorithm can be found on https://en.wikipedia.org/wiki/Prim%27s_algorithm
	We also define structure Point which is used in the maze generation as well tracking the position
	of various objects/structures(npc, traps, the playeble character ect)
	Not only do we generate the labyrith but we also run chechs,
	for where to place truesures(in dead-ends of the maze) and traps(on crossroads) 
*/



package Labyrinth

import "github.com/golang/The-Lagorinth/Point"
import "math/rand"

var Wall string = "0"
var Pass string = " "
var Treasure string = "$"
var Trap string = "*"
var Monster string = "i"
var CharSymbol string = "c"
var ExitPosition string = "E"
var StartPosition string = "S"
var Projectile string = "o"
var pointMap = map[Point.Point]bool{}
var chanceToBeTreasure int = 25
var chanceToBeTrap int = 10
var chanceToBeNpc int = 15

//The structure has 3 fields:
//	integers for X and Y coordinates
//	field for parent Point


//The stucture has two fields:
//integers width and height of the labyrinth/2D array
//2D array of strings(of characters) which will represent the generated maze
type Labyrinth struct{
	Width, Height int
	//rng int
	Labyrinth [40][40]string
}

//The main algorithm used to generate the maze
//"0" for wall cells
//" " for empty cells(a.k.a path)
func (lab *Labyrinth) Prim(seed int64) {
	frontier := make([]Point.Point, 0, 40)
	rand.Seed(seed)
	var start Point.Point = Point.Point{rand.Intn(lab.Width - 1) + 1, rand.Intn(lab.Width - 1) + 1, nil}
	lab.Labyrinth[start.X][start.Y] = StartPosition
	lab.neighbours(&start, &frontier)
	for {
		randomPoint := rand.Intn(len(frontier))
		current := frontier[randomPoint]
		frontier = append(frontier[:randomPoint], frontier[randomPoint +1:]...)

		opposite := current.Opposite()
		last := opposite
		
		if lab.Labyrinth[opposite.X][opposite.Y] == Wall {
			lab.Labyrinth[current.X][current.Y] = Pass

			if opposite.X != 0 && opposite.X != lab.Width - 1 && opposite.Y != 0 && opposite.Y != lab.Height - 1 {
				lab.Labyrinth[opposite.X][opposite.Y] = Pass
			}

			lab.neighbours(&opposite, &frontier)
		}
		if len(frontier) == 0 {
			lab.Labyrinth[last.X][last.Y] = "E"
			break
		}
	}
	lab.createTreasuresAndTraps()
}

//all neighbours(left, right, top, bottom) of a given Point will be passed to AddNeighbour
func(lab *Labyrinth) neighbours(point *Point.Point, frontier *[]Point.Point) {
	lab.addNeighbour(point.X + 1, point.Y  	 , point, frontier)
	lab.addNeighbour(point.X - 1, point.Y 	 , point, frontier)
	lab.addNeighbour(point.X    , point.Y + 1,point, frontier)
	lab.addNeighbour(point.X    , point.Y - 1,point, frontier)
}


//adds the Neighbours af a give point to the frontier list which is used in the maze generation algorithm
func(lab *Labyrinth) addNeighbour(x int, y int, parent *Point.Point, frontier *[]Point.Point) {
	if !pointMap[Point.Point{x, y, parent}] {
		if (x > 0 && x < lab.Width - 1) && (y > 0 && y < lab.Height - 1) {
			pointToBeAdd := Point.Point{x, y, parent}
			*frontier = append(*frontier, pointToBeAdd)
			pointMap[pointToBeAdd] = true
		}
	}
}

func(lab *Labyrinth) countWalls(x int, y int) int {
	var wallCount int = 0
	if lab.Labyrinth[x + 1][y] == Wall {
		wallCount++
	}
	if lab.Labyrinth[x - 1][y] == Wall {
		wallCount++
	}
	if lab.Labyrinth[x][y + 1] == Wall {
		wallCount++
	}
	if lab.Labyrinth[x][y - 1] == Wall {
		wallCount++
	}
	return wallCount
}

//Determines if a given point has 3 neighnours that are "wall" cells
func(lab *Labyrinth) isDeadEnd(x int, y int) bool {
	return lab.countWalls(x, y) == 3
}

// IsTreasure 25% to place a treasure at a dead-end in the maze
func(lab *Labyrinth) isTreasure() bool {
	return rand.Intn(100) < chanceToBeTreasure 
}

func(lab *Labyrinth) placeNpc(x int, y int) {
	if rand.Intn(100) < chanceToBeNpc {
		lab.Labyrinth[x][y] = Monster
	}
}

//places a "T" for treasure in the 2d array at x and y coordinates
func(lab *Labyrinth) createTreasuresAndTraps() {
	for i := 1; i < lab.Width - 1; i++ {
		for j := 1; j < lab.Height - 1; j++ {
			if lab.Labyrinth[i][j] == Pass {
				if lab.isDeadEnd(i, j) {
					if lab.isTreasure() {
						lab.Labyrinth[i][j] = Treasure
					} else {
						lab.placeNpc(i, j)
					}
				}
				if lab.isCrossRoad(i, j) {
					if lab.isTrap() {
						lab.Labyrinth[i][j] = Trap
					} else {
						lab.placeNpc(i, j)
					}
				}
			}
		}
	}
}

func (lab *Labyrinth) IsInBondaries(x int, y int) bool {
	return x > -1 && x < lab.Width && y > -1 && y < lab.Height
}

//Determines if a given point is a crossroad, a point that has 1 or 0 neighbours that are "wall" cells
func(lab *Labyrinth) isCrossRoad(x int, y int) bool {
	return lab.countWalls(x, y) < 2
}

//at a given crossroad randoms whethere the tile will be a trap
func(lab *Labyrinth) isTrap() bool {
	return rand.Intn(100) < chanceToBeTrap
}
