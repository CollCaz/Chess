package data

import "encoding/json"
import (
	"encoding/json"
	"github.com/jackc/pgx/v5/pgxpool"
)
type PlayerModel struct {
	DB pgxpool.Pool
}

type Player struct {
	ID             int
	Name           string
	Version        int
	CreatedAt      time.Time
	WorldRank      int
	BYear          int
	Federation     string
	Sex            string
	FideID         int
	FideTttle      string
	StandardRating int
	RapidRating    int
	BlitzRating    int
}

func (p Player) MarshalJSON() ([]byte, error) {
	rating := make(map[string]int)
	rating["Standard"] = p.StandardRating
	rating["Rapid"] = p.RapidRating
	rating["Blitz"] = p.BlitzRating

	aux := struct {
		Name       string
		WorldRank  int
		BYear      int
		Federation string
		Sex        string
		FideID     int
		FideTttle  string
		Rating     map[string]int
	}{
		Name:       p.Name,
		WorldRank:  p.WorldRank,
		BYear:      p.BYear,
		Federation: p.Federation,
		Sex:        p.Sex,
		FideID:     p.FideID,
		FideTttle:  p.FideTttle,
		Rating:     rating,
	}

	return json.Marshal(aux)
}

func GetPlayer(id int) (*Player, error) {
	player := Player{
		Name:           "Carlson, Magnus",
		WorldRank:      1,
		BYear:          1990,
		Federation:     "Norway",
		Sex:            "Male",
		FideID:         1503014,
		FideTttle:      "Grandmaster",
		StandardRating: 2830,
		RapidRating:    2823,
		BlitzRating:    2886,
	}
	return &player, nil
}
