package server

import (
	"Chess/internal/app"
	data "Chess/internal/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", s.helloWorldHandler)
	e.GET("/games/:id", s.getGameByID)
	e.GET("/players/:id", s.getPlayerByID)
	e.POST("/players", s.insertPlayer)
	e.POST("/games", s.insertGame)

	return e
}

func (s *Server) helloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *Server) getGameByID(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	resp, err := app.App.Models.Game.GetGame(id)
	if err != nil {
		panic("wtf")
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *Server) getPlayerByID(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	resp, err := app.App.Models.Player.GetPlayer(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *Server) insertPlayer(c echo.Context) error {
	player := data.Player{}
	// add player to database
	if err := c.Bind(&player); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := app.App.Models.Player.InsertPlayer(&player)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSONPretty(http.StatusOK, player, "  ")
}

func (s *Server) insertGame(c echo.Context) error {
	game := data.Game{}
	if err := c.Bind(&game); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := app.App.Models.Game.InsertGame(&game)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSONPretty(http.StatusOK, game, "  ")
}
