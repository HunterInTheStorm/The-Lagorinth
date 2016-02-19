package Game

import "github.com/golang/The-Lagorinth/Labyrinth"
import "github.com/golang/The-Lagorinth/Items"
import "github.com/golang/The-Lagorinth/Characters"
import "math/rand"
import "os"
import "bufio"
import "strings"

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

func (game *Game) characterMoveTo(character *Character.NPC, x int, y int) {
	game.restoreTile(character.Location.X, character.Location.Y)
	character.ChangeOrientation(x, y)
	character.Location.X = x
	character.Location.Y = y
	game.replaceTile(character.Location.X, character.Location.Y, character.Symbol)
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

func (game *Game) lootEnemy(character *Character.NPC) {
	weapon := character.Weapon
	character.UnequipWeapon()
	armor := character.Armor
	character.UnequipArmor()
	game.newArmorFound(armor)
	game.newWeaponFound(weapon)
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

func (game *Game) detectKeyPress() string{
	reader := bufio.NewReader(os.Stdin)
	key, _ := reader.ReadString('\n')
	return strings.Trim(key,"\r\n")
}

func (game *Game) applyRegenToAll() {
	game.player.Base.Regenerate()
	for _, monster := range game.monsterList {
		monster.Regenerate()
	}
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

func (game *Game) npcsTurn() {
	game.activateSpells()
	game.checkTraps()
	game.moveMonsters()
}

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

func (game *Game) removeMonster(place int) {
	game.monsterList = append(game.monsterList[:place], game.monsterList[place +1:]...)
}
