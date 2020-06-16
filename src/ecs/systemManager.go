package ecs

import (
	"ecs/system"
	"fmt"
)

type SystemManager struct {
	Systems map[string]*system.System
}

func NewSystemManager() SystemManager {
	return SystemManager{make(map[string]*system.System)}
}

func (sm *SystemManager) AddSystem(sys system.System) {
	if len(sm.Systems) == 0 {
		sm.Systems = make(map[string]*system.System)
	}
	if !sm.HasSystem(sys.GetName()) {
		sm.Systems[sys.GetName()] = &sys
		fmt.Println("Added System: " + sys.GetName())
	}
}

func (sm *SystemManager) HasSystem(systemName string) bool {
	_, ok := sm.Systems[systemName]
	return ok
}
