package Spell

import "github.com/golang/The-Lagorinth/Point"

var id int = 0

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

//type Effect struct {
// 	center Point.Point
// 	radius int
// 	duration int
// 	damage int
// 	buff *Buff
// }

// func (spell Spell) CreateAreaOfEffect(hero *Character) Effect {
// 	center Point
// 	radius int
// 	duration int
// 	damage int
// 	buff *Buff
// }
