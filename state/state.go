package state

import "observer-pattern/util"

var (
	State util.Subject[AppState]
)

func init() {
	State = util.NewConcreteSubject(AppState{})
}

type AppState struct {
	X int
	Y int
}

func SetX(x int) {
	if State.GetState().X == x {
		return
	}

	State.SetState(AppState{
		X: x,
		Y: State.GetState().Y,
	})
}

func SetY(y int) {
	if State.GetState().Y == y {
		return
	}

	State.SetState(AppState{
		X: State.GetState().X,
		Y: y,
	})
}
