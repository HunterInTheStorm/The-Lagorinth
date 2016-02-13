package Character

import "github.com/golang/The-Lagorinth/Items"
import "github.com/golang/The-Lagorinth/Spells"
import "github.com/golang/The-Lagorinth/Point"
import "math/rand"

type NPC struct {
	Loaction *Point.Point
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
	Buffs map[int]*Spells.Buff
}

//a npc will move one forward depemding on its orientation
func (npc *NPC) Move() {

}

//function that changes the orientation of a npc
func (npc *NPC) ChangeOrientation(x1 int, x2 int, y1 int, y2 int) {
	npc.Orientation.X = x1 - x2
	npc.Orientation.Y = y1- y2
}

//function that updates the coordinates of an character so that he move one up
func (npc *NPC) MoveNorth() {
	npc.Loaction.X = npc.Loaction.X + 1
	npc.ChangeOrientation(npc.Loaction.X, npc.Loaction.X - 1, npc.Loaction.Y, npc.Loaction.Y)
}

//function that updates the coordinates of an character so that he move one down
func (npc *NPC) MoveSouth() {
	npc.Loaction.X = npc.Loaction.X - 1
	npc.ChangeOrientation(npc.Loaction.X, npc.Loaction.X + 1, npc.Loaction.Y, npc.Loaction.Y)
}

//function that updates the coordinates of an character so that he move one left
func (npc *NPC) MoveWest() {
	npc.Loaction.Y = npc.Loaction.Y + 1
	npc.ChangeOrientation(npc.Loaction.X, npc.Loaction.X, npc.Loaction.Y, npc.Loaction.Y - 1)
}

//function that updates the coordinates of an character so that he move one right
func (npc *NPC) MoveEast() {
	npc.Loaction.Y = npc.Loaction.Y - 1
	npc.ChangeOrientation(npc.Loaction.X, npc.Loaction.X, npc.Loaction.Y, npc.Loaction.Y + 1)
}

//add all of the values of a weapon's properties to the character's
func (npc *NPC) EquipWeapon(newWeapon *Items.Weapon) {
	npc.Weapon = newWeapon
}

//remove all of the values of a weapon's properties from the character's
func (npc *NPC) UnequipWeapon() {
	npc.Weapon = nil
}

//add all of the values of a armor's properties to the character's
func (npc *NPC) EquipArmor(newArmor *Items.Armor) {
	npc.UnequipArmor()
	npc.Armor = newArmor
	npc.CurrentHealth = npc.CurrentHealth + newArmor.Health
	npc.MaxHealth = npc.MaxHealth + newArmor.Health
	npc.CurrentMana = npc.CurrentMana + newArmor.Mana
	npc.MaxMana = npc.MaxMana + newArmor.Mana
}

//remove all of the values of a armor's properties from the character's
func (npc *NPC) UnequipArmor() {
	if npc.Armor != nil {
		npc.CurrentHealth = npc.CurrentHealth - npc.Armor.Health
		npc.MaxHealth = npc.MaxHealth - npc.Armor.Health
		npc.CurrentMana = npc.CurrentMana - npc.Armor.Mana
		npc.MaxMana = npc.MaxMana - npc.Armor.Mana
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

//the function substracs the funcyion's argument "damage" from the characters currentHealth
func (npc *NPC) TakeDamage(damage float32) {
	var damageTaken float32
	if damage - float32(npc.Defence + npc.Armor.Defence) > 0 {
		npc.CurrentHealth = npc.CurrentHealth - damageTaken
	}
}

//Function will update character's currentHealth to a higher value
func (npc *NPC) RegenHealth() {
	npc.CurrentHealth = npc.CurrentHealth + npc.HealthRegen + npc.Armor.HealthRegen
	if npc.CurrentHealth > npc.MaxHealth {
		npc.CurrentHealth = npc.MaxHealth
	}
}

//Function will update character's currentMana to a higher value
func (npc *NPC) RegenMana() {
	npc.CurrentMana = npc.CurrentMana + npc.ManaRegen + npc.Armor.ManaRegen
	if npc.CurrentMana > npc.MaxMana {
		npc.CurrentMana = npc.MaxMana
	}
}

type Hero struct {
	Base *NPC
	ClassName string
	BackGround string
	SpellLList []*Spells.Spell
	Memory map[Point.Point]int
	MemoryDuration int
}

// //will replace the bonuses given from one weapon for the bonuses of another
// func (hero *Hero) SwapWeapon() {

// }

// //will replace the bonuses given from one armor for the bonuses of another
// func (hero *Hero) SwapArmor() {

// }

// //a spell from the list of targetble spells will be envoked
// func (hero *Hero) UseProjectileSpell() Projectile {

// }

// //a spell from the list of selfcast spells will be envoked
// func (hero *Hero) UseSelfTargetSpell() Buff {

// }

// func (hero *Hero) UseAreaOfEffectSpell() Effect {

// }

// func (hero *Hero) UseInstantSpell() {

// }

// //given an array of Points the character memory of the labyrinth will be updated with new tiles
// //that he remebers(tiles that will be displyed)
// func(hero *Hero) MemorizeLabyrinth(points []Point) {

// }

// //values in the memory array will be updated by lowering the duration integer by one 
// func(hero *Hero) UpdateMemory(){

// }


var	TrapTypes []string = []string{"DamageTrap","SpawnTrap","TeleportTrap","MemoryWhipeTrap","TabulaRasaTrap"}

type Trap struct {
	Location *Point.Point
	TrapType string
	DetectDifficulty int
	DisarmDifficulty int
	IsDisarmed bool
	DisarmAtempted bool
	MinDmg int
	MaxDmg int
}

// func (trap *Trap) Randomize() {

// }

// //the function returns the damage(integer) the character will do to an enemy
// func (trap Trap) DoDamage() int {
// 	return 99
// }

// //the function will return a point for where an enemy will spawn
// func (trap Trap) SpawnMonsters() Point {
// 	return Point{0,0}
// }

// func (trap Trap) TeleportPlayer(hero *Hero ) Point, Point {

// }

// func (trap Trap) WhipeMemory(hero *Hero) Point {

// }

// func (trap Trap) WhipeMemoryAndTeleport(hero *Hero) Point, Point {

// }

// type Mimic struct {
// 	currentHealth, maxHealth int
// 	x,y int
// 	orientation *Point
// 	symbol string
// }

// //the function returns the damage(integer) the character will do to an enemy
// func (mimic Mimic) DoDamage() int {
// 	return 0
// }

// //the function substracs the funcyion's argument "damage" from the characters currentHealth
// func (mimic Mimic) TakeDamage(damage int) {

// }

// //function will determine whether the character is hit or not
// func (mimic Mimic) WillBeHit() bool {
// 	return false
// }

// //function determies whether the character will hit or not
// func (mimic Mimic) WillHit() bool {
// 	return true
// }
