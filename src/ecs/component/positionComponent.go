package component

import (
	"strconv"
)

type PositionComponent struct {
	X int
	Y int
}

func (comp PositionComponent) GetType() string {
	return "Position"
}

func (comp PositionComponent) GetShoutText() string {
	return "(" + strconv.Itoa(comp.X) + ", " + strconv.Itoa(comp.Y) + ")"
}
