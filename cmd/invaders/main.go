package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"invaders/pkg/render"
)

func main() {
	goInvaders := app.New()
	goInvadersWindow := goInvaders.NewWindow("Go Invaders")
	goInvadersWindow.Resize(fyne.NewSize(640, 480))
	menuContainer := container.New(layout.NewVBoxLayout())
	gameObjectsContainer := container.New(render.NewNullLayout())
	goInvadersWindow.SetContent(menuContainer)
	renderer := render.NewFyneRenderer(goInvadersWindow, menuContainer, gameObjectsContainer)
	renderer.DrawMenu()
	goInvadersWindow.ShowAndRun()
}
