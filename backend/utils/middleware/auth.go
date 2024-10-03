package middleware

import (
	"encoding/json"

	"net/http"
	"strings"

	jwtToken "github.com/rchmachina/bpjs-tk/utils/jwt"

	"github.com/labstack/echo/v4"
)

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Auth(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		token := c.Request().Header.Get("Authorization")

		if token == "" {
			return c.JSON(http.StatusUnauthorized, Result{Code: http.StatusUnauthorized, Message: "unathorized"})
		}

		token = strings.Split(token, "Bearer ")[1]

		claims, err := jwtToken.DecodeToken(token)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, Result{Code: http.StatusUnauthorized, Message: "unathorized"})
		}

		c.Set("userLogin", claims)

		return next(c)
	}
}

func UnmarshalToken(c echo.Context) (User, error) {
	userMap := c.Get("userLogin")

	// Marshal the map back to JSON bytes
	jsonData, err := json.Marshal(userMap)
	if err != nil {
		return User{}, err
	}

	// Unmarshal the JSON bytes into the User struct
	var user User
	if err := json.Unmarshal(jsonData, &user); err != nil {
		return User{}, err
	}

	return user, nil
}

type User struct {
	Expiry   float64 `json:"expiry"`
	ID       string  `json:"id"`
	Roles    string  `json:"roles"`
	UserName string  `json:"userName"`
}
