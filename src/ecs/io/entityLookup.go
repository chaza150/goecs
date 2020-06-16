package io

import (
	"fmt"
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
	ComponentValues map[string]*ValueNode
}

type ValueNode struct {
	IsComponent    bool
	ComponentValue *ComponentNode
	BasicValue     string
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

func (lookup *EntityLookup) AddEntityNode(node *EntityNode) {
	if !lookup.HasEntityName(node.SearchName) {
		node.Parent.Children = append(node.Parent.Children, node)
		if !node.IsAbstract {
			lookup.QuickConcreteMap[node.SearchName] = node
		}
		lookup.QuickMap[node.SearchName] = node
	} else {
		fmt.Println("Err: Cannot have multiple " + node.SearchName + " entity types in lookup")
	}
}

func (lookup *EntityLookup) CreateEntityNode(searchName string, abstract bool, parentName string, components []*ComponentNode) {
	var children []*EntityNode
	lookup.AddEntityNode(EntityNode{searchName, abstract, lookup.GetEntityNode(parentName), children, components})
}

func (lookup *EntityLookup) HasEntityName(entityName string) bool {
	_, ok := lookup.QuickMap[entityName]
	return ok
}

func (lookup *EntityLookup) GetEntityNode(searchName string) *EntityNode {
	if lookup.HasEntityName(searchName) {
		return lookup.QuickMap[searchName]
	} else {
		fmt.Println("Err: Could not retrieve Entity type: " + searchName)
		return nil
	}
}

//Complete Parser for Data file to components, values, etc
//Complete Interpreter for creating instances of
