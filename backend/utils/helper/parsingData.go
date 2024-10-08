package helper

import (
	"fmt"

	"github.com/google/uuid"
)

func ParseUUID(id string) (uuid.UUID, error) {

	parsedUUID, err := uuid.Parse(id) // Convert the string to uuid.UUID
	if err != nil {
		fmt.Println("Error parsing UUID:", err)
		return parsedUUID, err
	}

	return parsedUUID, nil

}
