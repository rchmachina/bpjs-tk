package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/rchmachina/bpjs-tk/dto"

	"gorm.io/gorm"
)

// kontrak
type Soal1Repository interface {
	//postDataSoal1() ([]dto.GetDataSoal2, error)
	GetSoal1Repository(uuidData string) ([]dto.GetDataSoal2, error)
	PostSoal1Repository(dto.BulkDataSoal1) string
	DeleteSoal1Repository() error
}

func RepositorySoal1(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetSoal1Repository(uuidData string) ([]dto.GetDataSoal2, error) {
	var result string
	var getResult []dto.GetDataSoal2

	fmt.Println("masuk pak cik")
	paramsJSON, err := json.Marshal(map[string]interface{}{})
	if err != nil {
		return getResult, err
	}

	err = r.db.Raw("select * soal1.insert_bulk_data($1::jsonb)", string(paramsJSON)).Scan(&result).Error

	err = json.Unmarshal([]byte(result), &getResult)
	if err != nil {
		return getResult, errors.New("not found")
	}

	return getResult, err
}
func (r *repository) PostSoal1Repository(data dto.BulkDataSoal1) string {
	var result string

	fmt.Println("masuk pak cik")
	paramsJSON, err := json.Marshal(data)
	if err != nil {
		return result
	}
	log.Println("isi data", data)

	err = r.db.Raw("select * from soal1.insert_bulk_data($1::jsonb)", string(paramsJSON)).Scan(&result).Error

	return result
}
func (r *repository) DeleteSoal1Repository() error {
	query := "DELETE FROM soal1.\"data\" "

	result := r.db.Exec(query, )

	// Check for errors
	if result.Error != nil {
		return result.Error
	}

	// Optionally, you can check the number of rows affected
	if rowsAffected := result.RowsAffected; rowsAffected == 0 {
		return fmt.Errorf("no rows deleted")
	}

	return nil
}
