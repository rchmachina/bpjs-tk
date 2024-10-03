package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	routes "github.com/rchmachina/bpjs-tk/route"
	"github.com/rchmachina/bpjs-tk/utils/database"
	env "github.com/rchmachina/bpjs-tk/utils/env"
)

func main() {
	//get the env file
	log.Println("asuuu")
	e := echo.New()
	//cors for api
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	fmt.Println("cors already activated")
	log.Println("u here?")
	//connnect to db
	_ = database.DatabaseConnection()
	fmt.Println("connected")

	baseApi := fmt.Sprint(env.GetConfigWithDefaultSetting("app.server.api_uri_base", "/api/v1"))
	baseAdress := fmt.Sprint(env.GetConfigWithDefaultSetting("app.server.address", ":8888"))

	routes.RouteInit(e.Group(baseApi))
	fmt.Println("server running localhost", baseAdress)

	e.Logger.Fatal(e.Start(baseAdress))

}
