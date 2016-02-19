package Game

import "github.com/golang/The-Lagorinth/Labyrinth"
import "github.com/golang/The-Lagorinth/Characters"
import "github.com/golang/The-Lagorinth/Point"
import "fmt"
import "os"
import "os/exec"

func (game *Game) draw() {
	game.clearScreen()
	game.drawCharacters()
	game.drawLabyrinth()
	game.drawHeroStats(game.player.Base)
}

// //function replaces an element fro the 2d array for the maze with the character symbol
func (game *Game) drawCharacters() {
	game.drawTraps()
	game.drawHero()
	game.drawProjectiles()
	game.drawMonsters()
}

//function to draw the labyrinth
func (game *Game) drawLabyrinth() {
	var maxX int = game.camera.X + game.cameraRadius
	var minX int = game.camera.X - game.cameraRadius
	var maxY int = game.camera.Y + game.cameraRadius
	var minY int = game.camera.Y - game.cameraRadius
	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			if game.labyrinth.IsInBondaries(i, j) {
				if point, ok := game.player.Memory[Point.Point{i, j, nil}]; ok && point > -1 {
					if game.labyrinth.Labyrinth[i][j] == Labyrinth.Trap {
						trap := game.trapList[Point.Point{i, j, nil}]
						if trap.IsDetected {
							fmt.Print(game.labyrinth.Labyrinth[i][j])
						} else {
							fmt.Print(Labyrinth.Pass)
						}
					} else {
						fmt.Print(game.labyrinth.Labyrinth[i][j])
					}
				} else {
					fmt.Print("+")
				}
			} else {
				fmt.Print("-")
			}
		}
		fmt.Println()
	}
}

func (game *Game) drawHeroStats(hero *Character.NPC) {
	fmt.Printf("HP: %v\\%v\tMP: %v\\%v\n", int(hero.CurrentHealth), hero.MaxHealth, int(hero.CurrentMana), hero.MaxMana)
	fmt.Printf("HP Regen: %v\tMP Regen: %v\n", int(hero.HealthRegen), int(hero.ManaRegen))
	damageMin := hero.DmgMultuplier * float32((hero.Weapon.MinDmg + hero.Weapon.BonusDmg))
	damageMax := hero.DmgMultuplier * float32((hero.Weapon.MaxDmg + hero.Weapon.BonusDmg))
	fmt.Printf("DMG: %v - %v\tDef:%v\n", int(damageMin), int(damageMax), int(hero.CombinedDefence()))
	fmt.Printf("Evs:%v\t\tCrit:%v\n", hero.Evasion, hero.CritChance)
}

func (game *Game) drawTraps() {
	for trapPoint, _ := range game.trapList {
		game.labyrinth.Labyrinth[trapPoint.X][trapPoint.Y] = Labyrinth.Trap
	}
}

func (game *Game) drawMonsters() {
	for _, mon := range game.monsterList {
		game.labyrinth.Labyrinth[mon.Location.X][mon.Location.Y] = mon.Symbol
	}
}

func (game *Game) drawProjectiles() {
	for _, projectile := range game.projectileList {
		game.labyrinth.Labyrinth[projectile.Location.X][projectile.Location.Y] = projectile.Symbol
	}
}

func (game *Game) drawHero() {
	game.labyrinth.Labyrinth[game.player.Base.Location.X][game.player.Base.Location.Y] = Labyrinth.CharSymbol
}

func (game *Game) clearScreen() {
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}

func (game *Game) restoreTile(x int, y int) {
	game.labyrinth.Labyrinth[x][y] = Labyrinth.Pass
	if game.start.X != game.player.Base.Location.X || game.start.Y != game.player.Base.Location.Y {
		game.labyrinth.Labyrinth[game.start.X][game.start.Y] = Labyrinth.StartPosition
	}
	if game.end.X != game.player.Base.Location.X && game.end.Y != game.player.Base.Location.Y {
		game.labyrinth.Labyrinth[game.end.X][game.end.X] = Labyrinth.ExitPosition
	}
}

func(game *Game) replaceTile(x int, y int, symbol string) {
	if game.labyrinth.Labyrinth[x][y] != Labyrinth.Wall {
		game.labyrinth.Labyrinth[x][y] = symbol
	}
}