package routes

import (
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	handlers "github.com/rchmachina/bpjs-tk/handler"
	repo "github.com/rchmachina/bpjs-tk/repositories"
	"github.com/rchmachina/bpjs-tk/utils/database"
)

func Soal1Route(e *echo.Group) {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})
	soal2Repo := repo.RepositorySoal1(database.DB)
	h := handlers.HandlerSoal1(soal2Repo, redisClient)
	e.POST("/concurent-soal1", h.SendConcurentData)
	e.POST("/soal1", h.PostData)
	e.GET("/get-redis", h.GetRedisData)
	e.DELETE("/soal1", h.DeleteDataSoal1)

}
