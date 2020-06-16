package system

import "ecs/entity"

type System interface {
	Update(entities []*entity.Entity)
	IsGoable() bool
	GoUpdate(entities []*entity.Entity) error
	GetRequiredComponents() []string
	GetName() string
}
