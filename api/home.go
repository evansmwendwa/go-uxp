package api

import (
	"net/http"

	"github.com/labstack/echo"
)

// Home - index Page
func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Application Index")
}
