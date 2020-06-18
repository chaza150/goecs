package io

import (
	"ecs/component"
	"ecs/entity"
	"fmt"
)

var componentNecessaryValues = map[string][]string{
	"Shout":    []string{"ShoutText"},
	"Position": []string{"X", "Y"},
}

func (lookup *EntityLookup) InstantiateEntity(searchName string, entityID string) *entity.Entity {
	entityNode := lookup.GetConcreteEntityNode(searchName)
	if entityNode != nil {

		entity := entity.NewEntity(entityID)
		AddComponentsFromEntityNode(entityNode, entity)

		return entity

	} else {
		fmt.Println("Unable to instantiate entity: " + entityID + " (Type: " + searchName + ")")
		return nil
	}
}

func AddComponentsFromEntityNode(node *EntityNode, entity *entity.Entity) {
	for _, compNode := range node.Components {
		AddComponentToEntityFromNode(compNode.ComponentType, compNode.ComponentValues, entity)
	}
	if node.Parent != nil {
		fmt.Println("adding components from parent: " + node.Parent.SearchName)
		AddComponentsFromEntityNode(node.Parent, entity)
	}
}

//Adds a Specific (componentType) Component To an Entity with values (values)
func AddComponentToEntityFromNode(componentType string, values map[string]string, entity *entity.Entity) {
	//Only add component if all necessary values are present
	if ComponentHasNecessaryValues(componentType, values) {
		//Based on the string componentType, add the respective component
		var err error = nil

		switch componentType {
		case "Shout":
			err = entity.AddComponent(component.ShoutComponent{ShoutText: ParseQuotedString(values["ShoutText"])})
		case "Position":
			err = entity.AddComponent(component.PositionComponent{X: ParseInt(values["X"]), Y: ParseInt(values["Y"])})
		default:
			fmt.Println("Entity, " + entity.ID + "'s \"" + componentType + "\" component, does not have parsing rules")
		}

		if err != nil {
			fmt.Println("Could not add component, \"" + componentType + "\" to entity: " + entity.ID)
		}

	} else {
		fmt.Println("Entity, " + entity.ID + "'s \"" + componentType + "\" component, does not have necessary values.")
	}
}

func ComponentHasNecessaryValues(componentType string, values map[string]string) bool {
	necessaryValues, ok := componentNecessaryValues[componentType]
	if ok {
		for _, valueName := range necessaryValues {
			_, contains := values[valueName]
			if !contains {
				return false
			}
		}
		return true
	} else {
		fmt.Println("Component Type, \"" + componentType + "\"'s necessary values are unknown")
		return false
	}
}