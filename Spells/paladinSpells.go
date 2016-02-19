//package Spell handles the creation and functionalities of the spells in the game.
package Spell

import "github.com/golang/The-Lagorinth/Point"
import "github.com/golang/The-Lagorinth/Labyrinth"

//PaladinSpellHeal creates and return a spell.
//It takes a Point structure as an argument, to be used for coordinates of future created projectiles.
//The spell is created for the use of the Paladin class in the game.
//Function makes it easier to modify and balance spells on the fly.
func PaladinSpellHeal(point *Point.Point) *Spell {
	//All spells have this
	origin := point
	var symbol string = ""
	var spellName string = "Great Heal"
	var isSelfTargeted bool = true
	var isProjectile bool = false
	var isAreaOfEffect bool = false
	var isBuff bool = false
	var cooldownTime int = 7
	var isOnCoolDown bool = false
	var coolDownTimeLeft int = 0
	//instant spells
	var regainHealth float32 = 42.0
	var manaCost float32 = 20.0
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

//PaladinSpellHolyArmor creates and return a spell.
//It takes a Point structure as an argument, to be used for coordinates of future created projectiles.
//The spell is created for the use of the Paladin class in the game.
//Function makes it easier to modify and balance spells on the fly.
func PaladinSpellHolyArmor(point *Point.Point) *Spell {
	//All spells have this
	origin := point
	var symbol string = ""
	var spellName string = "Holy Armor"
	var isSelfTargeted bool = true
	var isProjectile bool = false
	var isAreaOfEffect bool = false
	var isBuff bool = true
	var cooldownTime int = 20
	var isOnCoolDown bool = false
	var coolDownTimeLeft int = 0
	//instant spell
	var regainHealth float32 = 0.0
	var manaCost float32 = 10.0 
	//buff
	var buffId int = id
	var manaCostPerSecond float32= 3.0
	var bonusHealthRegen float32 = 1.0
	var bonusDamageMultiplier float32= 0.0 
	var duration int = 15
	var bonusDamage int = 0
	var bonusDefence int = 4
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

//PaladinSpellHolyBolt creates and return a spell.
//It takes a Point structure as an argument, to be used for coordinates of future created projectiles.
//The spell is created for the use of the Paladin class in the game.
//Function makes it easier to modify and balance spells on the fly.
func PaladinSpellHolyBolt(point *Point.Point) *Spell {
	//All spells have this
	origin := point
	var symbol string = Labyrinth.Projectile
	var spellName string = "Holy Bolt"
	var isSelfTargeted bool = false
	var isProjectile bool = true
	var isAreaOfEffect bool = false
	var isBuff bool = false
	var cooldownTime int = 3
	var isOnCoolDown bool = false
	var coolDownTimeLeft int = 0
	//instant spell
	var regainHealth float32 = 0.0
	var manaCost float32 = 25.0 
	//buff
	var buffID int = id
	var manaCostPerSecond float32= 0.0
	var bonusHealthRegen float32 = 0.0
	var bonusDamageMultiplier float32= 0.0
	var duration int = 0
	var bonusDamage int = 0
	var bonusDefence int = 0
	var bonusEvasion int = 0
	var bonusCritChance int = 0
	var willApplyToEnemies bool = true
	//AoE has buff
	var damage int = 400	
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

