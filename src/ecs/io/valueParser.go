package io

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	ValueRegex      = `(\".+\"|-?\d+(\.\d+)?|-?\d+|(true|false)|(\[.*\]))`
	listRegex       = `\[\s*(((\s*\".+\"\s*,\s*)*(\s*\".+\"\s*))|((\s*\d+\s*,\s*)*(\s*\d+\s*))|((\s*\d+(\.\d+)?\s*,\s*)*(\s*\d+(\.\d+)?\s*))|((\s*(true|false)\s*,\s*)*(\s*(true|false)\s*)))\]`
	stringListRegex = `\[\s*((\s*\".+\"\s*,\s*)*(\s*\".+\"\s*))\]`
	intListRegex    = `\[\s*((\s*-?\d+\s*,\s*)*(\s*-?\d+\s*))\]`
	floatListRegex  = `\[\s*((\s*-?\d+(\.\d+)?\s*,\s*)*(\s*-?\d+(\.\d+)?\s*))\]`
	boolListRegex   = `\[\s*((\s*(true|false)\s*,\s*)*(\s*(true|false)\s*))\]`
)

func ParseInt(text string) int {
	if i, err := strconv.Atoi(text); err == nil {
		return i
	} else {
		panic(err)
	}
}

func ParseFloat(text string) float64 {
	if f, err := strconv.ParseFloat(text, 64); err == nil {
		return f
	} else {
		panic(err)
	}
}

func ParseBool(text string) bool {
	if b, err := strconv.ParseBool(text); err == nil {
		return b
	} else {
		panic(err)
	}
}

func ParseQuotedString(text string) string {
	if matched, _ := regexp.MatchString(`\"\w+\"`, text); matched {
		return strings.Trim(text, "\"")
	} else {
		panic("Could not parse quoted string:\n" + text)
	}
}

func ParseStringList(text string) []string {
	if matched, _ := regexp.MatchString(stringListRegex, text); matched {
		listString := strings.Trim(text, "[]")
		items := strings.Split(listString, ",")
		var result []string
		for _, item := range items {
			cleanItem := strings.TrimSpace(item)
			result = append(result, ParseQuotedString(cleanItem))
		}
		return result
	} else {
		panic("Could not parse string list:\n" + text)
	}
}

func ParseIntList(text string) []int {
	if matched, _ := regexp.MatchString(intListRegex, text); matched {
		listString := strings.Trim(text, "[]")
		items := strings.Split(listString, ",")
		var result []int
		for _, item := range items {
			cleanItem := strings.TrimSpace(item)
			result = append(result, ParseInt(cleanItem))
		}
		return result
	} else {
		panic("Could not parse int list:\n" + text)
	}
}

func ParseFloatList(text string) []float64 {
	if matched, _ := regexp.MatchString(floatListRegex, text); matched {
		listString := strings.Trim(text, "[]")
		items := strings.Split(listString, ",")
		var result []float64
		for _, item := range items {
			cleanItem := strings.TrimSpace(item)
			result = append(result, ParseFloat(cleanItem))
		}
		return result
	} else {
		panic("Could not parse float list:\n" + text)
	}
}

func ParseBoolList(text string) []bool {
	if matched, _ := regexp.MatchString(boolListRegex, text); matched {
		listString := strings.Trim(text, "[]")
		items := strings.Split(listString, ",")
		var result []bool
		for _, item := range items {
			cleanItem := strings.TrimSpace(item)
			result = append(result, ParseBool(cleanItem))
		}
		return result
	} else {
		panic("Could not parse bool list:\n" + text)
	}
}
