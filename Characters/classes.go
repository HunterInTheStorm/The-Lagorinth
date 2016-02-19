package Character

import "github.com/golang/The-Lagorinth/Spells"
import "github.com/golang/The-Lagorinth/Point"
import "github.com/golang/The-Lagorinth/Labyrinth"

// //this function will create one of the 3 classes for the player
func CreatePaladin(charName string, charBackGround string, x int, y int) *Hero {
	//NPC
	location := Point.Point{x, y, nil}
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
	spellList := make([]*Spell.Spell, 0, 3)
	spell1 := Spell.PaladinSpellHeal(&location)
	spellList = append(spellList, spell1)
	spell2 := Spell.PaladinSpellHolyArmor(&location)
	spellList = append(spellList, spell2)
	spell3 := Spell.PaladinSpellHolyBolt(&location)
	spellList = append(spellList, spell3)
	memory := make(map[Point.Point]int)
	var memoryDuration int = 30

	base := NPC{&location, Labyrinth.CharSymbol, charName, &Point.Point{1, 0, nil},
		nil, nil, dmgMultuplier, defence, evasion, critChance, maxHealth, maxHealth, healthRegen, maxMana,
		maxMana, manaRegen, visionRadious, false, make(map[int]*Spell.Buff), true, trapHandling}
	
	hero := Hero{&base, PaladinClassName, charBackGround, spellList, 
		memory, memoryDuration}
	return &hero
}

// //this function will create one of the 3 classes for the player
func CreateMage(charName string, charBackGround string, x int, y int) *Hero {
	//NPC
	location := Point.Point{x, y, nil}
	var dmgMultuplier float32= 6.0
	var defence int = 4 
	var evasion int = 6
	var critChance int = 20
	var maxHealth float32 = 150.0
	var healthRegen float32 = 2.0
	var maxMana float32 = 200
	var manaRegen float32 = 5.0
	var visionRadious int = 5
	var trapHandling int = 5
	//Character
	// SpellLList []Spells.Spell
	spellList := make([]*Spell.Spell, 0, 3)
	spell1 := Spell.MageSpellSacrifice(&location)
	spellList = append(spellList, spell1)
	spell2 := Spell.MageSpellBallLightning(&location)
	spellList = append(spellList, spell2)
	spell3 := Spell.MageSpellFireBall(&location)
	spellList = append(spellList, spell3)
	memory := make(map[Point.Point]int)
	var memoryDuration int = 45

	base := NPC{&location, Labyrinth.CharSymbol, charName, &Point.Point{1, 0, nil},
		nil, nil, dmgMultuplier, defence, evasion, critChance, maxHealth, maxHealth, healthRegen, maxMana,
		maxMana, manaRegen, visionRadious, false, make(map[int]*Spell.Buff), true, trapHandling}
	
	hero := Hero{&base, MageClassName, charBackGround, make([]*Spell.Spell, 0, 3), 
		memory, memoryDuration}
	
	return &hero
}

// //this function will create one of the 3 classes for the player
func CreateRouge(charName string, charBackGround string, x int, y int) *Hero {
	//ROUGE
	//NPC
	location := Point.Point{x, y, nil}
	var dmgMultuplier float32 = 7.5
	var defence int = 5 
	var evasion int = 10
	var critChance int = 25
	var maxHealth float32 = 175.0
	var healthRegen float32 = 2.0
	var maxMana float32 = 150.0
	var manaRegen float32 = 2.5
	var visionRadious int = 5
	var trapHandling int = 6

	//Character
	spellList := make([]*Spell.Spell, 0, 3)
	spell1 := Spell.RougeSpellPrecision(&location)
	spellList = append(spellList, spell1)
	spell2 := Spell.RougeSpellShadow(&location)
	spellList = append(spellList, spell2)
	spell3 := Spell.RougeSpellAssassinMark(&location)
	spellList = append(spellList, spell3)
	memory := make(map[Point.Point]int)
	var memoryDuration int = 30


	base := NPC{&location, Labyrinth.CharSymbol, charName, &Point.Point{1, 0, nil},
		nil, nil, dmgMultuplier, defence, evasion, critChance, maxHealth, maxHealth, healthRegen, maxMana,
		maxMana, manaRegen, visionRadious, false, make(map[int]*Spell.Buff), true, trapHandling}
	
	hero := Hero{&base, RougeClassName, charBackGround, make([]*Spell.Spell, 0, 3), 
		memory, memoryDuration}

	return &hero
}
