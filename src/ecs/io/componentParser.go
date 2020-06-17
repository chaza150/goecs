package io

import (
	"regexp"
	"strings"
)

const (
	ComponentRegex = `\s*\"\w+\"\s*:\s*{\s*(\"\w+\"\s*:\s*` + ValueRegex + `\s*,\s*)*(\"\w+\"\s*:\s*` + ValueRegex + `)\s*}`
)

func ParseComponent(text string) *ComponentNode {
	if isValidComponent(text) {
		nameRegex, _ := regexp.Compile(`\"\w+\"`)
		componentType := ParseQuotedString(nameRegex.FindString(text))

		fieldRegex, _ := regexp.Compile(`\"\w+\"\s*:\s*` + ValueRegex)
		componentFields := fieldRegex.FindAllString(text, -1)
		var compValues map[string]string = make(map[string]string)
		for _, field := range componentFields {
			items := strings.Split(field, ":")
			fieldName := ParseQuotedString(strings.TrimSpace(items[0]))
			fieldValue := strings.TrimSpace(items[1])
			compValues[fieldName] = fieldValue
		}
		component := ComponentNode{ComponentType: componentType, ComponentValues: compValues}
		return &component
	} else {
		panic("Tried to Parse Invalid Component: " + text)
	}
}

func isValidComponent(text string) bool {
	if matched, _ := regexp.MatchString(ComponentRegex, text); matched {
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

//
// 	"Shout" : {
//		"ShoutText" :
// 	}
//
//
//
//
//
//
