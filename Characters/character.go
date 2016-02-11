package Characters

type NPC struct {
	loaction *Point
	symbol string
	name string
	orientation *Point
	weapon *Weapon
	armor *armor
	dmgMultuplier float32
	defence, evasion, critChance int
	currentHealth, maxHealth, healthRegen float32
	currentMana, maxMana, manaRegen float32
	visionRadious int
	isStunned bool
	flyingVision bool
	buffs map[int]*Buff
}

//a npc will move one forward depemding on its orientation
func (npc *NPC) Move() {

}

//function that changes the orientation of a npc
func (npc *NPC) ChangeOrientation(x1 int, x2 int, y1 int, y2 int) {
	npc.orientation.x = x1 - x2
	npc.orientation.y = y1- y2
}

//function that updates the coordinates of an character so that he move one up
func (npc *NPC) MoveNorth() {
	npc.loaction.x = npc.loaction.x + 1
	npc.ChangeOrientation(npc.loaction.x, npc.loaction.x - 1, npc.loaction.y, npc.loaction.y)
}

//function that updates the coordinates of an character so that he move one down
func (npc *NPC) MoveSouth() {
	npc.loaction.x = npc.loaction.x - 1
	npc.ChangeOrientation(npc.loaction.x, npc.loaction.x + 1, npc.loaction.y, npc.loaction.y)
}

//function that updates the coordinates of an character so that he move one left
func (npc *NPC) MoveWest() {
	npc.loaction.y = npc.loaction.y + 1
	npc.ChangeOrientation(npc.loaction.x, npc.loaction.x, npc.loaction.y, npc.loaction.y - 1)
}

//function that updates the coordinates of an character so that he move one right
func (npc *NPC) MoveEast() {
	npc.loaction.y = npc.loaction.y - 1
	npc.ChangeOrientation(npc.loaction.x, npc.loaction.x, npc.loaction.y, npc.loaction.y + 1)
}

//add all of the values of a weapon's properties to the character's
func (npc *NPC) EquipWeapon(newWeapon *Weapon) {
	npc.weapon = newWeapon
}

//remove all of the values of a weapon's properties from the character's
func (npc *NPC) UnequipWeapon() {
	npc.weapon = nil
}

//add all of the values of a armor's properties to the character's
func (npc *NPC) EquipArmor(newArmor *Armor) {
	npc.UnequipArmor()
	npc.armor = newArmor
	npc.currentHealth = npc.currentHealth + armor.health
	npc.maxHealth = npc.maxHealth + armor.health
	npc.currentMana = npc.currentMana + npc.armor.mana
	npc.maxMana = npc.maxMana + npc.armor.mana
}

//remove all of the values of a armor's properties from the character's
func (npc *NPC) UnequipArmor() {
	if npc.armor != nil {
		npc.currentHealth = npc.currentHealth - npc.armor.health
		npc.maxHealth = npc.maxHealth - npc.armor.health
		npc.currentMana = npc.currentMana - npc.armor.mana
		npc.maxMana = npc.maxMana - npc.armor.mana
		npc.armor = nil
	}
}

//the function returns the damage(integer) the character will do to an enemy
func (npc *NPC) DoDamage() float32 {
	if rand.Intn(100) < npc.critChance + npc.weapon.bonusCritChance {
		return 2 * npc.dmgMultuplier * npc.weapon.Damage()
	}
	return npc.dmgMultuplier * npc.weapon.Damage()
}

//the function substracs the funcyion's argument "damage" from the characters currentHealth
func (npc *NPC) TakeDamage(damage float32) {
	var damageTaken float32
	if damage - npc.defence + npc.armor.defence > 0 {
		npc.currentHealth = npc.currentHealth - damageTaken
	}
}

//Function will update character's currentHealth to a higher value
func (npc *NPC) RegenHealth() {
	npc.currentHealth = npc.currentHealth + npc.healthRegen + npc.armor.healthRegen
	if npc.currentHealth > npc.maxHealth {
		npc.currentHealth = npc.maxHealth
	}
}

//Function will update character's currentMana to a higher value
func (npc *NPC) RegenMana() {
	npc.currentMana = npc.currentMana + npc.manaRegen + npc.armor.manaRegen
	if npc.currentMana > npc.maxMana {
		npc.currentMana = npc.maxMana
	}
}

type Character struct {
	base *NPC
	className string
	spellLList []Spell
	memory map[Point]int
	memoryDuration int
}

//will replace the bonuses given from one weapon for the bonuses of another
func (hero *Character) SwapWeapon() {

}

//will replace the bonuses given from one armor for the bonuses of another
func (hero *Character) SwapArmor() {

}

//a spell from the list of targetble spells will be envoked
func (hero *Character) UseProjectileSpell() Projectile {

}

//a spell from the list of selfcast spells will be envoked
func (hero *Character) UseSelfTargetSpell() Buff {

}

func (hero *Character) UseAreaOfEffectSpell() Effect {

}

func (hero *Character) UseInstantSpell() {

}

//given an array of Points the character memory of the labyrinth will be updated with new tiles
//that he remebers(tiles that will be displyed)
func(hero *Character) MemorizeLabyrinth(points []Point) {

}

//values in the memory array will be updated by lowering the duration integer by one 
func(hero *Character) UpdateMemory(){

}

type Trap struct {
	x, y int
	symbol string
	disarmDifficulty int
	isDisarmed bool
	disarmAtempted bool
	minDmg int
	maxDmg int
}

//the function returns the damage(integer) the character will do to an enemy
func (trap Trap) DoDamage() int {
	return 99
}

//the function will return a point for where an enemy will spawn
func (trap Trap) SpawnMonsters() Point {
	return Point{0,0}
}

func (trap Trap) TeleportPlayer(hero *Character ) Point, Point {

}

func (trap Trap) WhipeMemory(hero *Character) Point {

}

func (trap Trap) WhipeMemoryAndTeleport(hero *Character) Point, Point {

}

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
