package ecs

import "fmt"

type ECS struct {
	EntManager *EntityManager
	SysManager *SystemManager
}

func NewECS() ECS {
	em := NewEntityManager()
	sm := NewSystemManager()
	return ECS{&em, &sm}
}

func (ecs ECS) UpdateSystems() {
	for _, sys := range ecs.SysManager.Systems {
		fmt.Println("Updating " + (*sys).GetName())
		(*sys).Update(ecs.EntManager.GetEntitiesAsSlice())
	}
}
