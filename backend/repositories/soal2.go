package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/rchmachina/bpjs-tk/model"

	_ "github.com/lib/pq"

	"gorm.io/gorm"
)

// kontrak
type Soal2Repository interface {
	GetDataSoal2() ([]model.GetDataSoal2, error)
	GetChildDataSoal2(uuidData string) ([]model.GetDataSoal2, error)
}

func RepositorySoal2(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetDataSoal2() ([]model.GetDataSoal2, error) {
	var result string
	var getResult []model.GetDataSoal2

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

func (r *repository) GetChildDataSoal2(uuidData string) ([]model.GetDataSoal2, error) {

	var getResult []model.GetDataSoal2
	var result string

	fmt.Println("masuk pak cik")

	id, err := uuid.Parse(uuidData)
	if err != nil {
		fmt.Println("Error parsing UUID:", err)
		return getResult, err
	}
	log.Println("passed 2 :", id)
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
