package Character

import "github.com/golang/The-Lagorinth/Items"
import "github.com/golang/The-Lagorinth/Spells"
import "github.com/golang/The-Lagorinth/Point"
import "github.com/golang/The-Lagorinth/Labyrinth"
import "math/rand"

type NPC struct {
	Location *Point.Point
	Symbol string
	Name string
	Orientation *Point.Point
	Weapon *Items.Weapon
	Armor *Items.Armor
	DmgMultuplier float32
	Defence, Evasion, CritChance int
	CurrentHealth, MaxHealth, HealthRegen float32
	CurrentMana, MaxMana, ManaRegen float32
	VisionRadious int
	IsStunned bool
	BuffList map[int]*Spell.Buff
	IsHuman bool
	TrapHandling int
}

func (npc *NPC) moveTowardsHero(labyrinth *Labyrinth.Labyrinth) (bool, *Point.Point) {
	var upTile string
	if labyrinth.IsInBondaries(npc.Location.X - 1, npc.Location.Y) {
		upTile = labyrinth.Labyrinth[npc.Location.X - 1][npc.Location.Y]
	}
	if upTile == Labyrinth.CharSymbol {
		return true, &Point.Point{npc.Location.X - 1, npc.Location.Y, nil}
	}

	var downTile string
	if labyrinth.IsInBondaries(npc.Location.X + 1, npc.Location.Y) {
		downTile =	labyrinth.Labyrinth[npc.Location.X + 1][npc.Location.Y]
	}
	if downTile == Labyrinth.CharSymbol {
		return true, &Point.Point{npc.Location.X + 1, npc.Location.Y, nil}
	}

	var leftTile string
	if labyrinth.IsInBondaries(npc.Location.X, npc.Location.Y - 1) {
		leftTile =	labyrinth.Labyrinth[npc.Location.X][npc.Location.Y - 1]
	}
	if leftTile == Labyrinth.CharSymbol {
		return true, &Point.Point{npc.Location.X, npc.Location.Y - 1, nil}
	}

	var rightTile string
	if labyrinth.IsInBondaries(npc.Location.X, npc.Location.Y + 1) {
		rightTile =	labyrinth.Labyrinth[npc.Location.X][npc.Location.Y + 1]
	}
	if rightTile == Labyrinth.CharSymbol {
		return true, &Point.Point{npc.Location.X, npc.Location.Y + 1, nil}
	}
	return false, &Point.Point{}
}

func (npc *NPC) makeDecisionWhereToMove(labyrinth *Labyrinth.Labyrinth) (bool, *Point.Point) {
	frontTile := labyrinth.Labyrinth[npc.Location.X + npc.Orientation.X][npc.Location.Y + npc.Orientation.Y]
	if frontTile != Labyrinth.Wall && frontTile != Labyrinth.Monster && frontTile != Labyrinth.Treasure {
		if rand.Intn(100) < 80 {
			return true, &Point.Point{npc.Location.X + npc.Orientation.X, npc.Location.Y + npc.Orientation.Y, nil}
		}
	} else {
		direction := make([]Point.Point, 0, 4)
		upTile := labyrinth.Labyrinth[npc.Location.X - 1][npc.Location.Y]
		if upTile != Labyrinth.Wall && upTile != Labyrinth.Monster && upTile != Labyrinth.Treasure {
			direction = append(direction, Point.Point{npc.Location.X - 1, npc.Location.Y, nil})
		}
		downTile := labyrinth.Labyrinth[npc.Location.X + 1][npc.Location.Y]
		if downTile != Labyrinth.Wall && downTile != Labyrinth.Monster && downTile != Labyrinth.Treasure {
			direction = append(direction, Point.Point{npc.Location.X + 1, npc.Location.Y, nil})
		}
		leftTile := labyrinth.Labyrinth[npc.Location.X][npc.Location.Y - 1]
		if leftTile != Labyrinth.Wall && leftTile != Labyrinth.Monster && leftTile != Labyrinth.Treasure {
			direction = append(direction, Point.Point{npc.Location.X, npc.Location.Y - 1, nil})
		}
		rightTile := labyrinth.Labyrinth[npc.Location.X][npc.Location.Y + 1]
		if rightTile != Labyrinth.Wall && rightTile != Labyrinth.Monster && rightTile != Labyrinth.Treasure {
			direction = append(direction, Point.Point{npc.Location.X, npc.Location.Y + 1, nil})
		}
		if len(direction) != 0 {
			return true, &direction[rand.Intn(len(direction))]
		}
	}
	return false, &Point.Point{-1,-1,nil}
}

//a npc will move one forward depemding on its orientation
func (npc *NPC) Move(labyritnh *Labyrinth.Labyrinth) *Point.Point {
	isNextToHero, location := npc.moveTowardsHero(labyritnh)
	if isNextToHero {
		return location
	}
	isDecisionMade, place := npc.makeDecisionWhereToMove(labyritnh)
	if isDecisionMade {
		return place
	}
	return npc.Location
}

//function that changes the orientation of a npc
func (npc *NPC) ChangeOrientation(x2 int, y2 int) {
	npc.Orientation.X = x2 - npc.Location.X
	npc.Orientation.Y = y2 - npc.Location.Y
}

//add all of the values of a weapon's properties to the character's
func (npc *NPC) EquipWeapon(newWeapon *Items.Weapon) {
	npc.Weapon = newWeapon
}

//remove all of the values of a weapon's properties from the character's
func (npc *NPC) UnequipWeapon() {
	if npc.Weapon != nil {
		npc.Weapon = nil
	}
}

//add all of the values of a armor's properties to the character's
func (npc *NPC) EquipArmor(newArmor *Items.Armor) {
	npc.Armor = newArmor
	npc.CurrentHealth += newArmor.Health
	npc.MaxHealth += newArmor.Health
	npc.CurrentMana += newArmor.Mana
	npc.MaxMana += newArmor.Mana
	npc.Evasion += newArmor.Evasion
	npc.HealthRegen += newArmor.HealthRegen
	npc.ManaRegen += newArmor.ManaRegen
}

//remove all of the values of a armor's properties from the character's
func (npc *NPC) UnequipArmor() {
	if npc.Armor != nil {
		npc.CurrentHealth -= npc.Armor.Health
		npc.MaxHealth -= npc.Armor.Health
		npc.CurrentMana -= npc.Armor.Mana
		npc.MaxMana -= npc.Armor.Mana
		npc.Evasion -= npc.Armor.Evasion
		npc.HealthRegen -= npc.Armor.HealthRegen
		npc.ManaRegen -= npc.Armor.ManaRegen
		npc.Armor = nil
	}
}

//the function returns the damage(integer) the character will do to an enemy
func (npc *NPC) DoDamage() float32 {
	if rand.Intn(100) < npc.CritChance + npc.Weapon.BonusCritChance {
		return 2 * npc.DmgMultuplier * npc.Weapon.Damage()
	}
	return npc.DmgMultuplier * npc.Weapon.Damage()
}

func (npc *NPC) CombinedDefence() float32 {
	return float32(npc.Defence + npc.Armor.Defence)
}

//the function substracs the funcyion's argument "damage" from the characters currentHealth
func (npc *NPC) TakeDamage(damage float32) {
	var damageTaken float32 = damage - npc.CombinedDefence()
	if damageTaken > 0 {
		npc.CurrentHealth = npc.CurrentHealth - damageTaken
	}
}

//Function will update character's currentHealth to a higher value
func (npc *NPC) RegenHealth() {
	npc.CurrentHealth = npc.CurrentHealth + npc.HealthRegen
	if npc.CurrentHealth > npc.MaxHealth {
		npc.CurrentHealth = npc.MaxHealth
	}
}

//Function will update character's currentMana to a higher value
func (npc *NPC) RegenMana() {
	npc.CurrentMana = npc.CurrentMana + npc.ManaRegen
	if npc.CurrentMana > npc.MaxMana {
		npc.CurrentMana = npc.MaxMana
	}
}

func (npc *NPC) ApplyBuff(buff *Spell.Buff) {
	npc.HealthRegen += buff.BonusHealthRegen
	npc.DmgMultuplier += buff.BonusDamageMultiplier
	npc.Weapon.BonusDmg += buff.BonusDamage
	npc.Defence += buff.BonusDefence
	npc.Evasion += buff.BonusEvasion
	npc.CritChance += buff.BonusCritChance
	npc.ManaRegen -= buff.ManaCostPerTurn
}

func (npc *NPC) RemoveBuff(buff *Spell.Buff) {
	npc.HealthRegen -= buff.BonusHealthRegen
	npc.DmgMultuplier -= buff.BonusDamageMultiplier
	npc.Weapon.BonusDmg -= buff.BonusDamage
	npc.Defence -= buff.BonusDefence
	npc.Evasion -= buff.BonusEvasion
	npc.CritChance -= buff.BonusCritChance
	npc.ManaRegen += buff.ManaCostPerTurn
}

func (npc *NPC) Regenerate() {
	npc.RegenMana()
	npc.RegenHealth()
}

func (npc *NPC) ProjectileToTheFace(projectile *Spell.Projectile) {
	npc.IsStunned = projectile.WillStun
	if projectile.Buff != nil {
		npc.ApplyBuff(projectile.Buff)
	}
}


