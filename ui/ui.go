package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"observer-pattern/state"
	"observer-pattern/util"
	"strconv"
	"time"
)

type Presenter struct {
	title     string
	w         fyne.Window
	l         *widget.Label
	statement string
	algo      func(state.AppState) int
	bag       *util.Bag
}

func NewPresenter(
	application fyne.App,
	title string,
	statement string,
	algo func(state.AppState) int,
) *Presenter {
	w := application.NewWindow(title)
	w.Resize(fyne.NewSize(150, 20))

	l := widget.NewLabel("")
	w.SetContent(l)

	bag := util.NewBag()

	w.SetOnClosed(bag.Clean)

	return &Presenter{title, w, l, statement, algo, bag}
}

func (p *Presenter) Show() {
	p.w.Show()

	state.State.Observe(func(appState state.AppState) {
		result := fmt.Sprintf("%s = %d", p.statement, p.algo(appState))
		p.l.SetText(result)
		fmt.Printf("%s - State consumed for %s\n", time.Now().Format(time.DateTime), p.title)
	}).StoreIn(p.bag)
}

func BuildInput(application fyne.App) fyne.Window {
	w := application.NewWindow("Input")
	w.Resize(fyne.NewSize(150, 50))

	xLbl := widget.NewLabel("X: ")
	xInp := widget.NewEntry()
	xInp.SetPlaceHolder("0")

	yLbl := widget.NewLabel("Y: ")
	yInp := widget.NewEntry()
	yInp.SetPlaceHolder("0")

	xInp.OnChanged = func(strX string) {
		x, _ := strconv.Atoi(strX)
		state.SetX(x)
	}

	yInp.OnChanged = func(strY string) {
		y, _ := strconv.Atoi(strY)
		state.SetY(y)
	}

	xBox := container.NewHBox(xLbl, xInp)
	xBox.Resize(fyne.Size{
		Width:  150,
		Height: 25,
	})

	yBox := container.NewHBox(yLbl, yInp)
	yBox.Resize(fyne.Size{
		Width:  150,
		Height: 25,
	})

	content := container.NewVBox(xBox, yBox)

	w.SetContent(content)

	return w
}
