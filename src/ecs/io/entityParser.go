package io

import (
	"regexp"
)

const (
	EntityRegex = `\s*\"\w+\"\s*(:|\?)\s*\"\w+\"\s*{((` + ComponentRegex + `,)*(` + ComponentRegex + `)+|\s*)\s*}`
)

func ParseEntity(text string) (*EntityNode, string) {
	if isValidEntity(text) {
		nameRegex, _ := regexp.Compile(`\"\w+\"`)
		searchName := ParseQuotedString(nameRegex.FindString(text))
		parentName := ParseQuotedString(nameRegex.FindAllString(text, 2)[1])
		abstract := false
		if matched, _ := regexp.MatchString(`\s*\"\w+\"\s*\?\s*\"\w+\"\s*{`, text); matched {
			abstract = true
		}
		componentRegex, _ := regexp.Compile(ComponentRegex)
		componentMatches := componentRegex.FindAllString(text, -1)
		var components []*ComponentNode

		for _, componentMatch := range componentMatches {
			components = append(components, ParseComponent(componentMatch))
		}
		var children []*EntityNode
		entityNode := EntityNode{SearchName: searchName, IsAbstract: abstract, Parent: nil, Children: children, Components: components}
		return &entityNode, parentName
	} else {
		panic("Tried to Parse Invalid Entity: " + text)
	}
}

func isValidEntity(text string) bool {
	if matched, _ := regexp.MatchString(EntityRegex, text); matched {
		braceCount := 0
		for _, char := range text {
			if string(char) == "{" {
				braceCount++
			} else if string(char) == "}" {
				braceCount--
			}
		}
		if braceCount == 0 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
