package handlers

import (
	"log"

	"github.com/rchmachina/bpjs-tk/repositories"
	"github.com/rchmachina/bpjs-tk/utils/helper"

	"github.com/labstack/echo/v4"
)

type soal2Handler struct {
	soal2 repositories.Soal2Repository
}

func HandlerSoal2(soal2 repositories.Soal2Repository) *soal2Handler {
	return &soal2Handler{soal2}
}

func (h *soal2Handler) GetDataSoal2(c echo.Context) error {

	getAllData, err := h.soal2.GetDataSoal2()
	if err != nil {
		return helper.JSONResponse(c, 501, err)

	}

	return helper.JSONBulkResponse(c, 200, getAllData)

}

func (h *soal2Handler) GetChild(c echo.Context) error {
	id := c.QueryParam("id")
	log.Println("iisi id", id)
	getAllChildData, err := h.soal2.GetChildDataSoal2(id)
	if err != nil {
		return helper.JSONResponse(c, 501, err.Error())

	}

	return helper.JSONBulkResponse(c, 200, getAllChildData)

}
