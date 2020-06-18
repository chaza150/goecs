package io

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type EntityLookup struct {
	Root             *EntityNode
	QuickMap         map[string]*EntityNode
	QuickConcreteMap map[string]*EntityNode
}

type EntityNode struct {
	SearchName string
	IsAbstract bool
	Parent     *EntityNode
	Children   []*EntityNode
	Components []*ComponentNode
}

type ComponentNode struct {
	ComponentType   string
	ComponentValues map[string]string
}

func NewEntityLookup() *EntityLookup {
	var children []*EntityNode
	var quickMap map[string]*EntityNode
	quickMap = make(map[string]*EntityNode)
	root := EntityNode{SearchName: "base", IsAbstract: true, Parent: nil, Children: children}
	quickMap[root.SearchName] = &root
	entityLookup := EntityLookup{Root: &root, QuickMap: quickMap, QuickConcreteMap: make(map[string]*EntityNode)}
	return &entityLookup
}

//Add an Entity Node to the Entity Lookup Tables
func (lookup *EntityLookup) AddEntityNode(node *EntityNode) {
	if !lookup.HasCyclicParenthood(node.Parent.SearchName, node.SearchName) {
		if !lookup.HasEntityName(node.SearchName) {
			node.Parent.Children = append(node.Parent.Children, node)
			if !node.IsAbstract {
				lookup.QuickConcreteMap[node.SearchName] = node
			}
			lookup.QuickMap[node.SearchName] = node
		} else {
			fmt.Println("Err: Cannot have multiple " + node.SearchName + " entity types in lookup")
		}
	} else {
		panic("Cyclic Parenthood when Adding Entity: " + node.SearchName)
	}
}

//Create and Add Entity Node
func (lookup *EntityLookup) CreateEntityNode(searchName string, abstract bool, parentName string, components []*ComponentNode) {
	var children []*EntityNode
	entity := EntityNode{searchName, abstract, lookup.GetEntityNode(parentName), children, components}
	lookup.AddEntityNode(&entity)
}

//Check That Entity Node with Name Exists
func (lookup *EntityLookup) HasEntityName(entityName string) bool {
	_, ok := lookup.QuickMap[entityName]
	return ok
}

//Check That Non-Abstract Entity Node with Name Exists
func (lookup *EntityLookup) HasConcreteEntityName(entityName string) bool {
	_, ok := lookup.QuickConcreteMap[entityName]
	return ok
}

//Retrieve an Entity Node from SearchName
func (lookup *EntityLookup) GetEntityNode(searchName string) *EntityNode {
	if lookup.HasEntityName(searchName) {
		return lookup.QuickMap[searchName]
	} else {
		fmt.Println("Err: Could not retrieve Entity type: " + searchName + " - Node does not exist!")
		return nil
	}
}

//Retrieve an Entity Node from SearchName
func (lookup *EntityLookup) GetConcreteEntityNode(searchName string) *EntityNode {
	if lookup.HasConcreteEntityName(searchName) {
		return lookup.QuickMap[searchName]
	} else {
		if lookup.HasEntityName(searchName) {
			fmt.Println("Err: Could not retrieve Entity type: " + searchName + " - Abstract Nodes cannot be instantiated!")
		} else {
			fmt.Println("Err: Could not retrieve Entity type: " + searchName + " - Does not exist!")
		}
		return nil
	}
}

//Create a Component Node from map of Values
func (lookup *EntityLookup) CreateComponentNode(componentType string, values map[string]string) *ComponentNode {
	component := ComponentNode{componentType, values}
	return &component
}

//Parse Entity Types From Text and Add to Nodes
func (lookup *EntityLookup) ParseEntityData(text string) {
	entityRegex, _ := regexp.Compile(EntityRegex)
	entityMatches := entityRegex.FindAllString(text, -1)

	for _, entityMatch := range entityMatches {
		entityNode, parentName := ParseEntity(entityMatch)
		entityNode.Parent = lookup.GetEntityNode(parentName)
		lookup.AddEntityNode(entityNode)

	}
}

// Check for Cycles in Parenthood
func (lookup *EntityLookup) HasCyclicParenthood(parentName, entityName string) bool {
	if parentName == entityName {
		return true
	} else {
		parent := lookup.GetEntityNode(parentName).Parent
		if parent == nil {
			return false
		} else {
			return lookup.HasCyclicParenthood(parent.SearchName, entityName)
		}
	}
}

//Print Full Entity Type Tree
func (lookup *EntityLookup) PrintEntityTree() {
	lookup.PrintEntity(lookup.Root, 0)
}

//Recursively Print One Entity At a Time with Children
func (lookup *EntityLookup) PrintEntity(entityNode *EntityNode, indent int) {
	indentString := strings.Repeat("    ", indent)
	fmt.Println(indentString + "Name: " + entityNode.SearchName)
	fmt.Println(indentString + " - Abstract: " + strconv.FormatBool(entityNode.IsAbstract))
	fmt.Println(indentString + " - Components:")
	for _, componentNode := range entityNode.Components {
		fmt.Println(indentString + "   - " + componentNode.ComponentType)
		for field, value := range componentNode.ComponentValues {
			fmt.Println(indentString + "     - " + field + " : " + value)
		}
	}
	fmt.Println(indentString + " - Children:")
	for _, childNode := range entityNode.Children {
		lookup.PrintEntity(childNode, indent+1)
	}
}

//Complete Parser for Data file to components, values, etc
//Complete Interpreter for creating instances of
