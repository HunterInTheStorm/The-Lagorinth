package Game

import "github.com/golang/The-Lagorinth/Labyrinth"
import "github.com/golang/The-Lagorinth/Spells"
import "github.com/golang/The-Lagorinth/Characters"
import "math/rand"

func (game *Game) manageSpells(){
	game.lowerCoolDownOnSpells()
	game.lowerCharacterBuffDuration(game.player.Base)
	for _, monster := range game.monsterList {
		game.lowerCharacterBuffDuration(monster)
	}
}

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

func (game *Game) lowerCoolDownOnSpells() {
	for _, spell := range game.player.SpellList {
		spell.LowerCoolDownTime()
	}
}

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

func (game *Game) activateSpells() {
	for place, projectile := range game.projectileList {
		x := projectile.Location.X + projectile.Vector.X
		y := projectile.Location.Y + projectile.Vector.Y
		game.projectileActionEvent(x, y, projectile, place)
	}
}

func (game *Game) projectileMoveTo(projectile *Spell.Projectile, x int, y int) {
	game.restoreTile(projectile.Location.X, projectile.Location.Y)
	projectile.Move()
	game.replaceTile(projectile.Location.X, projectile.Location.Y, projectile.Symbol)
}

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

func (game *Game) removeProjectile(place int) {
	game.projectileList = append(game.projectileList[:place], game.projectileList[place +1:]...)
}

func (game *Game) findProjectile(x int, y int) int {
	for place, projectile := range game.projectileList {
		if x == projectile.Location.X && y == projectile.Location.Y {
			return place
		}
	}
	return -1
}

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
