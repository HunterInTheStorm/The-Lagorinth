//package Spell handles the creation and functionalities of the spells in the game.
package Spell

import "github.com/golang/The-Lagorinth/Point"
import "github.com/golang/The-Lagorinth/Labyrinth"

//MageSpellSacrifice creates and return a spell.
//It takes a Point structure as an argument, to be used for coordinates of future created projectiles.
//The spell is created for the use of the Mage class in the game.
//Function makes it easier to modify and balance spells on the fly.
func MageSpellSacrifice(point *Point.Point) *Spell {
	//All spells have this
	origin := point
	var symbol string = ""
	var spellName string = "Sacrifice"
	var isSelfTargeted bool = true
	var isProjectile bool = false
	var isAreaOfEffect bool = false
	var isBuff bool = false
	var cooldownTime int = 20
	var isOnCoolDown bool = false
	var coolDownTimeLeft int = 0
	//instant spells
	var regainHealth float32 = -50.0
	var manaCost float32 = -200.0
	//buff
	var buffId int = id
	var manaCostPerSecond float32 = 0.0
	var bonusHealthRegen float32 = 0.0
	var bonusDamageMultiplier float32 = 0.0 
	var duration int = 0
	var bonusDamage int = 0
	var bonusDefence int = 0
	var bonusEvasion int = 0
	var bonusCritChance int = 0
	var willApplyToEnemies bool = false
	//AoE has buff
	var damage int = 0
	var radius int = 0
	//Projectile has damage and buff
	var willStun bool = false
	var canDestroyWall bool = false

	spell := Spell{origin, symbol, spellName, isSelfTargeted, isProjectile, isAreaOfEffect,
		isBuff, regainHealth, bonusHealthRegen, manaCost, manaCostPerSecond,bonusDamageMultiplier,
		radius, duration, damage, bonusDamage, bonusDefence,bonusEvasion, bonusCritChance, willStun,
		canDestroyWall, willApplyToEnemies, cooldownTime, coolDownTimeLeft, isOnCoolDown, buffId}

	id++
	return &spell
}

//MageSpellBallLightning creates and return a spell.
//It takes a Point structure as an argument, to be used for coordinates of future created projectiles.
//The spell is created for the use of the Mage class in the game.
//Function makes it easier to modify and balance spells on the fly.
func MageSpellBallLightning(point *Point.Point) *Spell {
	//All spells have this
	origin := point
	var symbol string = Labyrinth.Projectile
	var spellName string = "Ball Lightning"
	var isSelfTargeted bool = false
	var isProjectile bool = true
	var isAreaOfEffect bool = false
	var isBuff bool = false
	var cooldownTime int = 10
	var isOnCoolDown bool = false
	var coolDownTimeLeft int = 0
	//instant spell
	var regainHealth float32 = 0.0
	var manaCost float32 = 30.0 
	//buff
	var buffId int = id
	var manaCostPerSecond float32= 0.0
	var bonusHealthRegen float32 = 0.0
	var bonusDamageMultiplier float32= 0.0 
	var duration int = 0
	var bonusDamage int = 0
	var bonusDefence int = 0
	var bonusEvasion int = 0
	var bonusCritChance int = 0
	var willApplyToEnemies bool = false
	//AoE has buff
	var damage int = 100
	var radius int = 0 
	//Projectile has damage and buff
	var willStun bool = true
	var canDestroyWall bool = false

	spell := Spell{origin, symbol, spellName, isSelfTargeted, isProjectile, isAreaOfEffect,
		isBuff, regainHealth, bonusHealthRegen, manaCost, manaCostPerSecond,bonusDamageMultiplier,
		radius, duration, damage, bonusDamage, bonusDefence,bonusEvasion, bonusCritChance, willStun,
		canDestroyWall, willApplyToEnemies, cooldownTime, coolDownTimeLeft, isOnCoolDown, buffId}

	id++
	return &spell
}

//MageSpellFireBall creates and return a spell.
//It takes a Point structure as an argument, to be used for coordinates of future created projectiles.
//The spell is created for the use of the Mage class in the game.
//Function makes it easier to modify and balance spells on the fly.
func MageSpellFireBall(point *Point.Point) *Spell {
	//All spells have this
	origin := point
	var symbol string = Labyrinth.Projectile
	var spellName string = "Fire Ball"
	var isSelfTargeted bool = false
	var isProjectile bool = true
	var isAreaOfEffect bool = false
	var isBuff bool = true
	var cooldownTime int = 2
	var isOnCoolDown bool = false
	var coolDownTimeLeft int = 0
	//instant spell
	var regainHealth float32 = 0.0
	var manaCost float32 = 15.0 
	//buff
	var buffID int = id
	var manaCostPerSecond float32= 0.0
	var bonusHealthRegen float32 = -30.0
	var bonusDamageMultiplier float32= 0.0
	var duration int = 4
	var bonusDamage int = 0
	var bonusDefence int = -20
	var bonusEvasion int = 0
	var bonusCritChance int = 0
	var willApplyToEnemies bool = true
	//AoE has buff
	var damage int = 35.0	
	var radius int = 0 
	//Projectile has damage and buff
	var willStun bool = false
	var canDestroyWall bool = false

	spell := Spell{origin, symbol, spellName, isSelfTargeted, isProjectile, isAreaOfEffect,
		isBuff, regainHealth, bonusHealthRegen, manaCost, manaCostPerSecond,bonusDamageMultiplier,
		radius, duration, damage, bonusDamage, bonusDefence,bonusEvasion, bonusCritChance, willStun,
		canDestroyWall, willApplyToEnemies, cooldownTime, coolDownTimeLeft, isOnCoolDown, buffID}

	id++
	return &spell
}
