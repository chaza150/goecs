package system

import "ecs/entity"

type System interface {
	Update(entities map[string]*entity.Entity)
	IsGoable() bool
	GoUpdate(entities map[string]*entity.Entity) error
	GetRequiredComponents() []string
	GetName() string
}
