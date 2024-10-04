package dto

import "time"

type Response struct {
	Message interface{} `json:"message"`
}


type DateRange struct {
	StartDate time.Time
	EndDate   time.Time
}
