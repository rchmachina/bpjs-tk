package repositories

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/rchmachina/bpjs-tk/model"

	_ "github.com/lib/pq"

	"gorm.io/gorm"
)

// kontrak
type UserRepository interface {
	CreateUserDb(user model.CreateUser) (string, error)
	LoginUserDB(string) (model.LoginResponse, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateUserDb(user model.CreateUser) (string, error) {
	var result string

	fmt.Println("masuk pak cik")
	paramsJSON, err := json.Marshal(user)
	if err != nil {
		return result, err
	}

	err = r.db.Raw("SELECT * from users.create_users($1::jsonb)", string(paramsJSON)).Scan(&result).Error

	return result, err
}

func (r *repository) LoginUserDB(userName string) (model.LoginResponse, error) {

	var responseLogin model.LoginResponse
	var result string

	paramsJSON, err := json.Marshal(map[string]interface{}{"userName": userName})
	if err != nil {
		return responseLogin, err
	}

	err = r.db.Raw("SELECT * FROM pklbram.login_user($1::jsonb)", string(paramsJSON)).Scan(&result).Error
	if err != nil {
		return responseLogin, err
	}
	err = json.Unmarshal([]byte(result), &responseLogin)
	if err != nil {
		return responseLogin, errors.New("not found")
	}
	fmt.Print("isi response", responseLogin)

	return responseLogin, err
}
