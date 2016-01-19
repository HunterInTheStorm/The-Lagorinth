package main

type NPC struct {
	x, y int
	symbol string
	name string
	orientation *Point
	weapon *Weapon
	armor *armor
	bonusDmg, bonusCritChance int
	dmgMultuplier double
	defence, evasion int
	currentHealth, maxHealth int
	visionRadious int
	isStunned bool
	flyingVision bool
}

//a npc will move one forward depemding on its orientation
func (npc NPC) Move() {

}

//function that changes the orientation of a npc
func (npc NPC) ChangeOrientation() {
	
}

//function that updates the coordinates of an character so that he move one up
func (npc NPC) MoveNorth() {

}

//function that updates the coordinates of an character so that he move one down
func (npc NPC) MoveSouth() {
	
}

//function that updates the coordinates of an character so that he move one left
func (npc NPC) MoveWest() {
	
}

//function that updates the coordinates of an character so that he move one right
func (npc NPC) MoveEast() {
	
}

//add all of the values of a weapon's properties to the character's
func (npc NPC) EquipWeapon(weapon Weapon) {

}

//remove all of the values of a weapon's properties from the character's
func (npc NPC) UnequipWeapon() {

}

//add all of the values of a armor's properties to the character's
func (npc NPC) EquipArmor(armor Armor) {

}

//remove all of the values of a armor's properties from the character's
func (npc NPC) UnequipArmor() {

}

//the function returns the damage(integer) the character will do to an enemy
func (npc NPC) DoDamage() int {
	return 0
}

//the function substracs the funcyion's argument "damage" from the characters currentHealth
func (npc NPC) TakeDamage(damage int) {

}

//function will determine whether the character is hit or not
func (npc NPC) WillBeHit() bool {
	return false
}

//function determies whether the character will hit or not
func (npc NPC) WillHit() bool {
	return true
}


//Function will update character's currentHealth to a higher value
func (npc NPC) RegenHealth() {

}

type Character struct {
	base *NPC
	className string
	currentMana, maxMana int
	spellLList []Spell
	memory []struct(point Point, duration int}
	memoryDuration int

}

//Function will update character's currentMana to a hiegher value
func (hero Character) RegenMana() {

}

//will replace the bonuses given from one weapon for the bonuses of another
func (hero Character) SwapWeapon() {

}

//will replace the bonuses given from one armor for the bonuses of another
func (hero Character) SwapArmor() {

}

//a spell from the list of targetble spells will be envoked
func (hero Character) UseTargetSpell() {

}

//a spell from the list of selfcast spells will be envoked
func (hero Character) UseSelfTargetSpell() {

}

//given an array of Points the character memory of the labyrinth will be updated with new tiles
//that he remebers(tiles that will be displyed)
func(hero Character) MemorizeLabyrinth(points []Point) {

}

//values in the memory array will be updated by lowering the duration integer by one 
func(hero Character) UpdateMemory(){

}

type Mimic struct {
	currentHealth, maxHealth int
	x,y int
	orientation *Point
	symbol string
}

//the function returns the damage(integer) the character will do to an enemy
func (mimic Mimic) DoDamage() int {
	return 0
}

//the function substracs the funcyion's argument "damage" from the characters currentHealth
func (mimic Mimic) TakeDamage(damage int) {

}

//function will determine whether the character is hit or not
func (mimic Mimic) WillBeHit() bool {
	return false
}

//function determies whether the character will hit or not
func (mimic Mimic) WillHit() bool {
	return true
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