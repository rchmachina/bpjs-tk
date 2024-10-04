package routes

import (
	//"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	handlers "github.com/rchmachina/bpjs-tk/handler"
	repo "github.com/rchmachina/bpjs-tk/repositories"
	"github.com/rchmachina/bpjs-tk/utils/database"
)

func Soal2Route(e *echo.Group) {

	soal2Repo := repo.RepositorySoal2(database.DB)
	h := handlers.HandlerSoal2(soal2Repo)
	e.GET("/get-data-soal2", h.GetDataSoal2)
	e.GET("/get-child-soal2", h.GetChild)
}
