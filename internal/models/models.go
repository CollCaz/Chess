package data

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Models struct {
	Player *PlayerModel
	Game   *GameModel
}

func NewModels(db *pgxpool.Pool) *Models {
	return &Models{
		Player: &PlayerModel{DB: db},
		Game:   &GameModel{DB: db},
	}
}
