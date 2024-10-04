package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rchmachina/bpjs-tk/dto"
	"github.com/google/uuid"

	_ "github.com/lib/pq"

	"gorm.io/gorm"
)

// kontrak
type Soal2Repository interface {
	GetDataSoal2() ([]dto.GetDataSoal2, error)
	GetChildDataSoal2(uuidData string) ([]dto.GetDataSoal2, error)
}

func RepositorySoal2(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetDataSoal2() ([]dto.GetDataSoal2, error) {
	var result string
	var getResult []dto.GetDataSoal2

	fmt.Println("masuk pak cik")
	paramsJSON, err := json.Marshal(map[string]interface{}{})
	if err != nil {
		return getResult, err
	}

	err = r.db.Raw("select * from soal2.get_all_data($1::jsonb)", string(paramsJSON)).Scan(&result).Error

	err = json.Unmarshal([]byte(result), &getResult)
	if err != nil {
		return getResult, errors.New("not found")
	}

	return getResult, err
}

func (r *repository) GetChildDataSoal2(uuidData string) ([]dto.GetDataSoal2, error) {

	var getResult []dto.GetDataSoal2
	var result string

	fmt.Println("masuk pak cik")

	id, err := uuid.Parse(uuidData)
	if err != nil {
		fmt.Println("Error parsing UUID:", err)
		return getResult, err
	}

	paramsJSON, err := json.Marshal(map[string]interface{}{"uuidParent": id})
	if err != nil {
		return getResult, err
	}
	err = r.db.Raw("select * from soal2.get_children_data($1::jsonb)", string(paramsJSON)).Scan(&result).Error
	if err != nil {
		return getResult, err
	}
	err = json.Unmarshal([]byte(result), &getResult)
	if err != nil {
		return getResult, errors.New("not found")
	}

	return getResult, err
}
