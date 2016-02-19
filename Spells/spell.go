package Spell

import "github.com/golang/The-Lagorinth/Point"
import "github.com/golang/The-Lagorinth/Labyrinth"
import "math/rand"

var id int = 0

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
	var damage int = 50	
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

type Spell struct {
	Origin *Point.Point
	Symbol string
	SpellName string
	IsSelfTargeted, IsProjectile, IsAreaOfEffect, IsBuff bool
	RegainHealth, BonusHealthRegen float32
	ManaCost, ManaCostPerTurn float32
	BonusDamageMultiplier float32
	Radius, Duration int
	Damage int
	BonusDamage, BonusDefence, BonusEvasion, BonusCritChance int
	WillStun, CanDestroyWall, WillApplyToEnemies bool
	CoolDownTime int
	CoolDownTimeLeft int
	IsOnCoolDown bool
	BuffId int
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

func (spell *Spell) CreateBuff() *Buff {
	var buffName string = spell.SpellName
	var buffID int = spell.BuffId
	var bonusHealthRegen float32 = spell.BonusHealthRegen
	var bonusDamageMultiplier float32 = spell.BonusDamageMultiplier
	var duration int = spell.Duration
	var bonusDamage int = spell.BonusDamage
	var bonusDefence int = spell.BonusDefence
	var bonusEvasion int = spell.BonusEvasion
	var bonusCritChance int = spell.BonusCritChance
	var willApplyToEnemies bool = spell.WillApplyToEnemies
	var manaCostPerTurn float32 = spell.ManaCostPerTurn

	buff := Buff{buffName, buffID, bonusHealthRegen, bonusDamageMultiplier, duration,
		bonusDamage, bonusDefence, bonusEvasion, bonusCritChance, willApplyToEnemies, manaCostPerTurn}

	return &buff
}

func (spell *Spell) CreateProjectile(vector *Point.Point, critical int) *Projectile {

	var symbol string = spell.Symbol
	var spellName string = spell.SpellName
	var location = Point.Point{spell.Origin.X, spell.Origin.Y, nil}
	var newVector = Point.Point{vector.X, vector.Y, nil}
	var willStun bool = spell.WillStun 
	var canDestroyWall bool = spell.CanDestroyWall
	var critChance int = critical
	var damage int = spell.Damage
	var buff *Buff
	if spell.IsBuff {
		buff = spell.CreateBuff()
	} else {
		buff = &Buff{}
	}
	projectile := Projectile{symbol, spellName, location, newVector, willStun, canDestroyWall,
		critChance, damage, buff}
	return &projectile
}

// func (spell Spell) CreateAreaOfEffect(hero *Character) Effect {
// 	center Point
// 	radius int
// 	duration int
// 	damage int
// 	buff *Buff
// }

type Projectile struct {
	Symbol string
	SpellName string
	Location Point.Point
	Vector Point.Point
	WillStun, CanDestroyWall bool
	CritChance int
	Damage int
	Buff *Buff
}

func (spell Projectile) DoDamage() float32 {
	if rand.Intn(100) < spell.CritChance {
		return float32(2 * spell.Damage)
	}
	return float32(spell.Damage)
}

func (spell *Projectile) Move() {
	spell.Location.X += spell.Vector.X
	spell.Location.Y += spell.Vector.Y
}

func (spell *Projectile) ProjectileImapact(labyrinth *Labyrinth.Labyrinth) {
	if spell.CanDestroyWall {
		labyrinth.Labyrinth[spell.Location.X][spell.Location.Y] = Labyrinth.Pass
	}
}

type Effect struct {
	center Point.Point
	radius int
	duration int
	damage int
	buff *Buff
}

type Buff struct {
	BuffName string
	BuffID int
	BonusHealthRegen float32
	BonusDamageMultiplier float32
	Duration int
	BonusDamage, BonusDefence, BonusEvasion, BonusCritChance int
	WillApplyToEnemies bool
	ManaCostPerTurn float32
}

func (buff *Buff) LowerDuration() {
	buff.Duration--
}