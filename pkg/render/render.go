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

// calculateButtonSize sets a button to 10 pixels or 10% of screen size
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

// RenderMain draws the main screen
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

// gameSetup draws the canvas for the game to start
func (c *Canvas) gameSetup(board *canvas.Rectangle) {
	gameboard := NewEmptyGameBoard()

	width := board.Size().Width
	height := board.Size().Height

	player, err := NewSprite(DEFENDER)
	if err != nil {
		//todo
	}
	image := canvas.NewImageFromFile(player.path)
	image.FillMode = canvas.ImageFillOriginal
	image.Resize(fyne.NewSize(26, 26))
	gameObject := NewGameObject(player, width/2, height-30, false, false)
	image.Move(fyne.NewPos(gameObject.x-13, gameObject.y))
	gameObject.canvasObject = image
	gameboard.gameObjects = append(gameboard.gameObjects, gameObject)

	c.Canvas.(*fyne.Container).Add(image)

	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			var y float32 = float32(i)*40 + 5
			var x float32 = float32(j)*40 + 5
			alien, err := NewSprite(ALIEN)
			if err != nil {
				//todo
			}
			image = canvas.NewImageFromFile(alien.path)
			image.FillMode = canvas.ImageFillOriginal
			image.Resize(fyne.NewSize(26, 26))
			gameObject = NewGameObject(alien, x, y, true, false)
			image.Move(fyne.NewPos(gameObject.x, gameObject.y))
			gameObject.canvasObject = image
			gameboard.gameObjects = append(gameboard.gameObjects, gameObject)
			c.Canvas.(*fyne.Container).Add(image)
		}
	}
}

// moveAliensHorizontal moves aliens left or right
func (c *Canvas) moveAliensHorizontal(g *GameBoard, direction string) {
	var distance float32
	if direction == "left" {
		distance = -40
	} else {
		distance = 40
	}
	for _, o := range g.gameObjects {
		if o.isAlien {
			o.x = o.x + distance
			pos := fyne.NewPos(o.x, o.y)
			o.canvasObject.Move(pos)
		}
	}
	c.Canvas.Refresh()
}

// moveAliensDown moves the aliens down one row
func (c *Canvas) moveAliensDown(g *GameBoard) {
	for _, o := range g.gameObjects {
		if o.isAlien {
			o.y = o.y + 40
			pos := fyne.NewPos(o.x, o.y)
			o.canvasObject.Move(pos)
		}
	}
	c.Canvas.Refresh()
}

// RenderGame is a method of the Canvas type used to render the game graphics.
// todo Implement game rendering logic here.
func (c *Canvas) RenderGame(gameLayout *fyne.Container, window fyne.Window) {
	gameLayout.Objects = nil

	blankRect := canvas.NewRectangle(color.Black)
	//original game size is 224x256
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
	c.gameSetup(blankRect)
}

// HandleInput is a method of the Canvas type used to handle user input events.
// todo Implement game input handling logic here.
func (c *Canvas) HandleInput(e *fyne.KeyEvent) {
}
