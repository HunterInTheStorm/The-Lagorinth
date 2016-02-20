//package Spell handles the creation and functionalities of the spells in the game.
package Spell

import "github.com/golang/The-Lagorinth/Point"
import "github.com/golang/The-Lagorinth/Labyrinth"
import "math/rand"

type Projectile struct {
	Symbol                   string
	SpellName                string
	Location                 Point.Point
	Vector                   Point.Point
	WillStun, CanDestroyWall bool
	CritChance               int
	Damage                   int
	Buff                     *Buff
}

//DoDamage returns the damage that will be dealt to a character.
//Depending on the projectile's critical chance the function my return double the damage.
func (spell Projectile) DoDamage() float32 {
	if rand.Intn(100) < spell.CritChance {
		return float32(2 * spell.Damage)
	}
	return float32(spell.Damage)
}

//Move updates the the coordinates of the projectile.
//The coordinates are updated via the vector field, which stores the direction the projectile is facing.
func (spell *Projectile) Move() {
	spell.Location.X += spell.Vector.X
	spell.Location.Y += spell.Vector.Y
}

//ProjectileImapact handles the collision of a projectile with a wall in the labyrinth.
//If the CanDestroyWall flag is set, the wall will be replace with a passage.
func (spell *Projectile) ProjectileImapact(labyrinth *Labyrinth.Labyrinth) {
	if spell.CanDestroyWall {
		labyrinth.Labyrinth[spell.Location.X][spell.Location.Y] = Labyrinth.Pass
	}
}
