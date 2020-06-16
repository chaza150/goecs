package io

import (
	"fmt"
	"io/ioutil"
)

func LoadFile() {
	data, err := ioutil.ReadFile("/res/ecs/data")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file: ", string(data))
}
