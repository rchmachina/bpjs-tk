package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	_ "github.com/lib/pq"
	"github.com/rchmachina/bpjs-tk/dto"

	//"github.com/rchmachina/bpjs-tk/utils/helper"
	"gorm.io/gorm"
)

// kontrak
type Soal1Repository interface {
	GetSoal1Repository(pageInt, limitInt int) ([]dto.GetDataSoal1, error)
	GetChildSoal1Repository(id string) ([]dto.ChildDataDetail, error)
	PostSoal1Repository(dto.BulkDataSoal1) string
	DeleteSoal1Repository() error
	PostSoal1ChildData(data dto.PostChildData) (string, error)
	PostSoal1ParentData(data dto.GetDataSoal1) (string, error)
}

func RepositorySoal1(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetSoal1Repository(pageInt, limitInt int) ([]dto.GetDataSoal1, error) {
	var result string
	var getResult []dto.GetDataSoal1

	fmt.Println("masuk pak cik")
	paramsJSON, err := json.Marshal(map[string]interface{}{
		"page":  pageInt,
		"limit": limitInt,
	})
	if err != nil {
		return getResult, err
	}

	err = r.db.Raw("select * from soal1.get_all_data_with_pagination($1::jsonb)", string(paramsJSON)).Scan(&result).Error

	err = json.Unmarshal([]byte(result), &getResult)
	if err != nil {
		return getResult, errors.New("not found")
	}

	return getResult, err
}

func (r *repository) PostSoal1Repository(data dto.BulkDataSoal1) string {
	// Step 1: Insert parent data
	baseQuery := `
    INSERT INTO soal1.data (start_date, end_date, nominal)
    VALUES `

	var parentValueStrings []string
	var parentValues []interface{}

	// Loop through the inputData and append each parent's values
	for i, parent := range data.Data {
		startIdx := i * 3
		parentValueStrings = append(parentValueStrings, fmt.Sprintf("($%d, $%d, $%d)", startIdx+1, startIdx+2, startIdx+3))
		parentValues = append(parentValues, parent.ParentData.StartDate, parent.ParentData.EndDate, parent.ParentData.Nominal)
	}

	// Finalize the parent query and execute
	parentQuery := baseQuery + strings.Join(parentValueStrings, ", ") + ` RETURNING id`
	rows, err := r.db.Raw(parentQuery, parentValues...).Rows() // Use Raw for custom queries and Rows for fetching results
	if err != nil {
		return fmt.Sprintf("failed to execute parent insert query: %v", err)
	}
	defer rows.Close()

	// Step 2: Process child data for each parent
	var childValueStrings []string
	var childValues []interface{}
	
	var childValueCounter = 0
	var parentCounter = 0
	baseQueryChild := `
		INSERT INTO soal1.data_detail (start_date, end_date, nominal, parent_id)
		VALUES `

	// Loop over returned parent IDs and corresponding child data
	for rows.Next() {
		var parentID string
		if err := rows.Scan(&parentID); err != nil {
			return fmt.Sprintf("failed to scan parent ID: %v", err)
		}

		// Process and append child data for the current parent
		currentParent := data.Data[parentCounter]
		for _, child := range currentParent.ChildData {
			childValueCounter += 4

			// Add placeholders for each child record
			childValueStrings = append(childValueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d)", childValueCounter-3, childValueCounter-2, childValueCounter-1, childValueCounter))
			childValues = append(childValues, child.StartDate, child.EndDate, child.Nominal, parentID)

			queryChild := baseQueryChild + strings.Join(childValueStrings, ", ")

			// Execute the query
			err = r.db.Exec(queryChild, childValues...).Error
			if err != nil {
				return fmt.Sprintf("failed to execute child insert query: %v", err)
			}
		
		}

		// Increment parent counter
		parentCounter++

	}



	log.Println("Successfully inserted all parent and child records.")
	return "Successfully inserted all records."
}


func (r *repository) PostSoal1ChildData(data dto.PostChildData) (string, error) {
	var result string

	fmt.Println("masuk pak cik di childnya ")
	paramsJSON, err := json.Marshal(data)
	if err != nil {
		return result, err
	}

	err = r.db.Raw("select * from soal1.create_child_data($1::jsonb)", string(paramsJSON)).Scan(&result).Error

	return result, err
}
func (r *repository) PostSoal1ParentData(data dto.GetDataSoal1) (string, error) {
	var result string

	fmt.Println("masuk pak cik di parennya")
	paramsJSON, err := json.Marshal(data)
	if err != nil {
		return result, err
	}
	log.Println("isi data", data)

	err = r.db.Raw("select * from soal1.create_parent_data($1::jsonb)", string(paramsJSON)).Scan(&result).Error
	log.Println("get uuid from post soal1 ", result)
	return result, err
}

func (r *repository) DeleteSoal1Repository() error {
	query := "DELETE FROM soal1.\"data\" "

	result := r.db.Exec(query)

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
func (r *repository) GetChildSoal1Repository(id string) ([]dto.ChildDataDetail, error) {

	var result string
	var getResult []dto.ChildDataDetail

	fmt.Println("masuk pak cik")
	paramsJSON, err := json.Marshal(map[string]interface{}{
		"parentId": id,
	})
	if err != nil {
		return getResult, err
	}

	err = r.db.Raw("select * from soal1.get_child_data($1::jsonb)", string(paramsJSON)).Scan(&result).Error

	err = json.Unmarshal([]byte(result), &getResult)
	if err != nil {
		return getResult, errors.New("not found")
	}

	return getResult, err
}
