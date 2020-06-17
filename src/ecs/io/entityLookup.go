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
	entity := EntityNode{searchName, abstract, lookup.GetEntityNode(parentName), children, components}
	lookup.AddEntityNode(&entity)
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

func CreateComponentNode(componentType string, values map[string]string) *ComponentNode {
	component := ComponentNode{componentType, values}
	return &component
}

func (lookup *EntityLookup) CreateFullEntityNode(searchName string, abstract bool, parentName string, components []*ComponentNode) {
	lookup.CreateEntityNode(searchName, abstract, parentName, components)
}

func (lookup *EntityLookup) ParseEntityData(text string) {
	entityRegex, _ := regexp.Compile(EntityRegex)
	entityMatches := entityRegex.FindAllString(text, -1)

	for _, entityMatch := range entityMatches {
		entityNode, parentName := ParseEntity(entityMatch)
		entityNode.Parent = lookup.GetEntityNode(parentName)
		lookup.AddEntityNode(entityNode)
	}
}

func (lookup *EntityLookup) PrintEntityTree() {
	lookup.PrintEntity(lookup.Root, 0)
}

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
