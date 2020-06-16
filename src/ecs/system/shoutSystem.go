package system

import (
	"ecs/component"
	"ecs/entity"
	"fmt"
)

type ShoutSystem struct {
}

func (sys ShoutSystem) Update(entities []*entity.Entity) {
	for _, ent := range entities {
		//fmt.Println(ent.ID)
		if ent.HasComponentByName("Shout") {
			shoutComponent, _ := ent.GetComponent("Shout")
			fmt.Println((*shoutComponent).(component.ShoutComponent).GetShoutText())
		}
	}
}

func (sys ShoutSystem) GoUpdate(entities []*entity.Entity) error {
	for _, ent := range entities {
		if ent.HasComponentByName("Shout") {
			shoutComponent, _ := ent.GetComponent("Shout")
			go fmt.Println((*shoutComponent).(component.ShoutComponent).GetShoutText())
		}
	}
	return nil
}

func (sys ShoutSystem) IsGoable() bool {
	return true
}

func (sys ShoutSystem) GetRequiredComponents() []string {
	return []string{"Shout"}
}

func (sys ShoutSystem) GetName() string {
	return "ShoutSys"
}
