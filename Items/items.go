package Items

import "math/rand"

//This structure has 4 fields:
//minimum and maximum damage of the weapon represented by integers
//integer for additional bonus damage
//integer for chance to do a critical hit
type Weapon struct {
	MinDmg, MaxDmg, BonusDmg, BonusCritChance int
	//rng int
}


//REPLACE MAGIC NUMBERS
func (weapon *Weapon) RandomizeWeapon() {
	weapon.MinDmg = rand.Intn(5) + 10
	weapon.MaxDmg = weapon.MinDmg + rand.Intn(15)
	weapon.BonusDmg = rand.Intn(8) + 1
	weapon.BonusCritChance = rand.Intn(11)/10
}

//return and integer between minimum and maximum damage
func (weapon Weapon) Damage() float32 {
	return float32(rand.Intn(weapon.MaxDmg + 1 - weapon.MinDmg) + weapon.MinDmg + weapon.BonusDmg)
}

type Armor struct {
	Defence, Evasion int
	Health, HealthRegen float32
	Mana, ManaRegen float32
	//mana regeneration
	//reduced mana cost
}

//REPLACE MAGIC NUMBERS
func (armor *Armor) RandomizeArmor() {
	armor.Defence = rand.Intn(6) + 8
	armor.Evasion = rand.Intn(6)
	armor.Health = float32(rand.Intn(151)) + 50.0
	armor.HealthRegen = float32(rand.Intn(5)) + float32(rand.Intn(11))/float32(10)
	armor.Mana = float32(rand.Intn(76)) + 25.0
	armor.ManaRegen =	float32(rand.Intn(3)) + float32(rand.Intn(6))/float32(10)
}
