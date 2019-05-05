package response

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// ModelError for decorating responses
type ModelError struct {
	Message string `json:"message"`
}

// APIResponse - returns a decorated json response
func APIResponse(err error, c echo.Context, model interface{}) error {
	var e ModelError

	if gorm.IsRecordNotFoundError(err) {
		e.Message = err.Error()
		return c.JSON(http.StatusNotFound, e)
	} else if err != nil {
		e.Message = err.Error()
		return c.JSON(http.StatusBadRequest, e)
	} else {
		return c.JSON(http.StatusOK, model)
	}
}
