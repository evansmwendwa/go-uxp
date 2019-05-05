package router

import (
	"github.com/evansmwendwa/uxp/api"
	"github.com/labstack/echo"
)

// Echo - Return Echo application instance
func Echo() *echo.Echo {
	e := echo.New()

	// web routes
	e.GET("/", api.Home)

	// api routes
	g := e.Group("/api/v1")

	g.GET("/speakers", api.GetSpeakers)
	g.POST("/speakers", api.AddSpeaker)

	return e
}
