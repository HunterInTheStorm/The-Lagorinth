package main

//This structure has 4 fields:
//minimum and maximum damage of the weapon represented by integers
//integer for additional bonus damage
//integer for chance to do a critical hit
type Weapon struct {
	minDmg, maxDmg, bonusDmg, bonusCritChance int
	//rng int
}

//return and integer between minimum and maximum damage
func (weapon Weapon) Damage() int {

}

//description of the item to be ppresented to the player while playing the game
//MOVE TO INTERFACE
func (weapon Weapon) Description() string {
	return ""
}


type Armor struct {
	defence, health, evasion int
}

//description of the item to be ppresented to the player while playing the game
//MOVE TO INTERFACE
func (armor Armor) Description() string {
	return ""
}