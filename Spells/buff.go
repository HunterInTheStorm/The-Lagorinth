package Spell

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
