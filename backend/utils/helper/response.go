package helper

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
)

func JSONBulkResponse(c echo.Context, statusCode int, data interface{}) error {
	var statusMessage string

	switch statusCode {
	case 200:
		statusMessage = "Success"
	case 400:
		statusMessage = "Bad Request"
	case 404:
		statusMessage = "Not Found"
	case 501:
		statusMessage = "internal server error"
	default:
		statusMessage = "Unknown Status"
	}

	if statusMessage == "Unknown Status" {
		return c.JSON(statusCode, map[string]interface{}{
			"data": data,
			"status": map[string]interface{}{
				"code":    501,
				"message": "internal server error",
			},
		})
	}

	return c.JSON(statusCode, map[string]interface{}{
		"data": data,
		"status": map[string]interface{}{
			"code":    statusCode,
			"message": statusMessage,
		},
	})

}

func JSONResponse(c echo.Context, statusCode int, data interface{}) error {
	var statusMessage string

	switch statusCode {
	case 200:
		statusMessage = "Success"
	case 400:
		statusMessage = "Bad Request"
	case 404:
		statusMessage = "Not Found"
	case 501:
		statusMessage = "internal server error"
	default:
		statusMessage = "Unknown Status"
	}

	if statusMessage == "Unknown Status" {
		return c.JSON(statusCode, map[string]interface{}{
			"data": map[string]interface{}{
				"data": data,
			},
			"status": map[string]interface{}{
				"code":    501,
				"message": "internal server error",
			},
		})
	}

	return c.JSON(statusCode, map[string]interface{}{
		"data": data,
		"status": map[string]interface{}{
			"code":    statusCode,
			"message": statusMessage,
		},
	})

}

func ToJSON(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return "{}"
	}

	return string(b)
}
