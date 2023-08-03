package render

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const (
	SPRITE_SIZE = float32(26) //todo move to configuration file
	BLOCK_SIZE  = float32(35)
	MARGIN_SIZE = float32(5)
)

type FyneRenderer struct {
	window               fyne.Window
	menuContainer        *fyne.Container
	gameObjectsContainer *fyne.Container
}

func NewFyneRenderer(window fyne.Window, menuContainer, objectContainer *fyne.Container) *FyneRenderer {
	return &FyneRenderer{
		window:               window,
		menuContainer:        menuContainer,
		gameObjectsContainer: objectContainer,
	}
}

func (r *FyneRenderer) startButtonPressed() {
	fmt.Println("Start Button pressed")
	//r.ClearScreen()
	r.window.SetContent(r.gameObjectsContainer)
	width := r.gameObjectsContainer.Size().Width
	height := r.gameObjectsContainer.Size().Height
	position := NewPosition(width/2-SPRITE_SIZE/2, height-40)
	r.drawBackButton()
	r.DrawDefender(*position)
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			isBottom := false
			if i == 4 {
				isBottom = true
			}
			var y float32 = float32(i)*BLOCK_SIZE + MARGIN_SIZE
			var x float32 = float32(j)*BLOCK_SIZE + MARGIN_SIZE
			position = NewPosition(x, y)
			r.DrawAlien(*position, isBottom)
		}
	}
	r.gameObjectsContainer.Refresh()
}

func (r *FyneRenderer) backButtonPressed() {
	fmt.Println("Back button pressed")
	r.ClearGame()
	r.ClearMenu()
	r.DrawMenu()
	r.window.SetContent(r.menuContainer)
	r.menuContainer.Refresh()
}

func (r *FyneRenderer) scoresButtonPressed() {
	fmt.Println("Scores button pressed")
	r.ClearScreen()
}

func (r *FyneRenderer) DrawMenu() {
	startButton := widget.NewButton("Start", r.startButtonPressed)
	scoresButton := widget.NewButton("Scores", r.scoresButtonPressed)

	buttonContainer := container.NewHBox(startButton, scoresButton)
	buttonContainer.Layout = layout.NewHBoxLayout()

	r.menuContainer.Add(buttonContainer)
}

func (r *FyneRenderer) DrawDefender(position Position) {
	defender := NewDefender(position)
	defenderImage := canvas.NewImageFromFile(defender.path)
	//if defenderImage == nil || defenderImage.Resource == nil {
	//	fmt.Println("Failed to load image from file:", defender.path)
	//	return
	//}
	defenderImage.FillMode = canvas.ImageFillOriginal
	defenderImage.Resize(fyne.NewSize(SPRITE_SIZE, SPRITE_SIZE))

	defenderImage.Move(fyne.NewPos(defender.position.x, defender.position.y))

	r.gameObjectsContainer.Add(defenderImage)
}

func (r *FyneRenderer) DrawAlien(position Position, isBottom bool) {
	alien := NewAlien(position, isBottom)
	alienImage := canvas.NewImageFromFile(alien.path)
	//if alienImage == nil || alienImage.Resource == nil {
	//	fmt.Println("Failed to load image from file:", alien.path)
	//	return
	//}
	alienImage.FillMode = canvas.ImageFillOriginal
	alienImage.Resize(fyne.NewSize(SPRITE_SIZE, SPRITE_SIZE))

	alienImage.Move(fyne.NewPos(alien.position.x, alien.position.y))

	r.gameObjectsContainer.Add(alienImage)
}

func (r *FyneRenderer) drawBackButton() {
	backButton := widget.NewButton("<-", r.backButtonPressed)
	backButton.Resize(fyne.NewSize(25, 25))
	backButton.Move(fyne.NewPos(0, 0))
	r.gameObjectsContainer.Add(backButton)
}

func (r *FyneRenderer) DrawScore(score int) {
}

func (r *FyneRenderer) ClearMenu() {
	for len(r.menuContainer.Objects) > 0 {
		r.menuContainer.Remove(r.menuContainer.Objects[0])
	}
	r.menuContainer.Refresh()
}

func (r *FyneRenderer) ClearGame() {
	for len(r.gameObjectsContainer.Objects) > 0 {
		r.gameObjectsContainer.Remove(r.gameObjectsContainer.Objects[0])
	}
	r.gameObjectsContainer.Refresh()
}

func (r *FyneRenderer) ClearScoreScreen() {
	//todo
}

func (r *FyneRenderer) ClearScreen() {
	r.ClearMenu()
	r.ClearGame()
	r.ClearScoreScreen()
}
