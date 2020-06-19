package system

import (
	"ecs/component"
	"ecs/entity"
	"errors"
	"fmt"
	"github.com/faiface/pixel/pixelgl"
)

type InputSystem struct {
	Window *pixelgl.Window
}

func (sys InputSystem) Update(entities []*entity.Entity) {
	for _, ent := range entities {
		//fmt.Println(ent.ID)
		if ent.HasComponentByName("Input") {

			X := sys.Window.MousePosition().X
			Y := sys.Window.MousePosition().Y
			deltaX := X - sys.Window.MousePreviousPosition().X
			deltaY := Y - sys.Window.MousePreviousPosition().Y

			stopped := sys.Window.Pressed(pixelgl.KeyR)

			ent.ReplaceComponent("Input", component.InputComponent{MouseDeltaX: deltaX, MouseDeltaY: deltaY, Stopped: stopped})

			fmt.Println("X: ", deltaX, ", Y: ", deltaY)

			if stopped {
				fmt.Println("Stopping")
				sys.Window.SetCursorVisible(true)
				sys.Window.SetClosed(true)
			}
		}
	}
}

func (sys InputSystem) GoUpdate(entities []*entity.Entity) error {
	return errors.New("Unable to Update in multithreaded mode!")
}

func (sys InputSystem) IsGoable() bool {
	return false
}

func (sys InputSystem) GetRequiredComponents() []string {
	return []string{"Input"}
}

func (sys InputSystem) GetName() string {
	return "InputSys"
}
