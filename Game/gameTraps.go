package Game

import "github.com/golang/The-Lagorinth/Labyrinth"
import "github.com/golang/The-Lagorinth/Characters"
import "github.com/golang/The-Lagorinth/Point"
import "time"
import "fmt"
import "math/rand"


func (game *Game) triggerDamageTrap(trap *Character.Trap, character *Character.NPC) {
	damage := trap.DamageTrap()
	if rand.Intn(100) < character.Evasion {
		game.avoidAttackMessage(character.Name, Character.TrapTypes[0])
	} else {
		character.TakeDamage(damage)
		game.takeDamageFromTrapMessage(damage, Character.TrapTypes[0], character)
	}
	if game.isCharacterDefeted(character) {
		game.CharacterDefeted(character, -1)
	}
}

func (game *Game) findEmptyTile(centerX int, centerY int) Point.Point{
	for e := 0; true; e++ {
		for i := centerX - 1 - e; i <= centerX + 1 + e; i++ {
			for j := centerY - 1 - e; j <= centerY + 1 + e; j++ {
				if i > -1 && j > -1 && game.labyrinth.Labyrinth[i][j] == Labyrinth.Pass {
					return Point.Point{i, j, nil}
				}
			}
		}
	}
	return Point.Point{}
}

func (game *Game) triggerTabulaRasaTrap(trap *Character.Trap, hero *Character.Hero) {
	game.triggerMemoryWhipeTrap(trap, hero)
	game.triggerTeleportTrap(trap, hero.Base)
}

func (game *Game) triggerMemoryWhipeTrap(trap *Character.Trap, hero *Character.Hero) {
	fmt.Println("MEMORY TRAP")
	time.Sleep(2000 * time.Millisecond)
	trap.WhipeMemory(hero)
}

func (game *Game) triggerTeleportTrap(trap *Character.Trap, character *Character.NPC) {
	fmt.Println("TELEPORT TRAP")
	time.Sleep(2000 * time.Millisecond)
	location := trap.NewLocation(game.labyrinth.Width, game.labyrinth.Height)
	if game.labyrinth.Labyrinth[location.X][location.Y] == Labyrinth.Pass {
		character.Location.X = location.X
		character.Location.Y = location.Y
	} else {
		location = game.findEmptyTile(location.X, location.Y)
		character.Location.X = location.X
		character.Location.Y = location.Y
	}
	game.cameraReset()
}

func (game *Game) triggerSpawnTrap(trap *Character.Trap) {
	fmt.Println("SPAWNTRAP ACTIVATED")
	time.Sleep(2000 * time.Millisecond)
	location := trap.NewLocation(game.labyrinth.Width, game.labyrinth.Height)
	if game.labyrinth.Labyrinth[location.X][location.Y] == Labyrinth.Pass {
		newMonster := game.createMonster(location.X, location.Y)
		game.monsterList = append(game.monsterList, &newMonster)
	} else {
		location = game.findEmptyTile(location.X, location.Y)			
		newMonster := game.createMonster(location.X, location.Y)
		game.monsterList = append(game.monsterList, &newMonster)
	}
	fmt.Printf("Monster Spawned at %v,%v", location.X, location.Y)
	time.Sleep(2000 * time.Millisecond)
}

func (game *Game) triggerTrap(trap *Character.Trap, character *Character.Hero) {
	switch trap.TrapType {
	case Character.TrapTypes[0]:
		game.triggerDamageTrap(trap, character.Base)
	case Character.TrapTypes[1]:
		game.triggerSpawnTrap(trap)
	case Character.TrapTypes[2]:
		game.triggerTeleportTrap(trap, character.Base)
		game.player.MemorizeLabyrinth(game.labyrinth, game.player.Base.Location)
		game.draw()
	case Character.TrapTypes[3]:
		game.triggerMemoryWhipeTrap(trap, character)
		game.player.MemorizeLabyrinth(game.labyrinth, game.player.Base.Location)
		game.draw()
	case Character.TrapTypes[4]:
		game.triggerTabulaRasaTrap(trap, character)
		game.player.MemorizeLabyrinth(game.labyrinth, game.player.Base.Location)
		game.draw()
	}
}

func (game *Game) checkTraps() {
	if trap, ok := game.isTrapTriggered(game.player.Base) ; ok {
		game.triggerTrap(trap, game.player)
	}
}

func (game *Game) isTrapTriggered(character *Character.NPC) (*Character.Trap, bool) {
	if trap, ok := game.trapList[Point.Point{character.Location.X, character.Location.Y, nil}]; ok {
    	return trap, true
	}
	return &Character.Trap{}, false
}

func (game *Game) removeTrap(trap *Character.Trap) {
	game.restoreTile(trap.Location.X, trap.Location.Y)
	delete(game.trapList, *trap.Location)
}

func (game *Game) calculateOddsVsTraps(difficulty int, trapHandlingSkill int) int {
	return 100 - difficulty * 10 + trapHandlingSkill * 5
}

func (game *Game) attempDisarmTrap(trap *Character.Trap, character *Character.NPC) {
	chance := game.calculateOddsVsTraps(trap.DisarmDifficulty, character.TrapHandling)
	if rand.Intn(100) < chance {
		fmt.Println("TRAP DISARMED!. HELL YEAH!!!!!!!!!!")
		time.Sleep(2000 * time.Millisecond)
		trap.IsDisarmed = true
		trap.CanBeDisarmed = false
		game.trapsDisarmed ++
		game.removeTrap(trap)
	} else {
		fmt.Println("YOU ARE SUCH A DISAPPOINTMENT")
		time.Sleep(2000 * time.Millisecond)
		trap.CanBeDisarmed = false
	}
}

func (game *Game) attempDetectTrap(trap *Character.Trap, character *Character.NPC) {
	chance := game.calculateOddsVsTraps(trap.DetectDifficulty, character.TrapHandling)
	if rand.Intn(100) < chance {
		fmt.Println("TRAP DETECTED!!!!!!!")
		time.Sleep(2000 * time.Millisecond)
		trap.IsDetected = true
		trap.CanBeDetected = false
	} else {
		fmt.Println("TRAP NOT DETECTED :(")
		time.Sleep(2000 * time.Millisecond)
		trap.CanBeDetected = false
	}
}

func (game *Game) encounterTrap(character *Character.NPC, x int, y int) {
	trap := game.trapList[Point.Point{x, y, nil}]
	if trap.CanBeDetected && !trap.IsDetected {
		game.attempDetectTrap(trap, character)
		if !trap.IsDetected {
			game.characterMoveTo(character, x, y)
		} 
	} else if trap.IsDetected && trap.CanBeDisarmed {
		fmt.Println("Do you want to disarm the trap(y/n)")
		answer := game.detectKeyPress()
		if answer == "y"{
			game.attempDisarmTrap(trap, character)
		} else if answer == "n" {
			game.characterMoveTo(character, x, y)
		}
	} else if !trap.IsDetected || !trap.IsDisarmed{
		game.characterMoveTo(character, x, y)
	} else {
		game.characterMoveTo(character, x, y)
	}
}