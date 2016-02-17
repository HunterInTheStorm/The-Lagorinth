package Character

import "github.com/golang/The-Lagorinth/Items"
import "github.com/golang/The-Lagorinth/Spells"
import "github.com/golang/The-Lagorinth/Point"
import "math/rand"

// import "github.com/golang/The-Lagorinth/Items"
// import "github.com/golang/The-Lagorinth/Spells"

var PaladinClassName string = "Paladin"
//PALADIN
//NPC
var PaladinDmgMultuplier float32 = 6.5
var PaladinDefence int = 6 
var PaladinEvasion int = 3
var PaladinCritChance int = 15
var PaladinMaxHealth float32 = 200.0
var PaladinHealthRegen float32 = 3.6
var PaladinMaxMana float32 = 75.0
var PaladinManaRegen float32= 1.9
var PaladinVisionRadious int = 4
var PaladinTrapHandling int = 1

//Character
	// SpellLList []Spells.Spell
	// Memory map[Point.Point]int
var PaladinMemoryDuration int = 10

var MageClassName string = "Mage"
//MAGE
//NPC
var MageDmgMultuplier float32= 1.2
var MageDefence int = 6 
var MageEvasion int = 3
var MageCritChance int = 15
var MageMaxHealth float32 = 200.0
var MageHealthRegen float32 = 3.6
var MageMaxMana float32 = 75.0
var MageManaRegen float32 = 1.9
var MageVisionRadious int = 4
var MageTrapHandling int = 3
//Character
	// SpellLList []Spells.Spell
	// Memory map[Point.Point]int
var MageMemoryDuration int = 12

var RougeClassName string = "Rouge"
//ROUGE
//NPC
var RougeDmgMultuplier float32 = 1.2
var RougeDefence int = 6 
var RougeEvasion int = 3
var RougeCritChance int = 15
var RougeMaxHealth float32 = 200.0
var RougeHealthRegen float32 = 3.6
var RougeMaxMana float32 = 75.0
var RougeManaRegen float32 = 1.9
var RougeVisionRadious int = 5
var RougeTrapHandling int = 6

//Character
	// SpellLList []Spells.Spell
	// Memory map[Point.Point]int
var RougeMemoryDuration int = 10




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
	Buffs map[int]*Spells.Buff
	IsHuman bool
	TrapHandling int
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
	npc.Location.X = npc.Location.X - 1
	npc.ChangeOrientation(npc.Location.X, npc.Location.X - 1, npc.Location.Y, npc.Location.Y)
}

//function that updates the coordinates of an character so that he move one down
func (npc *NPC) MoveSouth() {
	npc.Location.X = npc.Location.X + 1
	npc.ChangeOrientation(npc.Location.X, npc.Location.X + 1, npc.Location.Y, npc.Location.Y)
}

//function that updates the coordinates of an character so that he move one left
func (npc *NPC) MoveWest() {
	npc.Location.Y = npc.Location.Y - 1
	npc.ChangeOrientation(npc.Location.X, npc.Location.X, npc.Location.Y, npc.Location.Y - 1)
}

//function that updates the coordinates of an character so that he move one right
func (npc *NPC) MoveEast() {
	npc.Location.Y = npc.Location.Y + 1
	npc.ChangeOrientation(npc.Location.X, npc.Location.X, npc.Location.Y, npc.Location.Y + 1)
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

func (npc *NPC) CombinedDefence() float32 {
	return float32(npc.Defence + npc.Armor.Defence)
}

//the function substracs the funcyion's argument "damage" from the characters currentHealth
func (npc *NPC) TakeDamage(damage float32) {
	var damageTaken float32 = damage - npc.CombinedDefence()
	if damage > 0 {
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
	IsDetected bool
	CanBeDisarmed bool
	CanBeDetected bool
	MinDmg int
	MaxDmg int
}

func (trap *Trap) Randomize(loc *Point.Point) {
	trap.Location = loc
	trap.TrapType = TrapTypes[rand.Intn(len(TrapTypes))]
	trap.DetectDifficulty = rand.Intn(10) + 1
	trap.DisarmDifficulty = rand.Intn(10) + 1
	trap.IsDisarmed = false
	trap.IsDetected = false
	trap.CanBeDisarmed = true
	trap.CanBeDetected = true
	trap.MinDmg = rand.Intn(6)
	trap.MaxDmg = rand.Intn(6) + trap.MinDmg
}

// //the function returns the damage(integer) the character will do to an enemy
// func (trap Trap) DoDamage() int {
// 	return 99
// }

func (trap *Trap) DamageTrap() float32 {
	damageRange := trap.MaxDmg - trap.MinDmg
 	return float32(rand.Intn(damageRange) + trap.MinDmg)
}

// //the function will return a point for where an enemy will spawn
func (trap Trap) NewLocation(labWidth int, labHeight int) Point.Point {
 	return Point.Point{rand.Intn(labWidth), rand.Intn(labHeight), nil}
}


func (trap Trap) WhipeMemory(hero *Hero) {
	newMemory := make(map[Point.Point]int)
	hero.Memory = newMemory
}
