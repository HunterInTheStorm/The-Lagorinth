package Spells

var id int = 0

type Spell struct {
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

func (spell Spell) CreateBuff(hero *Character) Buff {
	buffName string
	buffID int
	bonusHealthRegen float32
	bonusDamageMultiplier float32
	duration int
	bonusDamage, bonusDefence, bonusEvasion, bonusCritChance int
}

func (spell Spell) CreateProjectile(hero *Character) Projectile {
	loaction Point
	vector Point
	willStun, canDestroyWall bool
	critChance int
	damage int
	buff *Buff
}

func (spell Spell) CreateAreaOfEffect(hero *Character) Effect {
	center Point
	radius int
	duration int
	damage int
	buff *Buff
}

func (spell Spell) InstantSpell(hero *Character) {
	
}



type Projectile struct {
	loaction Point
	vector Point
	willStun, canDestroyWall bool
	critChance int
	damage int
	buff *Buff
}

func (spell Projectile) DoDamage(enemy *NPC) float32 {

}

func (spell *Projectile) MoveForward() {

}

func (spell Projectile) BuffTarget(enemy *NPC) {

}



type Effect struct {
	center Point
	radius int
	duration int
	damage int
	buff *Buff
}

func (spell Spell) EffectCharacter(monsters []NPC) {

}

func (spell Projectile) BuffTarget(enemy *NPC) {

}



type Buff struct {
	buffName string
	buffID int
	bonusHealthRegen float32
	bonusDamageMultiplier float32
	duration int
	bonusDamage, bonusDefence, bonusEvasion, bonusCritChance int
}

func (buff Buff) ApplyBuff(hero *Character) {

}

func (buff Buff) RemoveBuff(hero *Character) {

}

func (buff Buff) LowerDuration(hero *Character) {

}