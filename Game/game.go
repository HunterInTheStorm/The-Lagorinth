package Game

import "github.com/golang/The-Lagorinth/Labyrinth"
import "github.com/golang/The-Lagorinth/Items"
import "github.com/golang/The-Lagorinth/Spells"
import "github.com/golang/The-Lagorinth/Characters"
import "github.com/golang/The-Lagorinth/Point"
import "time"
import "fmt"

//This structure discribes the rules of the game concerning character creation,
//player and npc movement, contain various array to trac monsters, traps and mimics,
//win/lose  conditions adn statistics effecting the final score
type Game struct {
	playerDefeted bool
	gameCompleted bool
	score int
	turns int
	monsterSlain int
	chestsLooted int
	trapsDisarmed int
	start *Point.Point
	end *Point.Point
	monsterList []*Character.NPC
	trapList map[Point.Point]*Character.Trap
	labyrinth *Labyrinth.Labyrinth
	player *Character.Character
}


// //this function will handle user input adn desired character creation
// func (game *Game) CreateCharacter() {

// }

// //this function will create one of the 3 classes for the player
// func (game *Game) CreatePaladin() {

// }

// //this function will create one of the 3 classes for the player
// func (game *Game) CreateMage() {

// }

// //Add some function for chace go hit/evade

// //this function will create one of the 3 classes for the player
// func (game *Game) CreateRouge() {

// }

// //function will handle choice of backgroung for the character and applying the coresponding bonuses
// func (game *Game) ChooseBackground() {

// }

// the function will create the monsters in the maze
func (game *Game) createMonster(x int, y int) Character.NPC{
	return Character.NPC{&Point.Point{x, y, nil}, Labyrinth.Monster, "Skeletron", &Point.Point{-1, 0, nil},
	&Items.Weapon{3,4,5,6}, &Items.Armor{2,3,3.1,3.2,3.2,3.4}, 3.2,  1, 1, 1, 31.1, 31.1, 31.1, 31.1, 31.1, 31.1, 4, false, false, make(map[int]*Spells.Buff)}
}

// the function will create the traps in the maze
func (game *Game) createTrap(x int, y int) Character.Trap{
	return Character.Trap{&Point.Point{x, y, nil}, 0, false, false, 0 , 0}
}

// // the function will create the maze in the game
// func (game *Game) CreateLabyrinth() {

// }

// // the function will calculate the final score the player has acheived
// //depending on turs passed, mosters killed, cheasts looted
// func (game *Game) CalculateFinalScore() {

// }


// //this function will save a highscore to a file
// func (game *Game) SaveHighScore() {

// }

// //function will determine if a character can move to a certain possition
// func (game *Game) CanMoveTo() bool {
// 	return true
// }

// //function will remove a monster from the monster list
// func (game *Game) MonsterDefeted() {

// }

// //function will remove a trap from the trap list
// func (game *Game) TrapDisarmed() {

// }


// //this function will handle the fight event
// func (game *Game) Fight(npc NPC) {

// }

// //this function will find the moster in the list with which the player will engage in combat
// func (game *Game) FindMonster() NPC{
// 	return NPC{}
// }

// //function takes user input and send it to PlayerActionEvent
// func (game *Game) PlayerAction() {

// }

// //given a user input the function handles the desired action the player wants to performe
// func (game *Game) PlyerActionEvent(action string) {

// }

// //function replaces an element fro the 2d array for the maze with the character symbol
// func (game *Game) DrawCharacters() {

// }

//function to draw the labyrinth
func (game *Game) DrawLabyrinth() {
	for i := 0; i < game.labyrinth.Width; i++ {
		for j := 0; j < game.labyrinth.Height; j++ {
			fmt.Print(game.labyrinth.Labyrinth[i][j])
		}
		fmt.Println()
	}
}

// //function to restore the original state of the 2d maze array
// func (game *Game) RestoreLabyrinth(x, y int) {

// }

// //function will create a new monster at random intervals of time
// func (game *Game) SpawnMonster() {

// }

// //function will set the win condotion when the end of the maze has benn found
// func (game *Game) EndReached() bool {
// 	return true
// }

// //function will set the win condotion when the end of the maze has benn found
// func (game *Game) Evaluation() {
// 	return true
// }

func (game *Game) setGameFieldValues() {
	game.playerDefeted = false
	game.gameCompleted = false
	game.score = 0
	game.turns = 0
	game.monsterSlain = 0
	game.chestsLooted = 0
	game.trapsDisarmed = 0
}

func (game *Game) createLabyrinth() {
	var seed int64 = time.Now().UTC().UnixNano()
	var maze [40][40]string
	for i := 0; i < 40; i++ {
		for j := 0; j < 40; j++ {
			maze[i][j] = Labyrinth.Wall
		}
	}
	game.labyrinth = &Labyrinth.Labyrinth{40,40, maze}
	game.labyrinth.Prim(seed)
}

func (game *Game) createMonsterAndTrapsLists() {
	game.monsterList = make([]*Character.NPC,0,4)
	game.trapList = make(map[Point.Point]*Character.Trap)
	for i := 0; i < 40; i++ {
		for j := 0; j < 40; j++ {
			tile := game.labyrinth.Labyrinth[i][j]
			if tile != Labyrinth.Wall {
				switch tile {
				case Labyrinth.StartPosition:
					game.start = &Point.Point{i, j, nil}
				case Labyrinth.ExitPosition:
					game.end = &Point.Point{i, j, nil}
				case Labyrinth.Trap:
				 	newTrap := game.createTrap(i, j)
				 	newTrap.IsDisarmed = true
					game.trapList[*newTrap.Location] = &newTrap
				case Labyrinth.Monster:
					newMonster := game.createMonster(i, j)
					game.monsterList = append(game.monsterList, &newMonster)
				}
			}
		}
	}
}

func (game *Game) initialize() {
	game.setGameFieldValues()
	game.createLabyrinth()
	game.createMonsterAndTrapsLists()
	//TRY TO DO SOME MAGIC WITH CHANNELS BY SENDING POINTS WITH LOCATION OF START,END,NPCs,TRAPS	
}

//main loop cycle for the game
func (game *Game) Run() {
	game.initialize()
	//CREATE PLAER CHARACTER
	//MAIN LOOP CYCLE
	for  {
		game.DrawLabyrinth()
		break
		if game.playerDefeted || game.gameCompleted {
			break
		}
	}

	//game.Evaluation()
}
