package services

import (
	"github.com/lvovaalina/score-board/models"
)

var initialTeamScore = 0

var board models.Board

func StartGame(homeTeamName string, awayTeamName string) {
	board.Games = append(board.Games, models.Game{
		HomeTeam: models.Team{
			Name:  homeTeamName,
			Score: initialTeamScore,
		},
		AwayTeam: models.Team{
			Name:  awayTeamName,
			Score: initialTeamScore,
		},
		TotalScore: initialTeamScore,
	})
}
