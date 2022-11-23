package models

import "time"

type Game struct {
	HomeTeam        Team
	AwayTeam        Team
	TotalScore      int
	CreatedDateTime time.Time
}
