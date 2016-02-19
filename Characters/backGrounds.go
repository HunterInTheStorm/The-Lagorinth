package Character

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