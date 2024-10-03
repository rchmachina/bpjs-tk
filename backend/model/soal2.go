package model

import (
    "github.com/google/uuid"
)

// GetDataSoal2 represents the structure for data retrieved from the database.
type GetDataSoal2 struct {
    ID       uuid.UUID `json:"id"`        // UUID type for ID
    NameData string     `json:"nameData"`  // String for the name
    ParentID *uuid.UUID `json:"parentId"`  // Pointer to UUID for parent ID, allows for null values
}
