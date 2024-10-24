package main

import (
	"fyne.io/fyne/v2/app"
	"observer-pattern/state"
	"observer-pattern/ui"
)

func main() {
	a := app.New()

	inpW := ui.BuildInput(a)
	inpW.Show()

	sumW := ui.NewPresenter(a, "Sum", "X + Y", func(appState state.AppState) int {
		return appState.X + appState.Y
	})

	subW := ui.NewPresenter(a, "Subtract", "X - Y", func(appState state.AppState) int {
		return appState.X - appState.Y
	})

	mulW := ui.NewPresenter(a, "Multiply", "X * Y", func(appState state.AppState) int {
		return appState.X * appState.Y
	})

	sumW.Show()
	subW.Show()
	mulW.Show()

	a.Run()
}
