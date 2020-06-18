package entity

import (
	"ecs/component"
)

type Entity struct {
	ID         string
	Components map[string]*component.Component
}

func (e *Entity) HasComponent(comp component.Component) bool {
	_, ok := e.Components[comp.GetType()]
	return ok
}

func (e *Entity) HasComponents(compList []component.Component) bool {
	for _, comp := range compList {
		if !e.HasComponent(comp) {
			return false
		}
	}
	return true
}

func (e *Entity) HasComponentByName(compName string) bool {
	_, ok := e.Components[compName]
	return ok
}

func (e *Entity) HasComponentsByName(compList []string) bool {
	for _, compName := range compList {
		if !e.HasComponentByName(compName) {
			return false
		}
	}
	return true
}

func (e *Entity) AddComponent(comp component.Component) error {
	if len(e.Components) == 0 {
		e.Components = make(map[string]*component.Component)
	}
	if !e.HasComponent(comp) {
		e.Components[comp.GetType()] = &comp
		return nil
	} else {
		return ComponentAlreadyPresentError{e.ID, comp.GetType()}
	}
}

func (e *Entity) AddComponents(compList []component.Component) error {
	for _, comp := range compList {
		if e.HasComponent(comp) {
			return ComponentAlreadyPresentError{e.ID, comp.GetType()}
		}
	}
	for _, comp := range compList {
		e.AddComponent(comp)
	}
	return nil
}

func (e *Entity) RemoveComponent(comp component.Component) error {
	if e.HasComponent(comp) {
		delete(e.Components, comp.GetType())
		return nil
	} else {
		return ComponentNotPresentError{e.ID, comp.GetType()}
	}
}

func (e *Entity) RemoveComponents(compList []component.Component) error {
	if e.HasComponents(compList) {
		for _, comp := range compList {
			e.RemoveComponent(comp)
		}
		return nil
	} else {
		for _, comp := range compList {
			if !e.HasComponent(comp) {
				return ComponentNotPresentError{e.ID, comp.GetType()}
			}
		}
	}
	return nil
}

func (e *Entity) RemoveComponentByName(compName string) error {
	if e.HasComponentByName(compName) {
		delete(e.Components, compName)
		return nil
	} else {
		return ComponentNotPresentError{e.ID, compName}
	}
}

func (e *Entity) RemoveComponentsByName(compList []string) error {
	if e.HasComponentsByName(compList) {
		for _, compName := range compList {
			e.RemoveComponentByName(compName)
		}
		return nil
	} else {
		for _, compName := range compList {
			if !e.HasComponentByName(compName) {
				return ComponentNotPresentError{e.ID, compName}
			}
		}
	}
	return nil
}

func (e *Entity) GetComponent(compName string) (*component.Component, error) {
	if e.HasComponentByName(compName) {
		return e.Components[compName], nil
	} else {
		return nil, ComponentNotPresentError{e.ID, compName}
	}
}

func NewEntity(ID string) *Entity {
	return &Entity{ID: ID}
}

func (e *Entity) With(comp component.Component) *Entity {
	e.AddComponent(comp)
	return e
}
