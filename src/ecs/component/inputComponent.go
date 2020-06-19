package component

type InputComponent struct {
	MouseX      float64
	MouseY      float64
	MouseDeltaX float64
	MouseDeltaY float64
	Stopped     bool
}

func (comp InputComponent) GetType() string {
	return "Input"
}
