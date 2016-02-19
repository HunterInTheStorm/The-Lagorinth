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
	camera *Point.Point
	cameraRadius int
	projectileList []*Spell.Projectile
}

func (game *Game) useInstantSpellMessage(spell *Spell.Spell, hero *Character.Hero) {
	fmt.Printf("%s uses %s.", hero.Base.Name, spell.SpellName)
	if spell.ManaCost > 0 {
		fmt.Printf("%s restores %v HP for %v MP", hero.Base.Name, spell.RegainHealth, spell.ManaCost)
	} else{
		fmt.Printf("%s restores %v MP for %v HP", hero.Base.Name, spell.ManaCost, spell.RegainHealth)
	}
	time.Sleep(2000 * time.Millisecond)
}

func (game *Game) spellOnCoolDownMessage(spell *Spell.Spell) {
	if spell.CoolDownTimeLeft != 1 {
		fmt.Printf("%s is on CD for %v turns", spell.SpellName, spell.CoolDownTimeLeft)
	} else {
		fmt.Printf("%s is on CD for %v turn", spell.SpellName, spell.CoolDownTimeLeft)
	}
	time.Sleep(2000 * time.Millisecond)
}

func (game *Game) useBuffSpellMessage(spell *Spell.Spell, hero *Character.Hero) {
	fmt.Printf("%s is now buffed with %s for %v turns", hero.Base.Name, spell.SpellName, spell.Duration)
	time.Sleep(2000 * time.Millisecond)
}

func (game *Game) buffFadeMessage(spell *Spell.Buff) {
	fmt.Printf("%s fades away", spell.BuffName)
	time.Sleep(2000 * time.Millisecond)
}

func (game *Game) lowManaMessage(spell *Spell.Spell) {
	fmt.Printf("Not enough MP to cast %s. Spell Requires %v MP", spell.SpellName, spell.ManaCost)
	time.Sleep(2000 * time.Millisecond)
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

func (game *Game) isCharacterDefeted(character *Character.NPC) bool {
	return character.CurrentHealth < 0
}

func (game *Game) CharacterDefeted(character *Character.NPC, place int) {
	switch true {
	case place == -1:
		game.playerDefeted = true
		game.playerDefetedMessage()
	case place > -1:
		game.removeMonster(place)
		game.monsterSlain++
		game.restoreTile(character.Location.X, character.Location.Y)
	}
}

func (game *Game) avoidAttackMessage(attackerName string, defenderName string) {
	fmt.Printf("%s dodges %s's attack",defenderName, attackerName)
	time.Sleep(2000 * time.Millisecond)
}

func (game *Game) takeDamageFromTrapMessage(damage float32, attackerName string, defender *Character.NPC) {
	fmt.Printf("%s strikes %s for %v points of damage.\n", attackerName, defender.Name,
						 int(damage))
	time.Sleep(2000 * time.Millisecond)
}

func (game *Game) takeDamageMessage(damage float32, attacker *Character.NPC, defender *Character.NPC) {
	fmt.Printf("%s strikes %s for %v points of damage. %s has %v HP left\n", attacker.Name, defender.Name,
						 int(damage),defender.Name, int(defender.CurrentHealth))
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

func (game *Game) lootEnemy(character *Character.NPC) {
	weapon := character.Weapon
	character.UnequipWeapon()
	armor := character.Armor
	character.UnequipArmor()
	game.newArmorFound(armor)
	game.newWeaponFound(weapon)
}

//this function will handle the fight event
func (game *Game) fight(character *Character.NPC, enemyX int, enemyY int, ) {
	character.ChangeOrientation(enemyX, enemyY)
	place, enemy := game.findEnemy(enemyX, enemyY)
	damage := character.DoDamage()
	//include backstab bonus somewhere around here
	if rand.Intn(100) < enemy.Evasion {
		game.avoidAttackMessage(character.Name, enemy.Name)
	} else {
	enemy.TakeDamage(damage)
	enemy.ChangeOrientation(character.Location.X, character.Location.Y)
	game.takeDamageMessage(damage, character, enemy)
	}
	if game.isCharacterDefeted(enemy) {
		if !enemy.IsHuman {
			game.monsterDefetedMessage(character.Name, game.player.Base.Name)
			game.lootEnemy(enemy)
		}
		game.CharacterDefeted(enemy, place)
	}
}

func (game *Game) characterMoveTo(character *Character.NPC, x int, y int) {
	game.restoreTile(character.Location.X, character.Location.Y)
	character.ChangeOrientation(x, y)
	character.Location.X = x
	character.Location.Y = y
	game.replaceTile(character.Location.X, character.Location.Y, character.Symbol)
}

// //given a user input the function handles the desired action the player wants to performe
func (game *Game) plyerActionEvent(x int, y int, character *Character.NPC) {
	if x >= 0 && y >=0 && x < 40 && y < 40 {
		switch game.labyrinth.Labyrinth[x][y] {
		case Labyrinth.Pass:
			if character.IsHuman {
				game.characterMoveTo(character, x, y)
				game.drawHero()
			} else {
				game.characterMoveTo(character, x, y)
			}
		case Labyrinth.StartPosition:
			if character.IsHuman {
				game.characterMoveTo(character, x, y)
				game.drawHero()
			} else {
				game.characterMoveTo(character, x, y)
			}
		case Labyrinth.Monster:
			if character.IsHuman {
				game.fight(character, x, y)
			}
		case Labyrinth.CharSymbol:
			if !character.IsHuman {
				game.draw()
				game.fight(character, x, y)
			}
		case Labyrinth.Trap:
			if character.IsHuman {
				game.encounterTrap(character, x, y)
			} else {
				game.characterMoveTo(character, x, y)
				game.drawHero()
			}
		case Labyrinth.Treasure:
			if character.IsHuman {
				game.openChest()
				game.characterMoveTo(character, x, y)
			}
		case Labyrinth.ExitPosition:
			if character.IsHuman {
				game.exitFound()
				game.characterMoveTo(character, x, y)
				game.drawHero()
			} else {
				game.characterMoveTo(character, x, y)
			}
		}
	}
}

func (game *Game) compareArmor(current *Items.Armor, found *Items.Armor) {
	fmt.Println("Your Armor\t\t\tNew Armor")
	fmt.Printf("+Health: %v\t\t\t+Health: %v\n",current.Health, found.Health)
	fmt.Printf ("+HP Regen: %v\t\t\t+HP Regen: %v\n",current.HealthRegen, found.HealthRegen)
	fmt.Printf("+Mana: %v\t\t\t+Mana: %v\n",current.Mana, found.Mana)
	fmt.Printf ("+MP Regen: %v\t\t\t+MP Regen: %v\n",current.ManaRegen, found.ManaRegen)
	fmt.Printf("+Defence: %v\t\t\t+Defence: %v\n",current.Defence, found.Defence)
	fmt.Printf("+Evasion: %v\t\t\t+Evasion: %v\n",current.Evasion, found.Evasion)
}

func (game *Game) compareWeapon(current *Items.Weapon, found *Items.Weapon) {
	fmt.Println("Your Weapon\t\t\tNew Weapon")
	fmt.Printf("Damage: %v-%v\t\t\tDamage: %v-%v\n", current.MinDmg, current.MaxDmg, found.MinDmg, found.MaxDmg)
	fmt.Printf("+BonusDmg: %v\t\t\t+BonusDmg: %v\n",current.BonusDmg, found.BonusDmg)
	fmt.Printf ("+CritChance: %v\t\t\t+CritChance: %v\n",current.BonusCritChance, found.BonusCritChance)
}

func (game *Game) newArmorFoundMessage() {
	fmt.Println("You found a new armor piece! :)\nWould you like to equip it?(y/n)")
}

func (game *Game) newWeaponFoundMessage() {
	fmt.Println("You found a new weapon! :)\nWould you like to equip it?(y/n)")
}

func (game *Game) newArmorFound(found *Items.Armor) {
	game.clearScreen()
	game.compareArmor(game.player.Base.Armor, found)
	game.newArmorFoundMessage()
	key := game.detectKeyPress()
	if key == "y" {
		game.player.SwapArmor(found)
	}
}

func (game *Game) newWeaponFound(found *Items.Weapon) {
	game.clearScreen()
	game.compareWeapon(game.player.Base.Weapon, found)
	game.newWeaponFoundMessage()
	key := game.detectKeyPress()
	if key == "y" {
		game.player.SwapWeapon(found)
	}
}

func (game *Game) openChest() {
	if rand.Intn(2) == 0 {
		armor := game.createArmor()
		game.newArmorFound(armor)
	} else {
		weapon := game.createWeapon()
		game.newWeaponFound(weapon)
	}
	game.chestsLooted++
}

func (game *Game) exitFound() {
	game.gameCompleted = true
}

func (game *Game) restoreTile(x int, y int) {
	game.labyrinth.Labyrinth[x][y] = Labyrinth.Pass
	if game.start.X != game.player.Base.Location.X || game.start.Y != game.player.Base.Location.Y {
		game.labyrinth.Labyrinth[game.start.X][game.start.Y] = Labyrinth.StartPosition
	}
	if game.end.X != game.player.Base.Location.X && game.end.Y != game.player.Base.Location.Y {
		game.labyrinth.Labyrinth[game.end.X][game.end.X] = Labyrinth.ExitPosition
	}
}

func(game *Game) replaceTile(x int, y int, symbol string) {
	if game.labyrinth.Labyrinth[x][y] != Labyrinth.Wall {
		game.labyrinth.Labyrinth[x][y] = symbol
	}
}

//function takes user input and send it to PlayerActionEvent
func (game *Game) playerAction() {
	if !game.player.Base.IsStunned {
		key := game.detectKeyPress()
		switch key {
		case "w":
			game.plyerActionEvent(game.player.Base.Location.X - 1, game.player.Base.Location.Y, game.player.Base)
			game.cameraReset()
		case "a":
			game.plyerActionEvent(game.player.Base.Location.X, game.player.Base.Location.Y - 1, game.player.Base)
			game.cameraReset()
		case "s":
			game.plyerActionEvent(game.player.Base.Location.X + 1, game.player.Base.Location.Y, game.player.Base)
			game.cameraReset()
		case "d":
			game.plyerActionEvent(game.player.Base.Location.X, game.player.Base.Location.Y + 1, game.player.Base)
			game.cameraReset()
		case "exit":
			game.playerDefeted = true
		case "1":
			game.useSpell(game.player.SpellList[0], game.player)
		case "2":
			game.useSpell(game.player.SpellList[1], game.player)
		case "3":
			game.useSpell(game.player.SpellList[2], game.player)
		case "4":
			game.cameraMoveLeft()
			game.draw()
			game.playerAction()
		case "5":
			game.cameraMoveDown()
			game.draw()
			game.playerAction()
		case "6":
			game.cameraMoveRight()
			game.draw()
			game.playerAction()
		case "8":
			game.cameraMoveUp()
			game.draw()
			game.playerAction()
		case "home":
			game.cameraReset()
			game.draw()
			game.playerAction()
		}
	} else {
		game.player.Base.IsStunned = false
	}
	game.player.MemorizeLabyrinth(game.labyrinth, game.player.Base.Location)
}

func (game *Game) cameraMoveLeft() {
	game.camera.Y--
}

func (game *Game) cameraMoveDown() {
	game.camera.X++
}

func (game *Game) cameraMoveRight() {
	game.camera.Y++
}

func (game *Game) cameraMoveUp() {
	game.camera.X--
}

func (game *Game) cameraReset() {
	game.camera.X = game.player.Base.Location.X
	game.camera.Y = game.player.Base.Location.Y
}
// //function will create a new monster at random intervals of time
// func (game *Game) SpawnMonster() {

// }

// //function will set the win condotion when the end of the maze has benn found
// func (game *Game) Evaluation() {
// 	return true
// }

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

func (game *Game) initialize() {
	game.setGameFieldValues()
	game.createLabyrinth()
	game.createMonsterAndTrapsLists()
	//TRY TO DO SOME MAGIC WITH CHANNELS BY SENDING POINTS WITH LOCATION OF START,END,NPCs,TRAPs
	game.createHero()
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
	if game.isCharacterDefeted(character) {
		game.CharacterDefeted(character, -1)
	}
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

func (game *Game) triggerTabulaRasaTrap(trap *Character.Trap, hero *Character.Hero) {
	game.triggerMemoryWhipeTrap(trap, hero)
	game.triggerTeleportTrap(trap, hero.Base)
}

func (game *Game) triggerMemoryWhipeTrap(trap *Character.Trap, hero *Character.Hero) {
	fmt.Println("MEMORY TRAP")
	time.Sleep(2000 * time.Millisecond)
	trap.WhipeMemory(hero)
}

func (game *Game) triggerTeleportTrap(trap *Character.Trap, character *Character.NPC) {
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
	game.cameraReset()
}

func (game *Game) triggerSpawnTrap(trap *Character.Trap) {
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
		game.triggerSpawnTrap(trap)
	case Character.TrapTypes[2]:
		game.triggerTeleportTrap(trap, character.Base)
		game.player.MemorizeLabyrinth(game.labyrinth, game.player.Base.Location)
		game.draw()
	case Character.TrapTypes[3]:
		game.triggerMemoryWhipeTrap(trap, character)
		game.player.MemorizeLabyrinth(game.labyrinth, game.player.Base.Location)
		game.draw()
	case Character.TrapTypes[4]:
		game.triggerTabulaRasaTrap(trap, character)
		game.player.MemorizeLabyrinth(game.labyrinth, game.player.Base.Location)
		game.draw()
	}
}

func (game *Game) checkTraps() {
	if trap, ok := game.isTrapTriggered(game.player.Base) ; ok {
		game.triggerTrap(trap, game.player)
	}
}

func (game *Game) moveMonsters() {
	for _, monster := range game.monsterList {
		if !monster.IsStunned {
			location := monster.Move(game.labyrinth)
			game.plyerActionEvent(location.X, location.Y, monster)
		} else {
			monster.IsStunned = false
		}
	}
}

func (game *Game) takeSpellDamageMessage(damage float32, character *Character.NPC) {
	fmt.Printf("%s is hit for %v spell damage", character.Name, int(damage))
	time.Sleep(2000 * time.Millisecond)
}

func (game *Game) avoidSpellMessage(character *Character.NPC) {
	fmt.Printf("%s dodges spell", character.Name)
	time.Sleep(2000 * time.Millisecond)
}

func (game *Game) npcsTurn() {
	game.activateSpells()
	game.checkTraps()
	game.moveMonsters()
}

func (game *Game) applyRegenToAll() {
	game.player.Base.Regenerate()
	for _, monster := range game.monsterList {
		monster.Regenerate()
	}
}

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
