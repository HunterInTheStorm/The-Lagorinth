//Handles the logic in the game.
package Game

import "github.com/golang/The-Lagorinth/Spells"
import "github.com/golang/The-Lagorinth/Characters"
import "github.com/golang/The-Lagorinth/Items"
import "time"
import "fmt"

//playerDefetedMessage prints a message when the player has been defeated.
func (game *Game) playerDefetedMessage() {
	fmt.Println("Player defeated")
	time.Sleep(2000 * time.Millisecond)
}

//monsterDefetedMessage prints a message when a monster has been defeated.
func (game *Game) monsterDefetedMessage(name string, playerName string) {
	fmt.Println("Player defeats monster")
	time.Sleep(2000 * time.Millisecond)
}

//avoidAttackMessage prints a message when a character avoids an attack.
func (game *Game) avoidAttackMessage(attackerName string, defenderName string) {
	fmt.Printf("%s dodges %s's attack",defenderName, attackerName)
	time.Sleep(2000 * time.Millisecond)
}

//takeDamageFromTrapMessage prints a message when the player takes damage from a trap.
func (game *Game) takeDamageFromTrapMessage(damage float32, attackerName string, defender *Character.NPC) {
	fmt.Printf("%s strikes %s for %v points of damage.\n", attackerName, defender.Name,
						 int(damage))
	time.Sleep(2000 * time.Millisecond)
}

//takeDamageMessage prints a message when a character takes damage from another character.
func (game *Game) takeDamageMessage(damage float32, attacker *Character.NPC, defender *Character.NPC) {
	fmt.Printf("%s strikes %s for %v points of damage. %s has %v HP left\n", attacker.Name, defender.Name,
						 int(damage),defender.Name, int(defender.CurrentHealth))
	time.Sleep(2000 * time.Millisecond)
}

//newArmorFoundMessage prints a message when a new armor has been found.
func (game *Game) newArmorFoundMessage() {
	fmt.Println("You found a new armor piece! :)\nWould you like to equip it?(y/n)")
}

//newWeaponFoundMessage prints a message when a new weapon has been found.
func (game *Game) newWeaponFoundMessage() {
	fmt.Println("You found a new weapon! :)\nWould you like to equip it?(y/n)")
}

//takeSpellDamageMessage prints a message when a character takes damage from a spell.
func (game *Game) takeSpellDamageMessage(damage float32, character *Character.NPC) {
	fmt.Printf("%s is hit for %v spell damage", character.Name, int(damage))
	time.Sleep(2000 * time.Millisecond)
}

//avoidSpellMessage prints a message when a character avoids a spell.
func (game *Game) avoidSpellMessage(character *Character.NPC) {
	fmt.Printf("%s dodges spell", character.Name)
	time.Sleep(2000 * time.Millisecond)
}

//useInstantSpellMessage prints a message when a character uses an instant spell.
func (game *Game) useInstantSpellMessage(spell *Spell.Spell, hero *Character.Hero) {
	fmt.Printf("%s uses %s.", hero.Base.Name, spell.SpellName)
	if spell.ManaCost > 0 {
		fmt.Printf("%s restores %v HP for %v MP", hero.Base.Name, spell.RegainHealth, spell.ManaCost)
	} else{
		fmt.Printf("%s restores %v MP for %v HP", hero.Base.Name, spell.ManaCost, spell.RegainHealth)
	}
	time.Sleep(2000 * time.Millisecond)
}

//spellOnCoolDownMessage prints a message when a character tries to use a spell that is on cool down.
func (game *Game) spellOnCoolDownMessage(spell *Spell.Spell) {
	if spell.CoolDownTimeLeft != 1 {
		fmt.Printf("%s is on CD for %v turns", spell.SpellName, spell.CoolDownTimeLeft)
	} else {
		fmt.Printf("%s is on CD for %v turn", spell.SpellName, spell.CoolDownTimeLeft)
	}
	time.Sleep(2000 * time.Millisecond)
}

//useBuffSpellMessage prints a message when a character uses a buff spell.
func (game *Game) useBuffSpellMessage(spell *Spell.Spell, hero *Character.Hero) {
	fmt.Printf("%s is now buffed with %s for %v turns", hero.Base.Name, spell.SpellName, spell.Duration)
	time.Sleep(2000 * time.Millisecond)
}

//buffFadeMessage prints a message when a buff's effect is over.
func (game *Game) buffFadeMessage(spell *Spell.Buff) {
	fmt.Printf("%s fades away", spell.BuffName)
	time.Sleep(2000 * time.Millisecond)
}

//lowManaMessage prints a message when a character low on mana tries to use a spell.
func (game *Game) lowManaMessage(spell *Spell.Spell) {
	fmt.Printf("Not enough MP to cast %s. Spell Requires %v MP", spell.SpellName, spell.ManaCost)
	time.Sleep(2000 * time.Millisecond)
}

//availableBackGroundsMessage prints a message with the backgrounds the player can choose from.
func (game *Game) availableBackGroundsMessage() {
	fmt.Println("And what are you renowned for?")
	fmt.Println("Giant-like strength(giant)")
	fmt.Println("Agile toreador(toreador)")
	fmt.Println("World Cartographer(cartographer)")
	fmt.Println("Wise Librarian(librarian)")
}

//askNameMessage prints a message that prompts the player for his name.
func (game *Game) askNameMessage() {
	fmt.Println("What is your name mighty adventurer?")
}

//askClassNameMessage prints a message with the classes available in the game.
func (game *Game) askClassNameMessage() {
	fmt.Println("What is your profession traveler?")
	fmt.Println("(Paladin/Mage/Rouge)")
}

//compareArmor prints a message containing the information of newly found armor and the currently equipped one.
func (game *Game) compareArmor(current *Items.Armor, found *Items.Armor) {
	fmt.Println("Your Armor\t\t\tNew Armor")
	fmt.Printf("+Health: %v\t\t\t+Health: %v\n",current.Health, found.Health)
	fmt.Printf ("+HP Regen: %v\t\t\t+HP Regen: %v\n",current.HealthRegen, found.HealthRegen)
	fmt.Printf("+Mana: %v\t\t\t+Mana: %v\n",current.Mana, found.Mana)
	fmt.Printf ("+MP Regen: %v\t\t\t+MP Regen: %v\n",current.ManaRegen, found.ManaRegen)
	fmt.Printf("+Defence: %v\t\t\t+Defence: %v\n",current.Defence, found.Defence)
	fmt.Printf("+Evasion: %v\t\t\t+Evasion: %v\n",current.Evasion, found.Evasion)
}

//compareWeapon prints a message containing the information of newly found weapon and the currently equipped one.
func (game *Game) compareWeapon(current *Items.Weapon, found *Items.Weapon) {
	fmt.Println("Your Weapon\t\t\tNew Weapon")
	fmt.Printf("Damage: %v-%v\t\t\tDamage: %v-%v\n", current.MinDmg, current.MaxDmg, found.MinDmg, found.MaxDmg)
	fmt.Printf("+BonusDmg: %v\t\t\t+BonusDmg: %v\n",current.BonusDmg, found.BonusDmg)
	fmt.Printf ("+CritChance: %v\t\t\t+CritChance: %v\n",current.BonusCritChance, found.BonusCritChance)
}
