//package Spell handles the creation and functionalities of the spells in the game.
package Spell

import "github.com/golang/The-Lagorinth/Point"
import "github.com/golang/The-Lagorinth/Labyrinth"

//RougeSpellPrecision creates and return a spell.
//It takes a Point structure as an argument, to be used for coordinates of future created projectiles.
//The spell is created for the use of the Rouge class in the game.
//Function makes it easier to modify and balance spells on the fly.
func RougeSpellPrecision(point *Point.Point) *Spell {
	//All spells have this
	origin := point
	var symbol string = ""
	var spellName string = "Precision"
	var isSelfTargeted bool = true
	var isProjectile bool = false
	var isAreaOfEffect bool = false
	var isBuff bool = true
	var cooldownTime int = 10
	var isOnCoolDown bool = false
	var coolDownTimeLeft int = 0
	//instant spells
	var regainHealth float32 = 0.0
	var manaCost float32 = 30.0
	//buff
	var buffId int = id
	var manaCostPerSecond float32 = 5.0
	var bonusHealthRegen float32 = 0.0
	var bonusDamageMultiplier float32 = 2.0
	var duration int = 5
	var bonusDamage int = 0
	var bonusDefence int = 0
	var bonusEvasion int = 0
	var bonusCritChance int = 15
	var willApplyToEnemies bool = false
	//AoE has buff
	var damage int = 0
	var radius int = 0
	//Projectile has damage and buff
	var willStun bool = false
	var canDestroyWall bool = false

	spell := Spell{origin, symbol, spellName, isSelfTargeted, isProjectile, isAreaOfEffect,
		isBuff, regainHealth, bonusHealthRegen, manaCost, manaCostPerSecond, bonusDamageMultiplier,
		radius, duration, damage, bonusDamage, bonusDefence, bonusEvasion, bonusCritChance, willStun,
		canDestroyWall, willApplyToEnemies, cooldownTime, coolDownTimeLeft, isOnCoolDown, buffId}

	id++
	return &spell
}

//RougeSpellShadow creates and return a spell.
//It takes a Point structure as an argument, to be used for coordinates of future created projectiles.
//The spell is created for the use of the Rouge class in the game.
//Function makes it easier to modify and balance spells on the fly.
func RougeSpellShadow(point *Point.Point) *Spell {
	//All spells have this
	origin := point
	var symbol string = ""
	var spellName string = "Shadow"
	var isSelfTargeted bool = true
	var isProjectile bool = false
	var isAreaOfEffect bool = false
	var isBuff bool = true
	var cooldownTime int = 10
	var isOnCoolDown bool = false
	var coolDownTimeLeft int = 0
	//instant spell
	var regainHealth float32 = 0.0
	var manaCost float32 = 30.0
	//buff
	var buffId int = id
	var manaCostPerSecond float32 = 6.0
	var bonusHealthRegen float32 = -3.5
	var bonusDamageMultiplier float32 = 0.0
	var duration int = 10
	var bonusDamage int = 0
	var bonusDefence int = 3
	var bonusEvasion int = 15
	var bonusCritChance int = 0
	var willApplyToEnemies bool = false
	//AoE has buff
	var damage int = 0
	var radius int = 0
	//Projectile has damage and buff
	var willStun bool = false
	var canDestroyWall bool = false

	spell := Spell{origin, symbol, spellName, isSelfTargeted, isProjectile, isAreaOfEffect,
		isBuff, regainHealth, bonusHealthRegen, manaCost, manaCostPerSecond, bonusDamageMultiplier,
		radius, duration, damage, bonusDamage, bonusDefence, bonusEvasion, bonusCritChance, willStun,
		canDestroyWall, willApplyToEnemies, cooldownTime, coolDownTimeLeft, isOnCoolDown, buffId}

	id++
	return &spell
}

//RougeSpellAssassinMark creates and return a spell.
//It takes a Point structure as an argument, to be used for coordinates of future created projectiles.
//The spell is created for the use of the Rouge class in the game.
//Function makes it easier to modify and balance spells on the fly.
func RougeSpellAssassinMark(point *Point.Point) *Spell {
	//All spells have this
	origin := point
	var symbol string = Labyrinth.Projectile
	var spellName string = "Assassin's Mark"
	var isSelfTargeted bool = false
	var isProjectile bool = true
	var isAreaOfEffect bool = false
	var isBuff bool = true
	var cooldownTime int = 20
	var isOnCoolDown bool = false
	var coolDownTimeLeft int = 0
	//instant spell
	var regainHealth float32 = 0.0
	var manaCost float32 = 25.0
	//buff
	var buffID int = id
	var manaCostPerSecond float32 = 0.0
	var bonusHealthRegen float32 = -5.0
	var bonusDamageMultiplier float32 = 0.0
	var duration int = 5
	var bonusDamage int = -5
	var bonusDefence int = -10
	var bonusEvasion int = -3
	var bonusCritChance int = 0
	var willApplyToEnemies bool = true
	//AoE has buff
	var damage int = 0
	var radius int = 0
	//Projectile has damage and buff
	var willStun bool = false
	var canDestroyWall bool = false

	spell := Spell{origin, symbol, spellName, isSelfTargeted, isProjectile, isAreaOfEffect,
		isBuff, regainHealth, bonusHealthRegen, manaCost, manaCostPerSecond, bonusDamageMultiplier,
		radius, duration, damage, bonusDamage, bonusDefence, bonusEvasion, bonusCritChance, willStun,
		canDestroyWall, willApplyToEnemies, cooldownTime, coolDownTimeLeft, isOnCoolDown, buffID}

	id++
	return &spell
}
