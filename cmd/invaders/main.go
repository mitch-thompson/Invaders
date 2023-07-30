package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"invaders/pkg/render"
)

func main() {
	goInvaders := app.New()
	goInvadersWindow := goInvaders.NewWindow("Go Invaders")

	gameCanvas := render.NewCanvas()
	startButton := widget.NewButton("Start", func() {
		//todo
	})

	gameLayout := container.NewVBox(
		gameCanvas.Canvas,
		startButton,
	)

	goInvadersWindow.SetContent(gameLayout)
	goInvadersWindow.ShowAndRun()
}
