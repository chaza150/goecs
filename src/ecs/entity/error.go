package entity

import (
	"fmt"
)

type ComponentAlreadyPresentError struct {
	entityID string
	compType string
}

type ComponentNotPresentError struct {
	entityID string
	compType string
}

func (e ComponentAlreadyPresentError) Error() string {
	return fmt.Sprintf("Entity " + e.entityID + " already has component: " + e.compType)
}

func (e ComponentNotPresentError) Error() string {
	return fmt.Sprintf("Entity " + e.entityID + " does not have component: " + e.compType)
}
