// invaders/pkg/render/render.go

package render

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

// Canvas represents the game canvas that encapsulates the Fyne canvas and related functionalities.
type Canvas struct {
	Canvas fyne.CanvasObject
}

// NewCanvas creates a new instance of the game canvas.
func NewCanvas(width, height float32) *Canvas {
	gameCanvas := fyne.NewContainer()
	gameCanvas.Resize(fyne.NewSize(width, height))

	return &Canvas{
		Canvas: gameCanvas,
	}
}

//todo
func (c *Canvas) RenderScore() {
}

func calculateButtonSize(windowSize fyne.Size) fyne.Size {
	buttonWidth := 0.1 * windowSize.Width
	buttonHeight := 0.1 * windowSize.Height
	if buttonWidth > 10 {
		buttonWidth = 10
	}
	if buttonHeight > 10 {
		buttonHeight = 10
	}

	return fyne.NewSize(buttonWidth, buttonHeight)
}

func (c *Canvas) RenderMain(window fyne.Window, gameLayout *fyne.Container) {
	startButton := widget.NewButton("Start", func() {
		c.RenderGame(gameLayout, window)
	})

	scoresButton := widget.NewButton("Scores", func() {
		// TODO: Handle scores button click
	})

	startButton.OnTapped = func() {
		c.RenderGame(gameLayout, window)
	}

	buttonsContainer := container.NewHBox(startButton, scoresButton)
	buttonsContainer.Layout = layout.NewHBoxLayout()

	gameLayout.Objects = []fyne.CanvasObject{
		c.Canvas,
		buttonsContainer,
	}

	window.SetContent(gameLayout)
}

// RenderGame is a method of the Canvas type used to render the game graphics.
// todo Implement game rendering logic here.
func (c *Canvas) RenderGame(gameLayout *fyne.Container, window fyne.Window) {
	gameLayout.Objects = nil

	blankRect := canvas.NewRectangle(color.Black)
	blankRect.Resize(fyne.NewSize(640, 480))

	backButton := widget.NewButton("<-", func() {
		c.RenderMain(window, gameLayout)
	})
	backButton.Resize(calculateButtonSize(window.Canvas().Size()))

	buttonsContainer := container.NewHBox(backButton)
	buttonsContainer.Layout = layout.NewHBoxLayout()

	gameLayout.Objects = []fyne.CanvasObject{
		c.Canvas,
		buttonsContainer,
	}

	c.Canvas.(*fyne.Container).Objects = []fyne.CanvasObject{blankRect}
	c.Canvas.Refresh()

	gameLayout.Refresh()
}

// HandleInput is a method of the Canvas type used to handle user input events.
// todo Implement game input handling logic here.
func (c *Canvas) HandleInput(e *fyne.KeyEvent) {
}
