//Handles the logic in the game.
package Game

//cameraMoveLeft lowers the Y coordinate of the camera's center by 1.
func (game *Game) cameraMoveLeft() {
	game.camera.Y--
}

//cameraMoveDown increases the X coordinate of the camera's center by 1.
func (game *Game) cameraMoveDown() {
	game.camera.X++
}

//cameraMoveRight increases the Y coordinate of the camera's center by 1.
func (game *Game) cameraMoveRight() {
	game.camera.Y++
}

//cameraMoveUp lowers the X coordinate of the camera's center by 1.
func (game *Game) cameraMoveUp() {
	game.camera.X--
}

//cameraReset reset the camera center.
func (game *Game) cameraReset() {
	game.camera.X = game.player.Base.Location.X
	game.camera.Y = game.player.Base.Location.Y
}
