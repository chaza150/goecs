package io

import (
	"io/ioutil"
)

func LoadFile() string {
	data, err := ioutil.ReadFile("res/ecs/data")
	if err != nil {
		panic(err)
	}
	return string(data)
}
