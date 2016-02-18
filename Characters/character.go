package Character

import "github.com/golang/The-Lagorinth/Items"
import "github.com/golang/The-Lagorinth/Spells"
import "github.com/golang/The-Lagorinth/Point"
import "github.com/golang/The-Lagorinth/Labyrinth"
import "math/rand"

var PaladinClassName string = "Paladin"
var RougeClassName string = "Rouge"
var MageClassName string = "Mage"

var BackGroundNameGiant string = "giant"
var BackGroundNameToreador string =	"toreador"
var BackGroundNameCartographer string = "cartographer"
var BackGroundNameLibrarian string = "librarian"

// //this function will create one of the 3 classes for the player
func CreatePaladin(charName string, charBackGround string, x int, y int) *Hero {
	//NPC
	var dmgMultuplier float32 = 6.5
	var defence int = 6 
	var evasion int = 3
	var critChance int = 15
	var maxHealth float32 = 200.0
	var healthRegen float32 = 3.6
	var maxMana float32 = 75.0
	var manaRegen float32= 1.9
	var visionRadious int = 4
	var trapHandling int = 1

	//Character
	// SpellLList []Spells.Spell
	memory := make(map[Point.Point]int)
	var memoryDuration int = 17

	base := NPC{&Point.Point{x, y, nil}, Labyrinth.CharSymbol, charName, &Point.Point{1, 0, nil},
		nil, nil, dmgMultuplier, defence, evasion, critChance, maxHealth, maxHealth, healthRegen, maxMana,
		maxMana, manaRegen, visionRadious, false, make(map[int]*Spells.Buff), true, trapHandling}
	
	hero := Hero{&base, PaladinClassName, charBackGround, make([]*Spells.Spell, 0, 3), 
		memory, memoryDuration}
	return &hero
}

// //this function will create one of the 3 classes for the player
func CreateMage(charName string, charBackGround string, x int, y int) *Hero {
	//NPC
	var dmgMultuplier float32= 1.2
	var defence int = 6 
	var evasion int = 3
	var critChance int = 15
	var maxHealth float32 = 200.0
	var healthRegen float32 = 3.6
	var maxMana float32 = 75.0
	var manaRegen float32 = 1.9
	var visionRadious int = 4
	var trapHandling int = 3
	//Character
	// SpellLList []Spells.Spell
	memory := make(map[Point.Point]int)
	var memoryDuration int = 12

	base := NPC{&Point.Point{x, y, nil}, Labyrinth.CharSymbol, charName, &Point.Point{1, 0, nil},
		nil, nil, dmgMultuplier, defence, evasion, critChance, maxHealth, maxHealth, healthRegen, maxMana,
		maxMana, manaRegen, visionRadious, false, make(map[int]*Spells.Buff), true, trapHandling}
	
	hero := Hero{&base, MageClassName, charBackGround, make([]*Spells.Spell, 0, 3), 
		memory, memoryDuration}
	
	return &hero
}

// //this function will create one of the 3 classes for the player
func CreateRouge(charName string, charBackGround string, x int, y int) *Hero {
	//ROUGE
	//NPC
	var dmgMultuplier float32 = 1.2
	var defence int = 6 
	var evasion int = 3
	var critChance int = 15
	var maxHealth float32 = 200.0
	var healthRegen float32 = 3.6
	var maxMana float32 = 75.0
	var manaRegen float32 = 1.9
	var visionRadious int = 5
	var trapHandling int = 6

	//Character
	// SpellLList []Spells.Spell
	memory := make(map[Point.Point]int)
	var memoryDuration int = 10


	base := NPC{&Point.Point{x, y, nil}, Labyrinth.CharSymbol, charName, &Point.Point{1, 0, nil},
		nil, nil, dmgMultuplier, defence, evasion, critChance, maxHealth, maxHealth, healthRegen, maxMana,
		maxMana, manaRegen, visionRadious, false, make(map[int]*Spells.Buff), true, trapHandling}
	
	hero := Hero{&base, RougeClassName, charBackGround, make([]*Spells.Spell, 0, 3), 
		memory, memoryDuration}

	return &hero
}


func CreateBgGiant() *BackGround{
	var memoryDuration int = 0
	var visionRadius int = 0
	var mana float32 = 0.0
	var manaRegen float32 = 0.0
	var health float32 = 50.0
	var healthRegen float32 = 1.3
	var armor int = 0
	var evasion int = 0
	var critChance int = 0
	var dmgMultuplier float32 = 0.7

	background := BackGround{BackGroundNameGiant, memoryDuration, visionRadius, mana,
		manaRegen, health, healthRegen, armor, evasion, critChance, dmgMultuplier}

	return &background
}

func CreateBgToreador() *BackGround {
	var memoryDuration int = 0
	var visionRadius int = 1
	var mana float32 = 10.0
	var manaRegen float32 = 0.3 
	var health float32 = 10.0
	var healthRegen float32 = 0.3
	var armor int = 0
	var evasion int = 5
	var critChance int = 5
	var dmgMultuplier float32 = 0.3

	background := BackGround{BackGroundNameToreador, memoryDuration, visionRadius, mana,
		manaRegen, health, healthRegen, armor, evasion, critChance, dmgMultuplier}

	return &background
}

func CreateBgCartographer() *BackGround {
	var memoryDuration int = 20
	var visionRadius int = 3
	var mana float32 = 30.0
	var manaRegen float32 = 0.8
	var health float32 = 30.0
	var healthRegen float32 = 0.8
	var armor int = 0
	var evasion int = 0
	var critChance int = 0
	var dmgMultuplier float32 = 0

	background := BackGround{BackGroundNameCartographer, memoryDuration, visionRadius, mana,
		manaRegen, health, healthRegen, armor, evasion, critChance, dmgMultuplier}

	return &background
}

func CreateBgLibrarian() *BackGround {
	var memoryDuration int = 12
	var visionRadius int = 0
	var mana float32 = 60.0
	var manaRegen float32 = 3.5
	var health float32 = 0
	var healthRegen float32 = 0
	var armor int = 0
	var evasion int = 0
	var critChance int = 0
	var dmgMultuplier float32 = 0.2

	background := BackGround{BackGroundNameLibrarian, memoryDuration, visionRadius, mana,
		manaRegen, health, healthRegen, armor, evasion, critChance, dmgMultuplier}

	return &background
}

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

func (npc *NPC) moveTowardsHero(labyrinth *Labyrinth.Labyrinth) (bool, *Point.Point) {
	upTile :=	labyrinth.Labyrinth[npc.Location.X - 1][npc.Location.Y]
	if upTile == Labyrinth.CharSymbol {
		return true, &Point.Point{npc.Location.X - 1, npc.Location.Y, nil}
	}
	downTile := labyrinth.Labyrinth[npc.Location.X + 1][npc.Location.Y]
	if downTile == Labyrinth.CharSymbol {
		return true, &Point.Point{npc.Location.X + 1, npc.Location.Y, nil}
	}
	leftTile := labyrinth.Labyrinth[npc.Location.X][npc.Location.Y - 1]
	if leftTile == Labyrinth.CharSymbol {
		return true, &Point.Point{npc.Location.X, npc.Location.Y - 1, nil}
	}
	rightTile := labyrinth.Labyrinth[npc.Location.X][npc.Location.Y + 1]
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

func (npc *NPC) Regenerate() {
	npc.RegenMana()
	npc.RegenHealth()
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
func (hero *Hero) SwapWeapon(weapon *Items.Weapon) {
	hero.Base.UnequipWeapon()
	hero.Base.EquipWeapon(weapon)
}

// //will replace the bonuses given from one armor for the bonuses of another
func (hero *Hero) SwapArmor(armor *Items.Armor) {
	hero.Base.UnequipArmor()
	hero.Base.EquipArmor(armor)
}

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

// //a spell from the list of targetble spells will be envoked
// func (hero *Hero) UseProjectileSpell() Projectile {

// }

// //a spell from the list of selfcast spells will be envoked
// func (hero *Hero) UseSelfTargetSpell() Buff {

// }

// func (hero *Hero) UseAreaOfEffectSpell() Effect {

// }

func (hero *Hero) UpdateMemory() {
	for point, _ := range hero.Memory {
		if hero.Memory[point] > -1 {
			hero.Memory[point]--
		} else {
			delete(hero.Memory, point)
		}
	}
}

// //given an array of Points the character memory of the labyrinth will be updated with new tiles
// //that he remebers(tiles that will be displyed)
func (hero *Hero) MemorizeLabyrinth(labyrinth *Labyrinth.Labyrinth, center *Point.Point) {
	var minX int = center.X - hero.Base.VisionRadious
	var maxX int = center.X + hero.Base.VisionRadious
	var minY int = center.Y - hero.Base.VisionRadious
	var maxY int = center.Y + hero.Base.VisionRadious
	var y int
	for currentY := minY; currentY <= maxY; currentY++ {
		for xAscend := center.X; xAscend <= maxX; xAscend++ {
			y = lineEquationRegardsToX(xAscend, center.X, center.Y, maxX, currentY)
			if labyrinth.IsInBondaries(xAscend, y) && labyrinth.Labyrinth[xAscend][y] != Labyrinth.Wall {
				hero.Memory[Point.Point{xAscend, y, nil}] = hero.MemoryDuration
			} else {
				hero.Memory[Point.Point{xAscend, y, nil}] = hero.MemoryDuration
				break 
			}
		}
		for xDescend := center.X; xDescend >= minX; xDescend-- {
			y = lineEquationRegardsToX(xDescend, center.X, center.Y, minX, currentY)
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
			x = lineEquationRegardsToY(yDescend, center.X, center.Y, currentX, minY)
			if labyrinth.IsInBondaries(x, yDescend) && labyrinth.Labyrinth[x][yDescend] != Labyrinth.Wall {
				hero.Memory[Point.Point{x, yDescend, nil}] = hero.MemoryDuration
			} else {
				hero.Memory[Point.Point{x, yDescend, nil}] = hero.MemoryDuration
				break 
			}
		}
		for yAscend := center.Y; yAscend <= maxY; yAscend++ {
			x = lineEquationRegardsToY(yAscend, center.X, center.Y, currentX, maxY)
			if labyrinth.IsInBondaries(x, yAscend) && labyrinth.Labyrinth[x][yAscend] != Labyrinth.Wall {
				hero.Memory[Point.Point{x, yAscend, nil}] = hero.MemoryDuration
			} else {
				hero.Memory[Point.Point{x, yAscend, nil}] = hero.MemoryDuration
				break 
			}
		}
	}
}

func lineEquationRegardsToX(x int, x0 int, y0 int, x1 int, y1 int) int{
	 y := float32((x - x0)*(y1 - y0)/(x1 - x0) + y0)
	 return int(y + 0.5)
}

func lineEquationRegardsToY(y int, x0 int, y0 int, x1 int, y1 int) int{
	 x := float32((y - y0)*(x1 - x0)/(y1 - y0) + x0)
	 return int(x + 0.5)
}

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
	trap.MinDmg = rand.Intn(20) + 10
	trap.MaxDmg = rand.Intn(15) + trap.MinDmg
}

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

type BackGround struct {
	Name string
	BonusMemoryDuration int
	BonusVisionRadius int
	BonusMana float32
	BonusManaRegen float32
	BonusHealth float32
	BonusHealthRegen float32
	BonusArmor int
	BonusEvasion int
	BonusCritChance int
	BonusDmgMultuplier float32
}