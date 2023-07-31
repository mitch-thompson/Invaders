package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"invaders/pkg/render"
)

func main() {
	goInvaders := app.New()
	goInvadersWindow := goInvaders.NewWindow("Go Invaders")
	goInvadersWindow.Resize(fyne.NewSize(640, 480))
	gameCanvas := render.NewCanvas(640, 480)

	gameLayout := container.NewVBox()
	gameCanvas.RenderMain(goInvadersWindow, gameLayout)
	goInvadersWindow.ShowAndRun()
}
