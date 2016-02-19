//Package Character handles the creation and management of characters in the game.
package Character

import "github.com/golang/The-Lagorinth/Point"
import "github.com/golang/The-Lagorinth/Spells"
import "github.com/golang/The-Lagorinth/Labyrinth"
import "github.com/golang/The-Lagorinth/Items"

type Hero struct {
	Base *NPC
	ClassName string
	BackGround string
	SpellList []*Spell.Spell
	Memory map[Point.Point]int
	MemoryDuration int
}

//SwapArmor replaces the currently equipped Weapon with a new one.
func (hero *Hero) SwapWeapon(weapon *Items.Weapon) {
	hero.Base.UnequipWeapon()
	hero.Base.EquipWeapon(weapon)
}

//SwapArmor replaces the currently equipped Armor with a new one.
func (hero *Hero) SwapArmor(armor *Items.Armor) {
	hero.Base.UnequipArmor()
	hero.Base.EquipArmor(armor)
}

//ApplyBackground add the arguments' field values to those of the character.
func (hero *Hero) ApplyBackground(background *BackGround) {
	hero.BackGround = background.Name
	hero.MemoryDuration += background.BonusMemoryDuration
	hero.Base.VisionRadious += background.BonusVisionRadius
	hero.Base.MaxMana += background.BonusMana
	hero.Base.CurrentMana += background.BonusMana
	hero.Base.ManaRegen += background.BonusManaRegen
	hero.Base.MaxHealth += background.BonusHealth
	hero.Base.CurrentHealth += background.BonusHealth
	hero.Base.HealthRegen += background.BonusHealthRegen
	hero.Base.Defence += background.BonusArmor
	hero.Base.Evasion += background.BonusEvasion
	hero.Base.CritChance += background.BonusCritChance
	hero.Base.DmgMultuplier += background.BonusDmgMultuplier
}


//UseInstantSpell handles the use of an instant spell
func (hero *Hero) UseInstantSpell(spell *Spell.Spell) {
	hero.Base.CurrentMana -= spell.ManaCost
	hero.Base.CurrentHealth += spell.RegainHealth

	if hero.Base.CurrentMana > hero.Base.MaxMana {
		hero.Base.CurrentMana = hero.Base.MaxMana
	}

	if hero.Base.CurrentHealth > hero.Base.MaxHealth {
		hero.Base.CurrentHealth = hero.Base.MaxHealth
	}
}

//UseBuffSpell handles the use of a Buff Spell.
//Creates and applies a buff to the character.
func (hero *Hero) UseBuffSpell(spell *Spell.Spell) {
	hero.Base.CurrentMana -= spell.ManaCost
	buff := spell.CreateBuff()
	if _, ok := hero.Base.BuffList[spell.BuffId]; !ok {
		hero.Base.BuffList[spell.BuffId] = buff
		hero.Base.ApplyBuff(buff)
	} else {
		currentBuff := hero.Base.BuffList[spell.BuffId]
		currentBuff.Duration = spell.Duration
	}
}



//UseProjectileSpell returns a projectile created from the Spell argument. 
func (hero *Hero) UseProjectileSpell(spell *Spell.Spell) *Spell.Projectile {
	crit := hero.Base.CritChance + hero.Base.Weapon.BonusCritChance
	hero.Base.CurrentMana -= spell.ManaCost
	return spell.CreateProjectile(hero.Base.Orientation, crit)
}

//UpdateMemory will lower the duration for which a character remembers a tile.
func (hero *Hero) UpdateMemory() {
	for point, _ := range hero.Memory {
		if hero.Memory[point] > -1 {
			hero.Memory[point]--
		} else {
			delete(hero.Memory, point)
		}
	}
}

//MemorizeLabyrinth determines which tiles a character can see and add them to a map.
//Takes a labyrinth as an argument.
func (hero *Hero) MemorizeLabyrinth(labyrinth *Labyrinth.Labyrinth, center *Point.Point) {
	var minX int = center.X - hero.Base.VisionRadious
	var maxX int = center.X + hero.Base.VisionRadious
	var minY int = center.Y - hero.Base.VisionRadious
	var maxY int = center.Y + hero.Base.VisionRadious
	var y int
	for currentY := minY; currentY <= maxY; currentY++ {
		for xAscend := center.X; xAscend <= maxX; xAscend++ {
			y = Point.LineEquationRegardsToX(xAscend, center.X, center.Y, maxX, currentY)
			if labyrinth.IsInBondaries(xAscend, y) && labyrinth.Labyrinth[xAscend][y] != Labyrinth.Wall {
				hero.Memory[Point.Point{xAscend, y, nil}] = hero.MemoryDuration
			} else {
				hero.Memory[Point.Point{xAscend, y, nil}] = hero.MemoryDuration
				break 
			}
		}
		for xDescend := center.X; xDescend >= minX; xDescend-- {
			y = Point.LineEquationRegardsToX(xDescend, center.X, center.Y, minX, currentY)
			if labyrinth.IsInBondaries(xDescend, y) && labyrinth.Labyrinth[xDescend][y] != Labyrinth.Wall {
				hero.Memory[Point.Point{xDescend, y, nil}] = hero.MemoryDuration
			} else {
				hero.Memory[Point.Point{xDescend, y, nil}] = hero.MemoryDuration
				break 
			}
		}
	}
	var x int
	for currentX := minX; currentX <= maxX; currentX++ {
		for yDescend := center.Y; yDescend >= minY; yDescend-- {
			x = Point.LineEquationRegardsToY(yDescend, center.X, center.Y, currentX, minY)
			if labyrinth.IsInBondaries(x, yDescend) && labyrinth.Labyrinth[x][yDescend] != Labyrinth.Wall {
				hero.Memory[Point.Point{x, yDescend, nil}] = hero.MemoryDuration
			} else {
				hero.Memory[Point.Point{x, yDescend, nil}] = hero.MemoryDuration
				break 
			}
		}
		for yAscend := center.Y; yAscend <= maxY; yAscend++ {
			x = Point.LineEquationRegardsToY(yAscend, center.X, center.Y, currentX, maxY)
			if labyrinth.IsInBondaries(x, yAscend) && labyrinth.Labyrinth[x][yAscend] != Labyrinth.Wall {
				hero.Memory[Point.Point{x, yAscend, nil}] = hero.MemoryDuration
			} else {
				hero.Memory[Point.Point{x, yAscend, nil}] = hero.MemoryDuration
				break 
			}
		}
	}
}
