//Package Character handles the creation and management of characters in the game.
package Character

import "github.com/golang/The-Lagorinth/Point"
import "math/rand"

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

//Randomize creates a trap at random.
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

//DamageTrap returns the damage a trap will do.
func (trap *Trap) DamageTrap() float32 {
	damageRange := trap.MaxDmg - trap.MinDmg
 	return float32(rand.Intn(damageRange) + trap.MinDmg)
}

//NewLocation return a random point.
func (trap Trap) NewLocation(labWidth int, labHeight int) Point.Point {
 	return Point.Point{rand.Intn(labWidth), rand.Intn(labHeight), nil}
}

//WhipeMemory deletes the memory of the player's character
func (trap Trap) WhipeMemory(hero *Hero) {
	newMemory := make(map[Point.Point]int)
	hero.Memory = newMemory
}