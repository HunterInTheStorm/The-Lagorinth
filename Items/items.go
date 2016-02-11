package Items

//This structure has 4 fields:
//minimum and maximum damage of the weapon represented by integers
//integer for additional bonus damage
//integer for chance to do a critical hit
type Weapon struct {
	minDmg, maxDmg, bonusDmg, bonusCritChance int
	//rng int
}


//REPLACE MAGIC NUMBERS
func (weapon *Weapon) RandomizeWeapon() {
	weapon.minDmg = rand.Intn(5) + 10
	weapon.maxDmg = weapon.minDmg + rand.Intn(15)
	weapon.bonusDmg = rand.Intn(8) + 1
	weapon.bonusCritChance = float32(rand.Intn(11)/float32(10)
}

//return and integer between minimum and maximum damage
func (weapon Weapon) Damage() float32 {
	return float32(rand.Intn(weapon.maxDmg + 1 - weapon.minDmg) + wepon.minDmg + weapon.bonusDmg)
}

type Armor struct {
	defence, evasion int
	health, healthRegen float32
	mana, manaRegen float32
	//mana regeneration
	//reduced mana cost
}

//REPLACE MAGIC NUMBERS
func (armor *Armor) RandomizeArmor() {
	armor.defence = rand.Intn(6) + 8
	armor.evasion = rand.Intn(6)
	armor.health = float32(rand.Intn(151)) + 50.0
	armor.healthRegen = float32(rand.Intn(5)) + float32(rand.Intn(11)/float32(10)
	armor.mana = float32(rand.Intn(76)) + 25.0
	armor.manaRegen =	float32(rand.Intn(3)) + float32(rand.Intn(6)/float32(10)
}
