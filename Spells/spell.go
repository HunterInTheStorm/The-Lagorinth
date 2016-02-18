package Spell

import "github.com/golang/The-Lagorinth/Point"

func PaladinSpellHeal(point *Point.Point) *Spell {
	//All spells have this
	origin := point
	var symbol string = ""
	var spellName string = "Greate Heal"
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
		canDestroyWall, willApplyToEnemies, cooldownTime, coolDownTimeLeft, isOnCoolDown}

	return &spell
}

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
		canDestroyWall, willApplyToEnemies, cooldownTime, coolDownTimeLeft, isOnCoolDown}

	return &spell
}

func PaladinSpellHolyBolt(point *Point.Point) *Spell {
	//All spells have this
	origin := point
	var symbol string = ""
	var spellName string = "Holy Bolt"
	var isSelfTargeted bool = false
	var isProjectile bool = true
	var isAreaOfEffect bool = false
	var isBuff bool = false
	var cooldownTime int = 5
	var isOnCoolDown bool = false
	var coolDownTimeLeft int = 0
	//instant spell
	var regainHealth float32 = 0.0
	var manaCost float32 = 25.0 
	//buff
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
	var damage int = 50	
	var radius int = 0 
	//Projectile has damage and buff
	var willStun bool = true
	var canDestroyWall bool = false

	spell := Spell{origin, symbol, spellName, isSelfTargeted, isProjectile, isAreaOfEffect,
		isBuff, regainHealth, bonusHealthRegen, manaCost, manaCostPerSecond,bonusDamageMultiplier,
		radius, duration, damage, bonusDamage, bonusDefence,bonusEvasion, bonusCritChance, willStun,
		canDestroyWall, willApplyToEnemies, cooldownTime, coolDownTimeLeft, isOnCoolDown}

	return &spell
}


var id int = 0

type Spell struct {
	Origin *Point.Point
	Symbol string
	SpellName string
	IsSelfTargeted, IsProjectile, IsAreaOfEffect, IsBuff bool
	RegainHealth, BonusHealthRegen float32
	ManaCost, ManaCostPerSecond float32
	BonusDamageMultiplier float32
	Radius, Duration int
	Damage int
	BonusDamage, BonusDefence, BonusEvasion, BonusCritChance int
	WillStun, CanDestroyWall, WillApplyToEnemies bool
	CoolDownTime int
	CoolDownTimeLeft int
	IsOnCoolDown bool
}

func (spell *Spell) GoOnCoolDown() {
	spell.IsOnCoolDown = true
	spell.CoolDownTimeLeft = spell.CoolDownTime
}

func (spell *Spell) LowerCoolDownTime() {
	if spell.CoolDownTimeLeft > 0 {
		spell.CoolDownTimeLeft--
		if spell.CoolDownTimeLeft == 0 {
			spell.IsOnCoolDown = false
		}
	}
}


// func (spell Spell) CreateBuff(hero *Character) Buff {
// 	buffName string
// 	buffID int
// 	bonusHealthRegen float32
// 	bonusDamageMultiplier float32
// 	duration int
// 	bonusDamage, bonusDefence, bonusEvasion, bonusCritChance int
// }

func (spell *Spell) CreateProjectile() Projectile {
	// pnt := Point{0,1}
// 	loaction Point
// 	vector Point
// 	willStun, canDestroyWall bool
// 	critChance int
// 	damage int
// 	buff *Buff
	return Projectile{}
}

// func (spell Spell) CreateAreaOfEffect(hero *Character) Effect {
// 	center Point
// 	radius int
// 	duration int
// 	damage int
// 	buff *Buff
// }

// func (spell Spell) InstantSpell(hero *Character) {
	
// }


type Projectile struct {
	loaction Point.Point
	vector Point.Point
	willStun, canDestroyWall bool
	critChance int
	damage int
	buff *Buff
}

// func (spell Projectile) DoDamage() float32 {

// }

// func (spell *Projectile) MoveForward() {

// }

// func (spell Projectile) BuffTarget() {

// }



type Effect struct {
	center Point.Point
	radius int
	duration int
	damage int
	buff *Buff
}

// func (spell Effect) EffectCharacter() {

// }

// func (spell Effect) BuffTarget() {

// }



type Buff struct {
	buffName string
	buffID int
	bonusHealthRegen float32
	bonusDamageMultiplier float32
	duration int
	bonusDamage, bonusDefence, bonusEvasion, bonusCritChance int
}

// func (buff Buff) ApplyBuff() {

// }

// func (buff Buff) RemoveBuff() {

// }

// func (buff Buff) LowerDuration() {

// }