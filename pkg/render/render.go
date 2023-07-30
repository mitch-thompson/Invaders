// invaders/pkg/render/render.go

package render

import (
	"fyne.io/fyne/v2"
)

// Canvas represents the game canvas that encapsulates the Fyne canvas and related functionalities.
type Canvas struct {
	Canvas fyne.CanvasObject
}

// NewCanvas creates a new instance of the game canvas.
func NewCanvas() *Canvas {
	gameCanvas := fyne.NewContainer()

	return &Canvas{
		Canvas: gameCanvas,
	}
}

// RenderGame is a method of the Canvas type used to render the game graphics.
// todo Implement game rendering logic here.
func (c *Canvas) RenderGame() {

}

// HandleInput is a method of the Canvas type used to handle user input events.
// todo Implement game input handling logic here.
func (c *Canvas) HandleInput(e *fyne.KeyEvent) {
}
