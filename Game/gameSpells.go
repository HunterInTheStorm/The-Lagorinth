//Handles the logic in the game.
package Game

import "github.com/golang/The-Lagorinth/Labyrinth"
import "github.com/golang/The-Lagorinth/Spells"
import "github.com/golang/The-Lagorinth/Characters"
import "math/rand"

//manageSpells handles lowering the cool down on spells and the duration of buffs.
func (game *Game) manageSpells() {
	game.lowerCoolDownOnSpells()
	game.lowerCharacterBuffDuration(game.player.Base)
	for _, monster := range game.monsterList {
		game.lowerCharacterBuffDuration(monster)
	}
}

//lowerCharacterBuffDuration lower the duration of a character's buff.
func (game *Game) lowerCharacterBuffDuration(character *Character.NPC) {
	for _, buff := range character.BuffList {
		if buff.Duration > 0 {
			buff.LowerDuration()
		} else {
			character.RemoveBuff(buff)
			if character.IsHuman {
				game.buffFadeMessage(buff)
			}
			delete(character.BuffList, buff.BuffID)
		}
	}
}

//lowerCoolDownOnSpells lowers the cool down on spells.
func (game *Game) lowerCoolDownOnSpells() {
	for _, spell := range game.player.SpellList {
		spell.LowerCoolDownTime()
	}
}

//projectileActionEvent handles the interaction between projectiles and the surrounding world.
func (game *Game) projectileActionEvent(x int, y int, projectile *Spell.Projectile, place int) {
	switch game.labyrinth.Labyrinth[x][y] {
	case Labyrinth.Pass:
		game.projectileMoveTo(projectile, x, y)
	case Labyrinth.StartPosition:
		game.projectileMoveTo(projectile, x, y)
	case Labyrinth.ExitPosition:
		game.projectileMoveTo(projectile, x, y)
	case Labyrinth.Trap:
		game.projectileMoveTo(projectile, x, y)
	case Labyrinth.Treasure:
		game.projectileMoveTo(projectile, x, y)
		projectile.ProjectileImapact(game.labyrinth)
		game.removeProjectile(place)
	case Labyrinth.Projectile:
		spot := game.findProjectile(x, y)
		game.removeProjectile(place)
		game.removeProjectile(spot)
	case Labyrinth.Wall:
		game.projectileMoveTo(projectile, x, y)
		projectile.ProjectileImapact(game.labyrinth)
		game.removeProjectile(place)
	case Labyrinth.Monster:
		game.projectileMoveTo(projectile, x, y)
		game.projectileHitsCharacter(x, y, projectile)
		game.removeProjectile(place)
	case Labyrinth.CharSymbol:
		game.projectileMoveTo(projectile, x, y)
		game.projectileHitsCharacter(x, y, projectile)
		game.removeProjectile(place)
	}
}

//activateSpells determines where the projectile will move and call projectileActionEvent function.
func (game *Game) activateSpells() {
	for place, projectile := range game.projectileList {
		x := projectile.Location.X + projectile.Vector.X
		y := projectile.Location.Y + projectile.Vector.Y
		game.projectileActionEvent(x, y, projectile, place)
	}
}

//projectileMoveTo moves the projectile forward, deleting is symbol from it's previous tile and placing it on its new one.
func (game *Game) projectileMoveTo(projectile *Spell.Projectile, x int, y int) {
	game.restoreTile(projectile.Location.X, projectile.Location.Y)
	projectile.Move()
	game.replaceTile(projectile.Location.X, projectile.Location.Y, projectile.Symbol)
}

//projectileHitsCharacter handles the event when a projectile hits a character.
func (game *Game) projectileHitsCharacter(x int, y int, projectile *Spell.Projectile) {
	place, enemy := game.findEnemy(x, y)
	enemy.ProjectileToTheFace(projectile)
	damage := projectile.DoDamage()
	if rand.Intn(100) < enemy.Evasion {
		game.avoidSpellMessage(enemy)
	} else {
		enemy.TakeDamage(damage)
		game.takeSpellDamageMessage(damage, enemy)
	}
	if game.isCharacterDefeted(enemy) {
		game.CharacterDefeted(enemy, place)
	}
}

//removeProjectile removes a projectile from the list of projectiles.
func (game *Game) removeProjectile(place int) {
	game.projectileList = append(game.projectileList[:place], game.projectileList[place+1:]...)
}

//findProjectile find a projectile in the list given a set of coordinates.
func (game *Game) findProjectile(x int, y int) int {
	for place, projectile := range game.projectileList {
		if x == projectile.Location.X && y == projectile.Location.Y {
			return place
		}
	}
	return -1
}

//useSpell determines the type of the spell cast and call appropriate functions.
func (game *Game) useSpell(spell *Spell.Spell, hero *Character.Hero) {
	if spell.ManaCost < hero.Base.CurrentMana {
		if !spell.IsOnCoolDown {
			if spell.IsSelfTargeted && !spell.IsBuff {
				hero.UseInstantSpell(spell)
				game.useInstantSpellMessage(spell, hero)
				spell.GoOnCoolDown()
			} else if spell.IsSelfTargeted && spell.IsBuff {
				hero.UseBuffSpell(spell)
				game.useBuffSpellMessage(spell, hero)
				spell.GoOnCoolDown()
			} else if spell.IsProjectile {
				projectile := hero.UseProjectileSpell(spell)
				game.projectileList = append(game.projectileList, projectile)
				spell.GoOnCoolDown()
			} else if spell.IsAreaOfEffect {
			}
		} else {
			game.spellOnCoolDownMessage(spell)
		}
	} else {
		game.lowManaMessage(spell)
	}
}
