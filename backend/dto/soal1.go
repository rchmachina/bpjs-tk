package dto

import (
	"github.com/google/uuid"
)

// GetDataSoal2 represents the structure for data retrieved from the database.
type PostDataSoal1 struct {
	ID        uuid.UUID `json:"id"`
	StartDate string    `json:"startDate"`
	EndDate   string    `json:"endDate"`
	Nominal   float64   `json:"nominal"`
}
type GetDataSoal1 struct {
	ID        uuid.UUID `json:"id"`
	StartDate string    `json:"startDate"`
	EndDate   string    `json:"endDate"`
	Nominal   float64   `json:"nominal"`
}

type ChildDataDetail struct {
	ParentId  uuid.UUID  `json:"parentId"`
	StartDate string     `json:"startDate"`
	EndDate   string     `json:"endDate"`
	Id        *uuid.UUID `json:"id"`
	Nominal   float64    `json:"nominal"`
}

type PostChildData struct {
	//ParentId  uuid.UUID  `json:"parentId"`
	StartDate string  `json:"startDate"`
	EndDate   string  `json:"endDate"`
	Nominal   float64 `json:"nominal"`
}
type DataSoal1 struct {
	ParentData GetDataSoal1    `json:"parentData"`
	ChildData  []PostChildData `json:"childData"`
}
type BulkDataSoal1 struct {
	Data []DataSoal1 `json:"data"`
}
