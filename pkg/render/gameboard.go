package render

import "fyne.io/fyne/v2"

// GameBoard represents a container for all objects in the game
type GameBoard struct {
	gameObjects []*GameObject
}

// GameObject represents an object on the canvas
type GameObject struct {
	object   Sprite
	x, y     float32
	isAlien  bool
	isBottom bool
	//todo review this should to determine if it should be removed
	canvasObject fyne.CanvasObject
}

// NewEmptyGameBoard creates a new empty instance of GameBoard
func NewEmptyGameBoard() *GameBoard {
	return &GameBoard{[]*GameObject{}}
}

// NewGameObject creates a new instance of GameObject
func NewGameObject(sprite *Sprite, x, y float32, isAlien, isBottom bool) *GameObject {
	return &GameObject{
		*sprite,
		x,
		y,
		isAlien,
		isBottom,
		fyne.CanvasObject(nil),
	}
}
