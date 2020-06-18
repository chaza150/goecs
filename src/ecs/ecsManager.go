package ecs

import (
	"ecs/entity"
	"ecs/io"
	"fmt"
	"github.com/faiface/pixel/pixelgl"
)

type ECS struct {
	EntManager   *EntityManager
	SysManager   *SystemManager
	EntityLookup *io.EntityLookup
}

func NewECS(window *pixelgl.Window) ECS {
	em := NewEntityManager()
	sm := NewSystemManager()
	lookup := io.NewEntityLookup(window)
	return ECS{&em, &sm, lookup}
}

func (ecs *ECS) UpdateSystems() {
	for _, sys := range ecs.SysManager.Systems {
		fmt.Println("Updating " + (*sys).GetName())
		(*sys).Update(ecs.EntManager.GetEntitiesAsSlice())
	}
}

func (ecs *ECS) LoadEntityTypeData(fileName string) {
	ecs.EntityLookup.ParseEntityData(io.LoadFile(fileName))
}

func (ecs *ECS) InstantiateEntity(entitySearchName string, entityID string) *entity.Entity {
	entity := ecs.EntityLookup.InstantiateEntity(entitySearchName, entityID)
	if entity != nil {
		ecs.EntManager.AddEntity(entity)
	}
	return entity
}
