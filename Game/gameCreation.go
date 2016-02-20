//Handles the logic in the game.
package Game

import "github.com/golang/The-Lagorinth/Labyrinth"
import "github.com/golang/The-Lagorinth/Items"
import "github.com/golang/The-Lagorinth/Spells"
import "github.com/golang/The-Lagorinth/Characters"
import "github.com/golang/The-Lagorinth/Point"
import "time"
import "os"
import "bufio"
import "strings"

const dimensionX int = 40
const dimensionY int = 40

//initialize handles setting the game field values and creating the objects used in the game.
//Call other functions.
func (game *Game) initialize() {
	game.setGameFieldValues()
	game.createLabyrinth()
	game.createMonsterAndTrapsLists()
	//TRY TO DO SOME MAGIC WITH CHANNELS BY SENDING POINTS WITH LOCATION OF START,END,NPCs,TRAPs
	game.createHero()
}

//setGameFieldValues sets the game's field values.
func (game *Game) setGameFieldValues() {
	game.playerDefeted = false
	game.gameCompleted = false
	game.score = 1000
	game.turns = 0
	game.monsterSlain = 0
	game.chestsLooted = 0
	game.trapsDisarmed = 0
	game.cameraRadius = 12
	game.projectileList = make([]*Spell.Projectile, 0, 8)
}

//createLabyrinth creates the labyrinth.
func (game *Game) createLabyrinth() {
	var seed int64 = time.Now().UTC().UnixNano()
	var maze [dimensionX][dimensionY]string
	for i := 0; i < dimensionX; i++ {
		for j := 0; j < dimensionY; j++ {
			maze[i][j] = Labyrinth.Wall
		}
	}
	game.labyrinth = &Labyrinth.Labyrinth{dimensionX, dimensionY, maze}
	game.labyrinth.CreateLabyrinth(seed)
}

//createMonsterAndTrapsLists creates monsters and traps and places them in a respective list.
func (game *Game) createMonsterAndTrapsLists() {
	game.monsterList = make([]*Character.NPC, 0, 4)
	game.trapList = make(map[Point.Point]*Character.Trap)
	for i := 0; i < dimensionX; i++ {
		for j := 0; j < dimensionY; j++ {
			tile := game.labyrinth.Labyrinth[i][j]
			if tile != Labyrinth.Wall {
				switch tile {
				case Labyrinth.StartPosition:
					game.start = &Point.Point{i, j, nil}
					game.camera = &Point.Point{i, j, nil}
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

//createMonster handles the creation of a character.
func (game *Game) createMonster(x int, y int) Character.NPC {
	weapon := game.createWeapon()
	armor := game.createArmor()
	//TRANSFER VALUES TO SEPARATE FILE
	monster := Character.NPC{&Point.Point{x, y, nil}, Labyrinth.Monster, "Skeleton", &Point.Point{-1, 0, nil},
		nil, nil, 2.5, 10, 3, 5, 120.0, 120.0, 1.5, 30, 30, 0.2, 2, false, make(map[int]*Spell.Buff), false, 1}
	monster.EquipArmor(armor)
	monster.EquipWeapon(weapon)
	return monster
}

//createTrap handles the creation of a trap.
func (game *Game) createTrap(x int, y int) Character.Trap {
	trap := Character.Trap{}
	trap.Randomize(&Point.Point{x, y, nil})
	return trap
}

//createArmor handles the creation of an armor.
func (game *Game) createArmor() *Items.Armor {
	armor := Items.Armor{}
	armor.RandomizeArmor()
	return &armor
}

//createWeapon handles the creation of an weapon.
func (game *Game) createWeapon() *Items.Weapon {
	weapon := Items.Weapon{}
	weapon.RandomizeWeapon()
	return &weapon
}

//createHero handles the process of creating the player's character.
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
	game.addBackground(charBackGround, game.player)
	game.player.MemorizeLabyrinth(game.labyrinth, game.player.Base.Location)
}

//chooseBackground returns the name of the background the player has chosen.
func (game *Game) chooseBackground() string {
	game.availableBackGroundsMessage()
	reader := bufio.NewReader(os.Stdin)
	backGround, _ := reader.ReadString('\n')
	return strings.Trim(backGround, "\r\n")
}

//createEquipment create weapon and armor for the player's character and equips them.
func (game *Game) createEquipment(hero *Character.Hero) {
	weapon := game.createWeapon()
	armor := game.createArmor()
	hero.SwapWeapon(weapon)
	hero.SwapArmor(armor)
}

//createPaladin creates Paladin class character.
func (game *Game) createPaladin(charName string, charBackGround string) {
	hero := Character.CreatePaladin(charName, charBackGround, game.start.X, game.start.Y)
	game.createEquipment(hero)
	game.player = hero
}

//createMage creates Mage class character.
func (game *Game) createMage(charName string, charBackGround string) {
	hero := Character.CreateMage(charName, charBackGround, game.start.X, game.start.Y)
	game.createEquipment(hero)
	game.player = hero
}

//createRouge creates Rouge class character.
func (game *Game) createRouge(charName string, charBackGround string) {
	hero := Character.CreateRouge(charName, charBackGround, game.start.X, game.start.Y)
	game.createEquipment(hero)
	game.player = hero
}

//addBackground	determines what background to apply to the player's character.
func (game *Game) addBackground(charBackGround string, hero *Character.Hero) {
	switch charBackGround {
	case Character.BackGroundNameGiant:
		backGround := Character.CreateBgGiant()
		hero.ApplyBackground(backGround)
	case Character.BackGroundNameToreador:
		backGround := Character.CreateBgToreador()
		hero.ApplyBackground(backGround)
	case Character.BackGroundNameCartographer:
		backGround := Character.CreateBgCartographer()
		hero.ApplyBackground(backGround)
	case Character.BackGroundNameLibrarian:
		backGround := Character.CreateBgLibrarian()
		hero.ApplyBackground(backGround)
	}
}

//chooseName return a string of the desired name for the player's character.
func (game *Game) chooseName() string {
	game.askNameMessage()
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	return strings.Trim(name, "\r\n")
}

//chooseClass return a string of the desired class for the player's character.
func (game *Game) chooseClass() string {
	game.askClassNameMessage()
	reader := bufio.NewReader(os.Stdin)
	class, _ := reader.ReadString('\n')
	return strings.Trim(class, "\r\n")
}
