package router

import (
	"github.com/evansmwendwa/uxp/api"
	"github.com/labstack/echo"
)

// Echo - Return Echo application instance
func Echo() *echo.Echo {
	e := echo.New()

	e.GET("/", api.Home)

	// speakers
	e.GET("/speakers", api.GetSpeakers)
	e.POST("/speakers", api.AddSpeaker)

	return e
}
