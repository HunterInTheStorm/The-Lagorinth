package Items

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
	return rand.Intn(weapon.maxDmg + 1 - weapon.minDmg) + wepon.minDmg
}

type Armor struct {
	defence, health, evasion int
}
