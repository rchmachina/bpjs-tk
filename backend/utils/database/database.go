package database

import (
	"fmt"
	"log"

	"time"

	"github.com/google/uuid"
	env "github.com/rchmachina/bpjs-tk/utils/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email          string
	UserName       string
	HashedPassword string
	User_id        uuid.UUID `gorm:"primaryKey"`
	Roles          string
	CreatedAt      time.Time
}

var DB *gorm.DB

func DatabaseConnection() *gorm.DB {
	var err error
	//
	getDomainName := env.GetConfig("app.db.domain_name")
	getUserNameDb := env.GetConfig("app.db.DB_USER")
	getPassDb := env.GetConfig("app.db.DB_PASS")
	getPortDb := env.GetConfig("app.db.PORT_DB")
	getNameDb := env.GetConfig("app.db.DB_NAME")

	log.Println("isi", getDomainName)

	DB, err = gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable ", getDomainName, getUserNameDb, getPassDb, getNameDb, getPortDb)), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")

	return DB

}
