package component

import (
	"github.com/faiface/pixel/pixelgl"
)

type InputComponent struct {
	Window      *pixelgl.Window
	MouseDeltaX int
	MouseDeltaY int
	Stopped     bool
}

func (comp InputComponent) GetType() string {
	return "Input"
}
