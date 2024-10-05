package routes

import (
	"fmt"
	"log"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	handlers "github.com/rchmachina/bpjs-tk/handler"
	repo "github.com/rchmachina/bpjs-tk/repositories"
	"github.com/rchmachina/bpjs-tk/utils/database"
	env "github.com/rchmachina/bpjs-tk/utils/env"
)

func Soal1Route(e *echo.Group) {

	db := fmt.Sprint(env.GetConfigWithDefaultSetting("app.db.redis.db", 0))
	hostRedist := fmt.Sprint(env.GetConfigWithDefaultSetting("app.db.redis.host", "localhost:6379"))

	dbInt, err := strconv.Atoi(db)
	if err != nil {
		log.Fatal(err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: hostRedist, // Redis server address
		DB:   dbInt,      // Use default DB
	})
	soal2Repo := repo.RepositorySoal1(database.DB)
	h := handlers.HandlerSoal1(soal2Repo, redisClient)
	e.POST("/concurent-soal1", h.SendConcurentData)
	e.POST("/soal1", h.PostData)
	e.GET("/soal1", h.GetSoal1)
	e.GET("/child-soal1", h.GetSoal1Detail)
	e.GET("/get-redis", h.GetRedisData)
	e.DELETE("/soal1", h.DeleteDataSoal1)

}
