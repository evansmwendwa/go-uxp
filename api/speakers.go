package api

import (
	"github.com/evansmwendwa/uxp/db"
	"github.com/evansmwendwa/uxp/model"
	"github.com/evansmwendwa/uxp/response"
	"github.com/labstack/echo"
)

// GetSpeakers - Speaker page handler
func GetSpeakers(c echo.Context) error {
	db := db.Session()
	speakers := []model.Speaker{}
	err := db.Find(&speakers).Error

	return response.APIResponse(err, c, speakers)
}

// AddSpeaker - Add a new speaker
func AddSpeaker(c echo.Context) error {
	speaker := new(model.Speaker)
	if err := c.Bind(speaker); err != nil {
		return err
	}

	db := db.Session()
	err := db.Create(&speaker).Error

	return response.APIResponse(err, c, speaker)
}
