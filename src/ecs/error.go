package ecs

import (
	"fmt"
)

type NoSuchEntityError struct {
	entityID string
}

func (err NoSuchEntityError) Error() string {
	return fmt.Sprintf("Entity " + err.entityID + " does not exist!")
}
