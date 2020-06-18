package system

import (
	"ecs/component"
	"ecs/entity"
	"errors"
	"fmt"
	"github.com/faiface/pixel/pixelgl"
)

type InputSystem struct{}

func (sys InputSystem) Update(entities []*entity.Entity) {
	for _, ent := range entities {
		//fmt.Println(ent.ID)
		if ent.HasComponentByName("Input") {
			inputComponent, _ := ent.GetProperComponent("Input")
			inputComp := (inputComponent).(component.InputComponent)

			win := inputComp.Window

			ent.ReplaceComponent("Input", component.InputComponent{Window: win, MouseDeltaX: int(win.MousePosition().X - win.MousePreviousPosition().X), MouseDeltaY: int(win.MousePosition().Y - win.MousePreviousPosition().Y), Stopped: win.Pressed(pixelgl.KeyR)})

			fmt.Println("X: ", inputComp.MouseDeltaX, ", Y: ", inputComp.MouseDeltaY)

			inComp, _ := ent.GetProperComponent("Input")
			if inComp.(component.InputComponent).Stopped {
				fmt.Println("Stopping")
				win.SetCursorVisible(true)
				win.SetClosed(true)
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
