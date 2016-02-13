package Game


import "github.com/golang/The-Lagorinth/Labyrinth"
import "github.com/golang/The-Lagorinth/Items"
import "github.com/golang/The-Lagorinth/Spells"
import "github.com/golang/The-Lagorinth/Characters"
import "github.com/golang/The-Lagorinth/Point"
import "time"
import "fmt"
import "math/rand"
import "os"
import "bufio"
import "strings"

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
	player *Character.Hero
}


func(game *Game) chooseName() string {
	fmt.Println("What is your name mighty adventurer?")
 	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	return strings.Trim(name,"\r\n")
}

func(game *Game) chooseClass() string {
	fmt.Println("What is your profession traveller?")
	fmt.Println("(Paladin/Mage/Rouge)")
	reader := bufio.NewReader(os.Stdin)
	class, _ := reader.ReadString('\n')
	return strings.Trim(class,"\r\n")
}

func(game *Game) chooseBackground() string {
	fmt.Println("And what are you renowned for?")
	fmt.Println("Giant-like strenght(giant)")
	fmt.Println("Aagile toreador(toreador)")
	fmt.Println("World Cartographer(cartographer)")
	fmt.Println("Wise Librarian(librarian)")
	reader := bufio.NewReader(os.Stdin)
	backGround, _ := reader.ReadString('\n')
	return strings.Trim(backGround,"\r\n")
}

// //this function will create one of the 3 classes for the player
func (game *Game) createPaladin(charName string, charBackGround string) {
	weapon := game.createWeapon()
	armor := game.createArmor()
	
	base := Character.NPC{&Point.Point{game.start.X, game.start.Y, nil}, Labyrinth.CharSymbol, charName,
		&Point.Point{1, 0, nil}, &weapon, &armor, Character.PaladinDmgMultuplier, Character.PaladinDefence, 
		Character.PaladinEvasion,Character.PaladinCritChance, Character.PaladinMaxHealth, 
		Character.PaladinMaxHealth, Character.PaladinHealthRegen, Character.PaladinMaxMana,
		Character.PaladinMaxMana, Character.PaladinManaRegen, Character.PaladinVisionRadious, 
		false, make(map[int]*Spells.Buff)}

	game.player = &Character.Hero{&base, "Paladin", charBackGround, make([]*Spells.Spell, 0, 3), 
		make(map[Point.Point]int), Character.PaladinMemoryDuration}
}

// //this function will create one of the 3 classes for the player
func (game *Game) createMage(charName string, charBackGround string) {
	weapon := game.createWeapon()
	armor := game.createArmor()
	base := Character.NPC{&Point.Point{game.start.X, game.start.Y, nil}, Labyrinth.CharSymbol, charName,
		&Point.Point{1, 0, nil}, &weapon, &armor, Character.MageDmgMultuplier, Character.MageDefence, 
		Character.MageEvasion, Character.MageCritChance, Character.MageMaxHealth, Character.MageMaxHealth,
		Character.MageHealthRegen, Character.MageMaxMana, Character.MageMaxMana, Character.MageManaRegen, 
		Character.MageVisionRadious, false, make(map[int]*Spells.Buff)}

	game.player = &Character.Hero{&base, "Mage", charBackGround, make([]*Spells.Spell, 0, 3), 
		make(map[Point.Point]int), Character.MageMemoryDuration}
}

// //Add some function for chace go hit/evade

// //this function will create one of the 3 classes for the player
func (game *Game) createRouge(charName string, charBackGround string) {
	weapon := game.createWeapon()
	armor := game.createArmor()
	base := Character.NPC{&Point.Point{game.start.X, game.start.Y, nil}, Labyrinth.CharSymbol, charName,
		&Point.Point{1, 0, nil}, &weapon, &armor, Character.RougeDmgMultuplier, Character.RougeDefence, 
		Character.RougeEvasion, Character.RougeCritChance, Character.RougeMaxHealth, Character.RougeMaxHealth, 
		Character.RougeHealthRegen, Character.RougeMaxMana, Character.RougeMaxMana, Character.RougeManaRegen, 
		Character.RougeVisionRadious, false, make(map[int]*Spells.Buff)}

	game.player = &Character.Hero{&base, "Rouge", charBackGround, make([]*Spells.Spell, 0, 3), 
		make(map[Point.Point]int), Character.RougeMemoryDuration}
}

// //this function will handle user input adn desired character creation
func (game *Game) createHero() {
	charName := game.chooseName()
	charClass := string(game.chooseClass())
	charBackGround := game.chooseBackground()
	switch charClass {
	case "Paladin":
		fmt.Print(charClass)
		game.createPaladin(charName, charBackGround)
	case "Rouge":
		fmt.Print(charClass)
		game.createRouge(charName, charBackGround)
	case "Mage":
		fmt.Print(charClass)
		game.createMage(charName, charBackGround)
	}
}

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

func (game *Game) characterMoveTo(char *Character.NPC, x int, y int) {
	char.Location.X = x
	char.Location.Y = y
}

// //given a user input the function handles the desired action the player wants to performe
func (game *Game) plyerActionEvent(x int, y int) {
	switch game.labyrinth.Labyrinth[x][y] {
	case Labyrinth.Pass:
		game.characterMoveTo(game.player.Base, x, y)
	case Labyrinth.StartPosition:
		game.characterMoveTo(game.player.Base, x, y)
	case Labyrinth.Monster:
		//fight the beast
	case Labyrinth.Trap:
		//disarm the trap or lose an arm
	case Labyrinth.Treasure:
		//get the loot
	case Labyrinth.ExitPosition:
		//exit
	}
	// return ""
}

//function takes user input and send it to PlayerActionEvent
func (game *Game) playerAction(key string) {
	//restore maze here??
	fmt.Println(key)
	switch key {
	case "w":
		game.plyerActionEvent(game.player.Base.Location.X + 1, game.player.Base.Location.Y)
	case "a":
		game.plyerActionEvent(game.player.Base.Location.X, game.player.Base.Location.Y - 1)
	case "s":
		game.plyerActionEvent(game.player.Base.Location.X, game.player.Base.Location.Y + 1)
	case "d":
		game.plyerActionEvent(game.player.Base.Location.X - 1, game.player.Base.Location.Y)
	case "1":
		fmt.Println("SO FAR SO GOOD, SPELL CAST 1")
	case "2":
		fmt.Println("SO FAR SO GOOD, SPELL CAST 2")
	case "3":
		fmt.Println("SO FAR SO GOOD, SCELL CAST 3")
	case "4":
		fmt.Println("SO FAR SO GOOD, CAMERA MOVEMENT LEFT")
	case "5":
		fmt.Println("SO FAR SO GOOD, CAMERA MOVEMENT DOWN")
	case "6":
		fmt.Println("SO FAR SO GOOD, CAMERA MOVEMETN RIGHT")
	case "8":
		fmt.Println("SO FAR SO GOOD,CAMERA MOVEMENT UP")
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

func(game *Game) createArmor() Items.Armor {
	armor := Items.Armor{}
	armor.RandomizeArmor()
	return armor
}

func(game *Game) createWeapon() Items.Weapon {
	weapon := Items.Weapon{}
	weapon.RandomizeWeapon()
	return weapon
}

// the function will create the monsters in the maze
func (game *Game) createMonster(x int, y int) Character.NPC{
	weapon := game.createWeapon()
	armor := game.createArmor()
	//TRANSFER VALUES TO SEPARATE FILE
	return Character.NPC{&Point.Point{x, y, nil}, Labyrinth.Monster, "Skeleton", &Point.Point{-1, 0, nil},
	&weapon, &armor, 1,  10, 3, 5, 120.0, 120.0, 1.5, 30, 30, 0.2, 2, false, make(map[int]*Spells.Buff)}
}

// the function will create the traps in the maze
func (game *Game) createTrap(x int, y int) Character.Trap{
	trapType := Character.TrapTypes[rand.Intn(len(Character.TrapTypes))]
	detectDiff := rand.Intn(10) + 1
	disarmDiff := rand.Intn(5) + 1
	minDmg := rand.Intn(6)
	maxDmg := rand.Intn(6) + minDmg
	return Character.Trap{&Point.Point{x, y, nil},trapType,  detectDiff, disarmDiff, false, false, minDmg, maxDmg}
}

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
	//TRY TO DO SOME MAGIC WITH CHANNELS BY SENDING POINTS WITH LOCATION OF START,END,NPCs,TRAPs
	game.createHero()
}


// func (game *Game) printdata() {
// 	fmt.Println("playerDefeted")
// 	fmt.Println(game.playerDefeted)
// 	fmt.Println("gameCompleted")
// 	fmt.Println(game.gameCompleted)
// 	fmt.Println("score")
// 	fmt.Println(game.score)
// 	fmt.Println("turns")
// 	fmt.Println(game.turns)
// 	fmt.Println("monsterSlain")
// 	fmt.Println(game.monsterSlain)
// 	fmt.Println("chestsLooted")
// 	fmt.Println(game.chestsLooted)
// 	fmt.Println("trapsDisarmed")
// 	fmt.Println(game.trapsDisarmed)
// 	fmt.Println("start")
// 	fmt.Println(game.start)
// 	fmt.Println("end")
// 	fmt.Println(game.end)
// 	fmt.Println("monsterList")
// 	fmt.Println(game.monsterList)
// 	fmt.Println("trapList")
// 	fmt.Println(game.trapList)
// 	fmt.Println("player")
// 	fmt.Println(game.player)
//}

// //function replaces an element fro the 2d array for the maze with the character symbol
func (game *Game) drawCharacters() {
	game.labyrinth.Labyrinth[game.player.Base.Location.X][game.start.Y] = Labyrinth.CharSymbol
	for trapPoint, _ := range game.trapList {
		game.labyrinth.Labyrinth[trapPoint.X][trapPoint.Y] = Labyrinth.Trap
	}
	//add projectile here
	for _, mon := range game.monsterList {
		game.labyrinth.Labyrinth[mon.Location.X][mon.Location.Y] = Labyrinth.Monster
	}

}

//function to draw the labyrinth
func (game *Game) drawLabyrinth() {
	for i := 0; i < game.labyrinth.Width; i++ {
		for j := 0; j < game.labyrinth.Height; j++ {
			fmt.Print(game.labyrinth.Labyrinth[i][j])
		}
		fmt.Println()
	}
}

func (game *Game) detectKeyPress() string{
	reader := bufio.NewReader(os.Stdin)
	key, _ := reader.ReadString('\n')
	return strings.Trim(key,"\r\n")
}

//main loop cycle for the game
func (game *Game) Run() {
	game.initialize()

	for  {
		game.drawCharacters()
		game.drawLabyrinth()
		key := game.detectKeyPress()
		game.playerAction(key)
		if game.playerDefeted || game.gameCompleted {
			break
		}
	}

	//game.Evaluation()
}
