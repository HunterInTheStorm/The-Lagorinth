//package Spell handles the creation and functionalities of the spells in the game.
package Spell

import "github.com/golang/The-Lagorinth/Point"

//id is used in the creation of spells and more accurately Buffs.
//It assures the same Buff cannot be applied to a character again.
//Instead only refreshing the duration.
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

//GoOnCoolDown sets a cool down for the used spell so that it cannot be used again until it reaches 0. 
func (spell *Spell) GoOnCoolDown() {
	spell.IsOnCoolDown = true
	spell.CoolDownTimeLeft = spell.CoolDownTime
}

//LowerCoolDownTime lowers the cool down a spell by 1 for every call.
//When it reaches 0 the IsOnCoolDown flag is set to false
func (spell *Spell) LowerCoolDownTime() {
	if spell.CoolDownTimeLeft > 0 {
		spell.CoolDownTimeLeft--
		if spell.CoolDownTimeLeft == 0 {
			spell.IsOnCoolDown = false
		}
	}
}

//CreateBuff creates and returns a new Buff structure.
//Values from the Spell structure are used to create the values for the Buff structure
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

//CreateProjectile creates and returns a new Projectile structure.
//Values from the Spell structure are used to create the values for the Projectile structure
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
