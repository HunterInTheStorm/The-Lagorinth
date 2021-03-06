//Handles the logic in the game.
package Game

import "github.com/golang/The-Lagorinth/Labyrinth"
import "github.com/golang/The-Lagorinth/Items"
import "github.com/golang/The-Lagorinth/Characters"
import "math/rand"
import "os"
import "bufio"
import "strings"

var upMovement string = "w"
var leftMovement string = "a"
var downMovement string = "s"
var rightMovement string = "d"
var exitCommand string = "exit"
var castSpellOne string = "1"
var castSpellTwo string = "2"
var castSpellThree string = "3"
var cameraMovementUp string = "8"
var cameraMovementDown string = "5"
var cameraMovementLeft string = "4"
var cameraMovementRight string = "6"
var cameraCenter string = "home"
var yes string = "y"
var no string = "n"

//playerAction determines what event follows depending on the key pressed.
func (game *Game) playerAction() {
	if !game.player.Base.IsStunned {
		key := game.detectKeyPress()
		switch key {
		case upMovement:
			game.plyerActionEvent(game.player.Base.Location.X-1, game.player.Base.Location.Y, game.player.Base)
			game.cameraReset()
		case leftMovement:
			game.plyerActionEvent(game.player.Base.Location.X, game.player.Base.Location.Y-1, game.player.Base)
			game.cameraReset()
		case downMovement:
			game.plyerActionEvent(game.player.Base.Location.X+1, game.player.Base.Location.Y, game.player.Base)
			game.cameraReset()
		case rightMovement:
			game.plyerActionEvent(game.player.Base.Location.X, game.player.Base.Location.Y+1, game.player.Base)
			game.cameraReset()
		case exitCommand:
			game.playerDefeted = true
		case castSpellOne:
			game.useSpell(game.player.SpellList[0], game.player)
		case castSpellTwo:
			game.useSpell(game.player.SpellList[1], game.player)
		case castSpellThree:
			game.useSpell(game.player.SpellList[2], game.player)
		case cameraMovementLeft:
			game.cameraMoveLeft()
			game.draw()
			game.playerAction()
		case cameraMovementDown:
			game.cameraMoveDown()
			game.draw()
			game.playerAction()
		case cameraMovementRight:
			game.cameraMoveRight()
			game.draw()
			game.playerAction()
		case cameraMovementUp:
			game.cameraMoveUp()
			game.draw()
			game.playerAction()
		case cameraCenter:
			game.cameraReset()
			game.draw()
			game.playerAction()
		}
	} else {
		game.player.Base.IsStunned = false
	}
	game.player.MemorizeLabyrinth(game.labyrinth, game.player.Base.Location)
}

//plyerActionEvent determines what event follows depending on the tile ahead.
func (game *Game) plyerActionEvent(x int, y int, character *Character.NPC) {
	if x >= 0 && y >= 0 && x < 40 && y < 40 {
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

//characterMoveTo handles updating the character coordinates when he moves.
//Changing the way he is facing.
//Drawing the character symbol on the tile ahead and restoring the tile he was on.
func (game *Game) characterMoveTo(character *Character.NPC, x int, y int) {
	game.restoreTile(character.Location.X, character.Location.Y)
	character.ChangeOrientation(x, y)
	character.Location.X = x
	character.Location.Y = y
	game.replaceTile(character.Location.X, character.Location.Y, character.Symbol)
}

//defeat handles the event when a character is defeated in battle.
func (game *Game) defeat(character *Character.NPC, place int) {
	if game.isCharacterDefeted(character) {
		if !character.IsHuman {
			game.monsterDefetedMessage(character.Name, game.player.Base.Name)
			game.lootEnemy(character)
		}
		game.CharacterDefeted(character, place)
	}
}

//fight handles the event of 2 characters fighting.
func (game *Game) fight(character *Character.NPC, enemyX int, enemyY int) {
	character.ChangeOrientation(enemyX, enemyY)
	place, enemy := game.findEnemy(enemyX, enemyY)
	damage := character.DoDamage()
	if rand.Intn(100) < enemy.Evasion {
		game.avoidAttackMessage(character.Name, enemy.Name)
	} else {
		enemy.TakeDamage(damage)
		enemy.ChangeOrientation(character.Location.X, character.Location.Y)
		game.takeDamageMessage(damage, character, enemy)
	}
	game.defeat(enemy, place)
}

//lootEnemy handles the precess of acquiring items from defeated foes.
func (game *Game) lootEnemy(character *Character.NPC) {
	weapon := character.Weapon
	character.UnequipWeapon()
	armor := character.Armor
	character.UnequipArmor()
	game.newArmorFound(armor)
	game.newWeaponFound(weapon)
}

//newArmorFound handles the event when a new armor has been found.
func (game *Game) newArmorFound(found *Items.Armor) {
	game.clearScreen()
	game.compareArmor(game.player.Base.Armor, found)
	game.newArmorFoundMessage()
	key := game.detectKeyPress()
	if key == yes {
		game.player.SwapArmor(found)
	}
}

//newWeaponFound handles the event when a new weapon has been found.
func (game *Game) newWeaponFound(found *Items.Weapon) {
	game.clearScreen()
	game.compareWeapon(game.player.Base.Weapon, found)
	game.newWeaponFoundMessage()
	key := game.detectKeyPress()
	if key == yes {
		game.player.SwapWeapon(found)
	}
}

//openChest handles the event of encountering a treasure tile.
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

//exitFound handles the event of encountering the exit tile.
func (game *Game) exitFound() {
	game.gameCompleted = true
}

//detectKeyPress return a string of the information entered via the keyboard.
func (game *Game) detectKeyPress() string {
	reader := bufio.NewReader(os.Stdin)
	key, _ := reader.ReadString('\n')
	return strings.Trim(key, "\r\n")
}

//applyRegenToAll calls the Regen function on every character.
func (game *Game) applyRegenToAll() {
	game.player.Base.Regenerate()
	for _, monster := range game.monsterList {
		monster.Regenerate()
	}
}

//isCharacterDefeted return true if a characters health point drop below 0.
func (game *Game) isCharacterDefeted(character *Character.NPC) bool {
	return character.CurrentHealth < 0
}

//CharacterDefeted handles the event when a character has been defeated.
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

//checkForDefeatedCharacters check if characters' health points have dropped below 0 from damage over time.
func (game *Game) checkForDefeatedCharacters() {
	for place, monster := range game.monsterList {
		game.defeat(monster, place)
	}
}

//moveMonsters calls the Move function for every character and the function to determine the following event.
func (game *Game) moveMonsters() {
	game.checkForDefeatedCharacters()
	for _, monster := range game.monsterList {
		if !monster.IsStunned {
			location := monster.Move(game.labyrinth)
			game.plyerActionEvent(location.X, location.Y, monster)
		} else {
			monster.IsStunned = false
		}
	}
}

//npcsTurn call function responsible for spell and monster movement, and for triggered traps.
func (game *Game) npcsTurn() {
	game.activateSpells()
	game.checkTraps()
	game.moveMonsters()
}

//findEnemy return a character and his number in the list.
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

//removeMonster removes a character from the character list.
func (game *Game) removeMonster(place int) {
	game.monsterList = append(game.monsterList[:place], game.monsterList[place+1:]...)
}
