package models

import "fmt"

type Board struct {
	Games []Game
}

func (b Board) String() string {
	var gamesStr string
	for _, game := range b.Games {
		gamesStr += fmt.Sprintf("%s - %s: %d - %d\n",
			game.HomeTeam.Name, game.AwayTeam.Name, game.HomeTeam.Score, game.AwayTeam.Score)
	}

	return gamesStr
}
