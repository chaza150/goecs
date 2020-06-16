package entity

import (
	"ecs/component"
)

func NewPlayerEntity() *Entity {
	newPlayer := NewEntity("Player")
	newPlayer.AddComponent(component.ShoutComponent{"Player"})
	return newPlayer
}
