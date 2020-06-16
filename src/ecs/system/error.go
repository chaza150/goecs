package system

import "fmt"

type NotGoableError struct {
	systemName string
}

func (err NotGoableError) Error() string {
	return fmt.Sprintf("System " + err.systemName + " is not Goable!")
}
