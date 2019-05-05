package api

import (
	"github.com/evansmwendwa/uxp/response"
	"github.com/labstack/echo"
)

// Status message for the home page
type Status struct {
	Message string `json:"message"`
}

// Home - index Page
func Home(c echo.Context) (err error) {
	var m Status

	m.Message = "Application running on 0.0.0.0:8080"

	return response.APIResponse(err, c, m)
}
