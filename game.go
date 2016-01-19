package main


//This structure discribes the rules of the game concerning character creation,
//player and npc movement, contain various array to trac monsters, traps and mimics,
//win/lose  conditions adn statistics effecting the final score
type Game struct {
	labyrinth Labyrinth
	player Character
	playerDefeted bool
	gameCompleted bool
	monsters []NPC
	traps []Traps
	mimics []Mimic
	start Point
	end Point
	score int
	turns int
	monsterSlain int
	chestsLooted int
	trapsDisarmed
}


//this function will handle user input adn desired character creation
func (game Game) CreateCharacter() {

}

//this function will create one of the 3 classes for the player
func (game Game) CreatePaladin() {

}

//this function will create one of the 3 classes for the player
func (game Game) CreateMage() {

}

//this function will create one of the 3 classes for the player
func (game Game) CreateRouge() {

}

//function will handle choice of backgroung for the character and applying the coresponding bonuses
func (game Game) ChooseBackground() {

}

// the function will create the monsters in the maze
func (game Game) CreateMonsters() {

}

// the function will create the mimics in the maze
func (game Game) CreateMimics() {
	
}

// the function will create the traps in the maze
func (game Game) CreateTraps() {
	
}

// the function will create the maze in the game
func (game Game) CreateLabyrinth() {

}

// the function will calculate the final score the player has acheived
//depending on turs passed, mosters killed, cheasts looted
func (game Game) CalculateFinalScore() {

}


//this function will save a highscore to a file
func (game Game) SaveHighScore() {

}

//function will determine if a character can move to a certain possition
func (game Game) CanMoveTo() bool {
	return true
}

//function will remove a monster from the monster list
func (game Game) MonsterDefeted() {

}

//function will remove a trap from the trap list
func (game Game) TrapDisarmed() {

}

//function will remove a mimic from the mimic list
func (game Game) MimicDefeted() {

}

//this function will handle the fight event
func (game Game) Fight(npc NPC) {

}

//this function will find the moster in the list with which the player will engage in combat
func (game Game) FindMonster() NPC{
	return NPC{}
}

//function takes user input and send it to PlayerActionEvent
func (game Game) PlayerAction() {

}

//given a user input the function handles the desired action the player wants to performe
func (game Game) PlyerActionEvent(action string) {

}

//function replaces an element fro the 2d array for the maze with the character symbol
func (game Game) DrawCharacters() {

}

//function to draw the labyrinth
func (game Game) DrawLabyrinth() {

}

//function to restore the original state of the 2d maze array
func (game Game) RestoreLabyrinth(x, y int) {

}

//function will create a new monster at random intervals of time
func (game Game) SpawnMonster() {

}

//function will set the win condotion when the end of the maze has benn found
func (game Game) EndReached() bool {
	return true
}

//main loop cycle for the game
func (game Game) Run() {

}