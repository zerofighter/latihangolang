package controllers

import (
	"../models"
)

type (
	// For Get - /showtimes
	GoTestResource struct {
		Data []models.GoTest `json:"data"`
	}

	GoTestResourceOne struct {
		Data models.GoTest `json:"data"`
	}
)
