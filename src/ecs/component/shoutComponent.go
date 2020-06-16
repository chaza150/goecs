package component

type ShoutComponent struct {
	ShoutText string
}

func (comp ShoutComponent) GetType() string {
	return "Shout"
}

func (comp ShoutComponent) GetShoutText() string {
	return comp.ShoutText
}
