package ecs

import (
	"ecs/entity"
	"fmt"
)

type EntityManager struct {
	Entities map[string]*entity.Entity
}

func NewEntityManager() EntityManager {
	return EntityManager{make(map[string]*entity.Entity)}
}

func (em *EntityManager) AddEntity(ent *entity.Entity) {
	if len(em.Entities) == 0 {
		em.Entities = make(map[string]*entity.Entity)
	}
	em.Entities[ent.ID] = ent
	fmt.Println("Added Entity: " + ent.ID)
}

func (em *EntityManager) AddNewEntity(ID string) *entity.Entity {
	ent := entity.NewEntity(ID)
	fmt.Println("Created New Entity: " + ID)
	em.AddEntity(ent)
	return ent
}

func (em *EntityManager) GetEntity(ID string) *entity.Entity {
	return em.Entities[ID]
}

func (em *EntityManager) RemoveEntity(ID string) error {
	if em.HasEntity(ID) {
		delete(em.Entities, ID)
		return nil
	} else {
		return NoSuchEntityError{ID}
	}
}

func (em *EntityManager) HasEntity(ID string) bool {
	_, ok := em.Entities[ID]
	return ok
}

func (em *EntityManager) GetEntitiesAsSlice() []*entity.Entity {
	entities := make([]*entity.Entity, 0)
	for _, ent := range em.Entities {
		entities = append(entities, ent)
	}
	return entities
}
