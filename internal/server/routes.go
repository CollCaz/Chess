package server

import (
	data "Chess/internal/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", s.helloWorldHandler)
	e.GET("/games/:id", s.getGameByID)
	e.GET("/players/:id", s.getPlayerByID)
	e.POST("/players", s.insertPlayer)

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
	resp, err := data.GetGame(id)
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

	resp, err := data.GetGame(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *Server) insertPlayer(c echo.Context) error {
	fmt.Println("AAA")
	fmt.Println(c.Request())
	player := new(data.Player)
	// add player to database
	if err := c.Bind(player); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, player)
}
