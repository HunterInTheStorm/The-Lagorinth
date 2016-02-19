//package Spell handles the creation and functionalities of the spells in the game.
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

//LowerDuration lower the duration left on a buff
func (buff *Buff) LowerDuration() {
	buff.Duration--
}
