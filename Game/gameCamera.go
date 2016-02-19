package Game

func (game *Game) cameraMoveLeft() {
	game.camera.Y--
}

func (game *Game) cameraMoveDown() {
	game.camera.X++
}

func (game *Game) cameraMoveRight() {
	game.camera.Y++
}

func (game *Game) cameraMoveUp() {
	game.camera.X--
}

func (game *Game) cameraReset() {
	game.camera.X = game.player.Base.Location.X
	game.camera.Y = game.player.Base.Location.Y
}
