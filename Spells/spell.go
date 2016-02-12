package Spells

import "github.com/golang/The-Lagorinth/Point"

var id int = 0

type Spell struct {
	point Point.Point
	symbol string
	spellName string
	selfTargeted, projectile, areaOfEffect, buff bool
	regainHealth, bonusHealthRegen float32
	manaCost, manaCostPerSecond float32
	bonusDamageMultiplier float32
	radius, duration int
	damage int
	bonusDamage, bonusDefence, bonusEvasion, bonusCritChance int
	willStun, canDestroyWall bool
	cooldownTime int
}






// func (spell Spell) CreateBuff(hero *Character) Buff {
// 	buffName string
// 	buffID int
// 	bonusHealthRegen float32
// 	bonusDamageMultiplier float32
// 	duration int
// 	bonusDamage, bonusDefence, bonusEvasion, bonusCritChance int
// }

func (spell Spell) CreateProjectile() Projectile {
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