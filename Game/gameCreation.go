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

func (game *Game) initialize() {
	game.setGameFieldValues()
	game.createLabyrinth()
	game.createMonsterAndTrapsLists()
	//TRY TO DO SOME MAGIC WITH CHANNELS BY SENDING POINTS WITH LOCATION OF START,END,NPCs,TRAPs
	game.createHero()
}

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

func (game *Game) createLabyrinth() {
	var seed int64 = time.Now().UTC().UnixNano()
	var maze [40][40]string
	for i := 0; i < 40; i++ {
		for j := 0; j < 40; j++ {
			maze[i][j] = Labyrinth.Wall
		}
	}
	game.labyrinth = &Labyrinth.Labyrinth{40,40, maze}
	game.labyrinth.CreateLabyrinth(seed)
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

// the function will create the monsters in the maze
func (game *Game) createMonster(x int, y int) Character.NPC{
	weapon := game.createWeapon()
	armor := game.createArmor()
	//TRANSFER VALUES TO SEPARATE FILE
	monster := Character.NPC{&Point.Point{x, y, nil}, Labyrinth.Monster, "Skeleton", &Point.Point{-1, 0, nil},
	nil, nil, 2.5,  10, 3, 5, 120.0, 120.0, 1.5, 30, 30, 0.2, 2, false, make(map[int]*Spell.Buff), false, 1}
	monster.EquipArmor(armor)
	monster.EquipWeapon(weapon)
	return monster
}

// the function will create the traps in the maze
func (game *Game) createTrap(x int, y int) Character.Trap{
	trap := Character.Trap{}
	trap.Randomize(&Point.Point{x, y, nil})
	return trap
}

func(game *Game) createArmor() *Items.Armor {
	armor := Items.Armor{}
	armor.RandomizeArmor()
	return &armor
}

func(game *Game) createWeapon() *Items.Weapon {
	weapon := Items.Weapon{}
	weapon.RandomizeWeapon()
	return &weapon
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
	game.addBackground(charBackGround, game.player)
	game.player.MemorizeLabyrinth(game.labyrinth, game.player.Base.Location)
}

func(game *Game) chooseBackground() string {
	game.availableBackGroundsMessage()
	reader := bufio.NewReader(os.Stdin)
	backGround, _ := reader.ReadString('\n')
	return strings.Trim(backGround,"\r\n")
}

func (game *Game) createEquipment(hero *Character.Hero) {
	weapon := game.createWeapon()
	armor := game.createArmor()
	hero.SwapWeapon(weapon)
	hero.SwapArmor(armor)
}

// //this function will create one of the 3 classes for the player
func (game *Game) createPaladin(charName string, charBackGround string) {
	hero := Character.CreatePaladin(charName, charBackGround, game.start.X, game.start.Y)
	game.createEquipment(hero)
	game.player = hero
}

// //this function will create one of the 3 classes for the player
func (game *Game) createMage(charName string, charBackGround string) {
	hero := Character.CreateMage(charName, charBackGround, game.start.X, game.start.Y)
	game.createEquipment(hero)
	game.player = hero
}

// //this function will create one of the 3 classes for the player
func (game *Game) createRouge(charName string, charBackGround string) {
	hero := Character.CreateRouge(charName, charBackGround, game.start.X, game.start.Y)
	game.createEquipment(hero)
	game.player = hero
}

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

func(game *Game) chooseName() string {
	game.askNameMessage()
 	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	return strings.Trim(name,"\r\n")
}

func(game *Game) chooseClass() string {
	game.askClassNameMessage()
	reader := bufio.NewReader(os.Stdin)
	class, _ := reader.ReadString('\n')
	return strings.Trim(class,"\r\n")
}
