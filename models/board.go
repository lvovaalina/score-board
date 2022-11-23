package models

import "fmt"

type Board struct {
	Games []Game
}

func (b Board) String() string {
	var gamesStr string
	for i, game := range b.Games {
		gamesStr += fmt.Sprintf("%d. %s - %s: %d - %d\n",
			i+1, game.HomeTeam.Name, game.AwayTeam.Name, game.HomeTeam.Score, game.AwayTeam.Score)
	}

	return gamesStr
}
