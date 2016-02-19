package Spell

import "github.com/golang/The-Lagorinth/Point"
import "github.com/golang/The-Lagorinth/Labyrinth"
import "math/rand"

type Projectile struct {
	Symbol string
	SpellName string
	Location Point.Point
	Vector Point.Point
	WillStun, CanDestroyWall bool
	CritChance int
	Damage int
	Buff *Buff
}

func (spell Projectile) DoDamage() float32 {
	if rand.Intn(100) < spell.CritChance {
		return float32(2 * spell.Damage)
	}
	return float32(spell.Damage)
}

func (spell *Projectile) Move() {
	spell.Location.X += spell.Vector.X
	spell.Location.Y += spell.Vector.Y
}

func (spell *Projectile) ProjectileImapact(labyrinth *Labyrinth.Labyrinth) {
	if spell.CanDestroyWall {
		labyrinth.Labyrinth[spell.Location.X][spell.Location.Y] = Labyrinth.Pass
	}
}