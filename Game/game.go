package Game

import "github.com/golang/The-Lagorinth/Labyrinth"
import "github.com/golang/The-Lagorinth/Spells"
import "github.com/golang/The-Lagorinth/Characters"
import "github.com/golang/The-Lagorinth/Point"

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
	player *Character.Hero
	camera *Point.Point
	cameraRadius int
	projectileList []*Spell.Projectile
}

// //this function will save a high score to a file
// func (game *Game) SaveHighScore() {

// }

// //function will create a new monster at random intervals of time
// func (game *Game) SpawnMonster() {

// }

// //function will set the win condition when the end of the maze has been found
// func (game *Game) Evaluation() {
// 	return true
// }


//main loop cycle for the game
func (game *Game) Run() {
	game.initialize()

	for  {
		game.manageSpells()
		game.player.UpdateMemory()
		game.draw()
		game.playerAction()
		game.npcsTurn()
		game.applyRegenToAll()
		game.turns++
		if game.playerDefeted || game.gameCompleted {
			break
		}
	}

	//game.Evaluation()
}
