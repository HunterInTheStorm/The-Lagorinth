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
import "os/exec"
import "bufio"
import "strings"

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
		false, make(map[int]*Spells.Buff), true, Character.PaladinTrapHandling}

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
		Character.MageVisionRadious, false, make(map[int]*Spells.Buff), true, Character.MageTrapHandling}

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
		Character.RougeVisionRadious, false, make(map[int]*Spells.Buff), true, Character.RougeTrapHandling}

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
		game.createPaladin(charName, charBackGround)
	case "Rouge":
		game.createRouge(charName, charBackGround)
	case "Mage":
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

func (game *Game) isTrapTriggered(character *Character.NPC) (*Character.Trap, bool) {
	if trap, ok := game.trapList[Point.Point{character.Location.X, character.Location.Y, nil}]; ok {
    	return trap, true
	}
	return &Character.Trap{}, false
}

func (game *Game) removeTrap(trap *Character.Trap) {
	game.restoreTile(trap.Location.X, trap.Location.Y)
	delete(game.trapList, *trap.Location)
}

func (game *Game) calculateOddsVsTraps(difficulty int, trapHandlingSkill int) int {
	return 100 - difficulty * 10 + trapHandlingSkill * 5
}

func (game *Game) attempDisarmTrap(trap *Character.Trap, character *Character.NPC) {
	chance := game.calculateOddsVsTraps(trap.DisarmDifficulty, character.TrapHandling)
	if rand.Intn(100) < chance {
		fmt.Println("TRAP DISARMED!. HELL YEAH!!!!!!!!!!")
		time.Sleep(2000 * time.Millisecond)
		trap.IsDisarmed = true
		trap.CanBeDisarmed = false
		game.trapsDisarmed ++
		game.removeTrap(trap)
	} else {
		fmt.Println("YOU ARE SUCH A DISAPPOINTMENT")
		time.Sleep(2000 * time.Millisecond)
		trap.CanBeDisarmed = false
	}
}

func (game *Game) attempDetectTrap(trap *Character.Trap, character *Character.NPC) {
	chance := game.calculateOddsVsTraps(trap.DetectDifficulty, character.TrapHandling)
	if rand.Intn(100) < chance {
		fmt.Println("TRAP DETECTED!!!!!!!")
		time.Sleep(2000 * time.Millisecond)
		trap.IsDetected = true
		trap.CanBeDetected = false
	} else {
		fmt.Println("TRAP NOT DETECTED :(")
		time.Sleep(2000 * time.Millisecond)
		trap.CanBeDetected = false
	}
}

func (game *Game) encounterTrap(character *Character.NPC, x int, y int) {
	trap := game.trapList[Point.Point{x, y, nil}]
	if trap.CanBeDetected && !trap.IsDetected {
		game.attempDetectTrap(trap, character)
		if !trap.IsDetected {
			game.characterMoveTo(character, x, y)
		} 
	} else if trap.IsDetected && trap.CanBeDisarmed {
		fmt.Println("Do you want to disarm the trap(y/n)")
		answer := game.detectKeyPress()
		if answer == "y"{
			game.attempDisarmTrap(trap, character)
		} else if answer == "n" {
			game.characterMoveTo(character, x, y)
		}
	} else if !trap.IsDetected || !trap.IsDisarmed{
		game.characterMoveTo(character, x, y)
	} else {
		game.characterMoveTo(character, x, y)
	}
}

func (game *Game) removeMonster(place int) {
	game.monsterList = append(game.monsterList[:place], game.monsterList[place +1:]...)
}

func (game *Game) playerDefetedMessage() {
	fmt.Println("Player defeted")
	time.Sleep(2000 * time.Millisecond)
}

// //function will remove a monster from the monster list
func (game *Game) monsterDefetedMessage(name string, playerName string) {
	fmt.Println("Player defets monster")
	time.Sleep(2000 * time.Millisecond)
}

func (game *Game) isCharacterDefeted(character *Character.NPC, place int) {
	if character.CurrentHealth < 0 {
		switch true {
		case place == -1:
			game.playerDefeted = true
			game.playerDefetedMessage()
		case place > -1:
			game.removeMonster(place)
			game.monsterSlain++
			game.restoreTile(character.Location.X, character.Location.Y)
			game.monsterDefetedMessage(character.Name, game.player.Base.Name)
		}
	}
}

func (game *Game) avoidAttackMessage(attackerName string, defenderName string) {
	fmt.Println("Defender dodges")
	time.Sleep(2000 * time.Millisecond)
}

func (game *Game) takeDamageFromTrapMessage(damage float32, attackerName string, defender *Character.NPC) {
	fmt.Printf("%s strikes %s for %f points of damage.\n", attackerName, defender.Name,
						 damage)
	time.Sleep(2000 * time.Millisecond)
}

func (game *Game) takeDamageMessage(damage float32, attacker *Character.NPC, defender *Character.NPC) {
	fmt.Printf("%s strikes %s for %f points of damage. %s has %f HP left\n", attacker.Name, defender.Name,
						 damage,defender.Name, defender.CurrentHealth)
	time.Sleep(2000 * time.Millisecond)
}

// //this function will find the moster in the list with which the player will engage in combat
//ADDDDDDDD ERRRRRRRORRRRRRRRRRRRRRRRRRRR
func (game *Game) findEnemy(requiredX int, requiredY int) (int, *Character.NPC) {
	if game.player.Base.Location.X == requiredX && game.player.Base.Location.Y == requiredY {
		return -1, game.player.Base
	}
	for place, monster := range game.monsterList {
		if monster.Location.X == requiredX && monster.Location.Y == requiredY {
			return place, monster
		}
	}
 	return -2, &Character.NPC{}
}

//this function will handle the fight event
func (game *Game) fight(character *Character.NPC, enemyX int, enemyY int, ) {
	place, enemy := game.findEnemy(enemyX, enemyY)
	damage := character.DoDamage()
	//include backstab bonus somewhere around here
	if rand.Intn(100) < enemy.Evasion {
		game.avoidAttackMessage(character.Name, enemy.Name)
	} else {
	enemy.TakeDamage(damage)
	game.takeDamageMessage(damage, character, enemy)
	}
	game.isCharacterDefeted(enemy, place)
}

func (game *Game) characterMoveTo(character *Character.NPC, x int, y int) {
	character.Location.X = x
	character.Location.Y = y
}

// //given a user input the function handles the desired action the player wants to performe
func (game *Game) plyerActionEvent(x int, y int, character *Character.NPC) {
	switch game.labyrinth.Labyrinth[x][y] {
	case Labyrinth.Pass:
		game.characterMoveTo(character, x, y)
	case Labyrinth.StartPosition:
		game.characterMoveTo(character, x, y)
	case Labyrinth.Monster:
		if character.IsHuman {
			game.fight(character, x, y)
		}
	case Labyrinth.CharSymbol:
		if !character.IsHuman {
			game.fight(character, x, y)
		}
	case Labyrinth.Trap:
		if character.IsHuman {
			game.encounterTrap(character, x, y)
		} else {
			game.characterMoveTo(character, x, y)
		}
	case Labyrinth.Treasure:
		//get the loot
	case Labyrinth.ExitPosition:
		//exit
	}
}

func (game *Game) restoreTile(x int, y int) {
	game.labyrinth.Labyrinth[x][y] = Labyrinth.Pass
	game.labyrinth.Labyrinth[game.start.X][game.start.Y] = Labyrinth.StartPosition
	game.labyrinth.Labyrinth[game.end.X][game.end.X] = Labyrinth.ExitPosition
}

//function takes user input and send it to PlayerActionEvent
func (game *Game) playerAction(key string) {
	game.restoreTile(game.player.Base.Location.X, game.player.Base.Location.Y)
	switch key {
	case "w":
		game.plyerActionEvent(game.player.Base.Location.X - 1, game.player.Base.Location.Y, game.player.Base)
	case "a":
		game.plyerActionEvent(game.player.Base.Location.X, game.player.Base.Location.Y - 1, game.player.Base)
	case "s":
		game.plyerActionEvent(game.player.Base.Location.X + 1, game.player.Base.Location.Y, game.player.Base)
	case "d":
		game.plyerActionEvent(game.player.Base.Location.X, game.player.Base.Location.Y + 1, game.player.Base)
	case "e":
		game.playerDefeted = true
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
	&weapon, &armor, 1,  10, 3, 5, 120.0, 120.0, 1.5, 30, 30, 0.2, 2, false, make(map[int]*Spells.Buff), false, 1}
}

// the function will create the traps in the maze
func (game *Game) createTrap(x int, y int) Character.Trap{
	trap := Character.Trap{}
	trap.Randomize(&Point.Point{x, y, nil})
	return trap
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

// //function replaces an element fro the 2d array for the maze with the character symbol
func (game *Game) drawCharacters() {
	game.labyrinth.Labyrinth[game.player.Base.Location.X][game.player.Base.Location.Y] = Labyrinth.CharSymbol
	for trapPoint, _ := range game.trapList {
		game.labyrinth.Labyrinth[trapPoint.X][trapPoint.Y] = Labyrinth.Trap
	}
	//add projectile here
	for _, mon := range game.monsterList {
		game.labyrinth.Labyrinth[mon.Location.X][mon.Location.Y] = mon.Symbol
	}
}

//function to draw the labyrinth
func (game *Game) drawLabyrinth() {
	for i := 0; i < game.labyrinth.Width; i++ {
		for j := 0; j < game.labyrinth.Height; j++ {
			if game.labyrinth.Labyrinth[i][j] == Labyrinth.Trap {
				trap := game.trapList[Point.Point{i, j, nil}]
				if trap.IsDetected {
					fmt.Print(game.labyrinth.Labyrinth[i][j])
				} else {
					fmt.Print(Labyrinth.Pass)
				}
			} else {
				fmt.Print(game.labyrinth.Labyrinth[i][j])
			}
		}
		fmt.Println()
	}
}

func (game *Game) detectKeyPress() string{
	reader := bufio.NewReader(os.Stdin)
	key, _ := reader.ReadString('\n')
	return strings.Trim(key,"\r\n")
}

func (game *Game) triggerDamageTrap(trap *Character.Trap, character *Character.NPC) {
	damage := trap.DamageTrap()
	if rand.Intn(100) < character.Evasion {
		game.avoidAttackMessage(character.Name, Character.TrapTypes[0])
	} else {
		character.TakeDamage(damage)
		game.takeDamageFromTrapMessage(damage, Character.TrapTypes[0], character)
	}
	game.isCharacterDefeted(character, -1)
}


func (game *Game) findEmptyTile(centerX int, centerY int) Point.Point{
	for e := 0; true; e++ {
		for i := centerX - 1 - e; i <= centerX + 1 + e; i++ {
			for j := centerY - 1 - e; j <= centerY + 1 + e; j++ {
				if i > -1 && j > -1 && game.labyrinth.Labyrinth[i][j] == Labyrinth.Pass {
					return Point.Point{i, j, nil}
				}
			}
		}
	}
	return Point.Point{}
}

func (game *Game) teleportMemoryWhipeTrap(trap *Character.Trap, hero *Character.Hero) {
	game.memoryWhipeTrap(trap, hero)
	game.teleportTrap(trap, hero.Base)
}

func (game *Game) memoryWhipeTrap(trap *Character.Trap, hero *Character.Hero) {
	fmt.Println("MEMORY TRAP")
	time.Sleep(2000 * time.Millisecond)
	trap.WhipeMemory(hero)
}

func (game *Game) teleportTrap(trap *Character.Trap, character *Character.NPC) {
	fmt.Println("TELEPORT TRAP")
	time.Sleep(2000 * time.Millisecond)
	location := trap.NewLocation(game.labyrinth.Width, game.labyrinth.Height)
	if game.labyrinth.Labyrinth[location.X][location.Y] == Labyrinth.Pass {
		character.Location.X = location.X
		character.Location.Y = location.Y
	} else {
		location = game.findEmptyTile(location.X, location.Y)
		character.Location.X = location.X
		character.Location.Y = location.Y
	}
}

func (game *Game) spawnTrap(trap *Character.Trap) {
	fmt.Println("SPAWNTRAP ACTIVATED")
	time.Sleep(2000 * time.Millisecond)
	location := trap.NewLocation(game.labyrinth.Width, game.labyrinth.Height)
	if game.labyrinth.Labyrinth[location.X][location.Y] == Labyrinth.Pass {
		newMonster := game.createMonster(location.X, location.Y)
		game.monsterList = append(game.monsterList, &newMonster)
	} else {
		location = game.findEmptyTile(location.X, location.Y)			
		newMonster := game.createMonster(location.X, location.Y)
		game.monsterList = append(game.monsterList, &newMonster)
	}
	fmt.Printf("Monster Spawned at %v,%v", location.X, location.Y)
	time.Sleep(2000 * time.Millisecond)
}

func (game *Game) triggerTrap(trap *Character.Trap, character *Character.Hero) {
	switch trap.TrapType {
	case Character.TrapTypes[0]:
		game.triggerDamageTrap(trap, character.Base)
	case Character.TrapTypes[1]:
		game.spawnTrap(trap)
	case Character.TrapTypes[2]:
		game.teleportTrap(trap, character.Base)
	case Character.TrapTypes[3]:
		game.memoryWhipeTrap(trap, character)
	case Character.TrapTypes[4]:
		game.teleportMemoryWhipeTrap(trap, character)
	}
}

func (game *Game) checkTraps() {
	if trap, ok := game.isTrapTriggered(game.player.Base) ; ok {
		game.triggerTrap(trap, game.player)
	}
}

func (game *Game) npcsTurn() {
	//spells
	game.checkTraps()
	//monsters
}

//main loop cycle for the game
func (game *Game) Run() {
	game.initialize()

	for  {
		c := exec.Command("cmd", "/c", "cls")
		c.Stdout = os.Stdout
		c.Run()
		game.drawCharacters()
		game.drawLabyrinth()
		key := game.detectKeyPress()
		game.playerAction(key)
		game.npcsTurn()
		if game.playerDefeted || game.gameCompleted {
			break
		}
	}

	//game.Evaluation()
}
