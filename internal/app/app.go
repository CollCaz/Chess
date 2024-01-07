package app

import data "Chess/internal/models"

var App = Application{}

type Application struct {
	Models *data.Models
}
