package Game

import "github.com/golang/The-Lagorinth/Spells"
import "github.com/golang/The-Lagorinth/Characters"
import "github.com/golang/The-Lagorinth/Items"
import "time"
import "fmt"

func (game *Game) playerDefetedMessage() {
	fmt.Println("Player defeted")
	time.Sleep(2000 * time.Millisecond)
}

// //function will remove a monster from the monster list
func (game *Game) monsterDefetedMessage(name string, playerName string) {
	fmt.Println("Player defets monster")
	time.Sleep(2000 * time.Millisecond)
}

func (game *Game) avoidAttackMessage(attackerName string, defenderName string) {
	fmt.Printf("%s dodges %s's attack",defenderName, attackerName)
	time.Sleep(2000 * time.Millisecond)
}

func (game *Game) takeDamageFromTrapMessage(damage float32, attackerName string, defender *Character.NPC) {
	fmt.Printf("%s strikes %s for %v points of damage.\n", attackerName, defender.Name,
						 int(damage))
	time.Sleep(2000 * time.Millisecond)
}

func (game *Game) takeDamageMessage(damage float32, attacker *Character.NPC, defender *Character.NPC) {
	fmt.Printf("%s strikes %s for %v points of damage. %s has %v HP left\n", attacker.Name, defender.Name,
						 int(damage),defender.Name, int(defender.CurrentHealth))
	time.Sleep(2000 * time.Millisecond)
}

func (game *Game) newArmorFoundMessage() {
	fmt.Println("You found a new armor piece! :)\nWould you like to equip it?(y/n)")
}

func (game *Game) newWeaponFoundMessage() {
	fmt.Println("You found a new weapon! :)\nWould you like to equip it?(y/n)")
}

func (game *Game) takeSpellDamageMessage(damage float32, character *Character.NPC) {
	fmt.Printf("%s is hit for %v spell damage", character.Name, int(damage))
	time.Sleep(2000 * time.Millisecond)
}

func (game *Game) avoidSpellMessage(character *Character.NPC) {
	fmt.Printf("%s dodges spell", character.Name)
	time.Sleep(2000 * time.Millisecond)
}

func (game *Game) useInstantSpellMessage(spell *Spell.Spell, hero *Character.Hero) {
	fmt.Printf("%s uses %s.", hero.Base.Name, spell.SpellName)
	if spell.ManaCost > 0 {
		fmt.Printf("%s restores %v HP for %v MP", hero.Base.Name, spell.RegainHealth, spell.ManaCost)
	} else{
		fmt.Printf("%s restores %v MP for %v HP", hero.Base.Name, spell.ManaCost, spell.RegainHealth)
	}
	time.Sleep(2000 * time.Millisecond)
}

func (game *Game) spellOnCoolDownMessage(spell *Spell.Spell) {
	if spell.CoolDownTimeLeft != 1 {
		fmt.Printf("%s is on CD for %v turns", spell.SpellName, spell.CoolDownTimeLeft)
	} else {
		fmt.Printf("%s is on CD for %v turn", spell.SpellName, spell.CoolDownTimeLeft)
	}
	time.Sleep(2000 * time.Millisecond)
}

func (game *Game) useBuffSpellMessage(spell *Spell.Spell, hero *Character.Hero) {
	fmt.Printf("%s is now buffed with %s for %v turns", hero.Base.Name, spell.SpellName, spell.Duration)
	time.Sleep(2000 * time.Millisecond)
}

func (game *Game) buffFadeMessage(spell *Spell.Buff) {
	fmt.Printf("%s fades away", spell.BuffName)
	time.Sleep(2000 * time.Millisecond)
}

func (game *Game) lowManaMessage(spell *Spell.Spell) {
	fmt.Printf("Not enough MP to cast %s. Spell Requires %v MP", spell.SpellName, spell.ManaCost)
	time.Sleep(2000 * time.Millisecond)
}

func (game *Game) availableBackGroundsMessage() {
	fmt.Println("And what are you renowned for?")
	fmt.Println("Giant-like strenght(giant)")
	fmt.Println("Aagile toreador(toreador)")
	fmt.Println("World Cartographer(cartographer)")
	fmt.Println("Wise Librarian(librarian)")
}

func (game *Game) askNameMessage() {
	fmt.Println("What is your name mighty adventurer?")
}

func (game *Game) askClassNameMessage() {
	fmt.Println("What is your profession traveller?")
	fmt.Println("(Paladin/Mage/Rouge)")
}

func (game *Game) compareArmor(current *Items.Armor, found *Items.Armor) {
	fmt.Println("Your Armor\t\t\tNew Armor")
	fmt.Printf("+Health: %v\t\t\t+Health: %v\n",current.Health, found.Health)
	fmt.Printf ("+HP Regen: %v\t\t\t+HP Regen: %v\n",current.HealthRegen, found.HealthRegen)
	fmt.Printf("+Mana: %v\t\t\t+Mana: %v\n",current.Mana, found.Mana)
	fmt.Printf ("+MP Regen: %v\t\t\t+MP Regen: %v\n",current.ManaRegen, found.ManaRegen)
	fmt.Printf("+Defence: %v\t\t\t+Defence: %v\n",current.Defence, found.Defence)
	fmt.Printf("+Evasion: %v\t\t\t+Evasion: %v\n",current.Evasion, found.Evasion)
}

func (game *Game) compareWeapon(current *Items.Weapon, found *Items.Weapon) {
	fmt.Println("Your Weapon\t\t\tNew Weapon")
	fmt.Printf("Damage: %v-%v\t\t\tDamage: %v-%v\n", current.MinDmg, current.MaxDmg, found.MinDmg, found.MaxDmg)
	fmt.Printf("+BonusDmg: %v\t\t\t+BonusDmg: %v\n",current.BonusDmg, found.BonusDmg)
	fmt.Printf ("+CritChance: %v\t\t\t+CritChance: %v\n",current.BonusCritChance, found.BonusCritChance)
}
