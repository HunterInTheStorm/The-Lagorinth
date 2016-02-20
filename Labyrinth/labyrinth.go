//Package Labyrinth handles the creation of a creation of a labyrinth.
//Prim's algorithm is used for the creation pf said labyrinth.
//More detail can be found here: https://en.wikipedia.org/wiki/Prim%27s_algorithm
//The algorithm ensures a randomly generated labyrinth.
//In addition, the package expands on the idea providing algorithm for chest and enemy placement in the labyrinth(random).
package Labyrinth

import "github.com/golang/The-Lagorinth/Point"
import "math/rand"

//Global variables used in the project.
//This way it is easy to modify objects in the fly.
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

type Labyrinth struct {
	Width, Height int
	Labyrinth     [40][40]string
}

//CreateLabyrinth call Prim function which handles the creation pf the labyrinth.
//It sets the rand seed.
//It also expands the labyrinth by calling createTreasuresAndTraps function.
//That function determines the location of chests, traps and enemies.
func (lab *Labyrinth) CreateLabyrinth(seed int64) {
	rand.Seed(seed)
	lab.Prim()
	lab.createTreasuresAndTraps()
}

//Prim is the algorithm used to generate the maze.
//refer to package description for more details on Prim's algorithm.
func (lab *Labyrinth) Prim() {
	frontier := make([]Point.Point, 0, 40)
	var start Point.Point = Point.Point{rand.Intn(lab.Width-1) + 1, rand.Intn(lab.Width-1) + 1, nil}
	lab.Labyrinth[start.X][start.Y] = StartPosition
	lab.neighbours(&start, &frontier)
	for {
		randomPoint := rand.Intn(len(frontier))
		current := frontier[randomPoint]
		frontier = append(frontier[:randomPoint], frontier[randomPoint+1:]...)

		opposite := current.Opposite()
		last := opposite

		if lab.Labyrinth[opposite.X][opposite.Y] == Wall {
			lab.Labyrinth[current.X][current.Y] = Pass

			if opposite.X != 0 && opposite.X != lab.Width-1 && opposite.Y != 0 && opposite.Y != lab.Height-1 {
				lab.Labyrinth[opposite.X][opposite.Y] = Pass
			}

			lab.neighbours(&opposite, &frontier)
		}
		if len(frontier) == 0 {
			lab.Labyrinth[last.X][last.Y] = "E"
			break
		}
	}
}

//neighbours takes 2 arguments.
//A slice which is passed on to addNeighbour.
//And a point. Function determines all of the point's neighbours and passes their coordinates to addNeighbour.
//Point is passed on to addNeighbour as well to be used as a parent point for its neighbours.
func (lab *Labyrinth) neighbours(point *Point.Point, frontier *[]Point.Point) {
	lab.addNeighbour(point.X+1, point.Y, point, frontier)
	lab.addNeighbour(point.X-1, point.Y, point, frontier)
	lab.addNeighbour(point.X, point.Y+1, point, frontier)
	lab.addNeighbour(point.X, point.Y-1, point, frontier)
}

//addNeighbour add a point to frontier slice if the point is not has not been added already.
//Function takes 4 arguments.
//2 coordinates for the point to be added to frontier slice.
//A parent point.
//And the slice itself.
func (lab *Labyrinth) addNeighbour(x int, y int, parent *Point.Point, frontier *[]Point.Point) {
	if !pointMap[Point.Point{x, y, parent}] {
		if (x > 0 && x < lab.Width-1) && (y > 0 && y < lab.Height-1) {
			pointToBeAdd := Point.Point{x, y, parent}
			*frontier = append(*frontier, pointToBeAdd)
			pointMap[pointToBeAdd] = true
		}
	}
}

//countWalls counts how many of the neighbours of a point, in the labyrinth, are walls and returns that count.
//Function takes 2 arguments, the coordinates of the point.
func (lab *Labyrinth) countWalls(x int, y int) int {
	var wallCount int = 0
	if lab.Labyrinth[x+1][y] == Wall {
		wallCount++
	}
	if lab.Labyrinth[x-1][y] == Wall {
		wallCount++
	}
	if lab.Labyrinth[x][y+1] == Wall {
		wallCount++
	}
	if lab.Labyrinth[x][y-1] == Wall {
		wallCount++
	}
	return wallCount
}

//isDeadEnd returns true if a point in the labyrinth has 3 neighbours that are walls
//Function takes 2 arguments, the coordinates of the point.
func (lab *Labyrinth) isDeadEnd(x int, y int) bool {
	return lab.countWalls(x, y) == 3
}

//isTreasure return true if a random value is lower than chanceToBeTreasure global variable.
//Function is used for the placement of treasures in the labyrinth.
func (lab *Labyrinth) isTreasure() bool {
	return rand.Intn(100) < chanceToBeTreasure
}

//placeNpc places a monster symbol in the labyrinth if a random value is lower than chanceToBeNpc global variable.
//It takes 2 arguments, the coordinates of the point where the symbol will be places.
func (lab *Labyrinth) placeNpc(x int, y int) {
	if rand.Intn(100) < chanceToBeNpc {
		lab.Labyrinth[x][y] = Monster
	}
}

//createTreasuresAndTraps handles the placement of treasures, traps and monsters in the labyrinth.
func (lab *Labyrinth) createTreasuresAndTraps() {
	for i := 1; i < lab.Width-1; i++ {
		for j := 1; j < lab.Height-1; j++ {
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

//IsInBondaries determines of the arguments are within the labyrinth.
//Will return false if the arguments are negative or greater than the dimensions of the labyrinth.
func (lab *Labyrinth) IsInBondaries(x int, y int) bool {
	return x > -1 && x < lab.Width && y > -1 && y < lab.Height
}

//isCrossRoad returns true if 1 or less of the neighbours of point are walls.
func (lab *Labyrinth) isCrossRoad(x int, y int) bool {
	return lab.countWalls(x, y) < 2
}

//isTrap returns true if a random value is lower than the chanceToBeTrap global variable.
func (lab *Labyrinth) isTrap() bool {
	return rand.Intn(100) < chanceToBeTrap
}
