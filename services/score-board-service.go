package services

import (
	"errors"
	"strings"

	"github.com/lvovaalina/score-board/models"
)

var initialTeamScore = 0

var board models.Board

func StartGame(homeTeamName string, awayTeamName string) (err error) {
	if strings.Trim(homeTeamName, " ") == "" || strings.Trim(awayTeamName, " ") == "" {
		return errors.New("Team name(s) cannot be empty!")
	}

	if homeTeamName == awayTeamName {
		return errors.New("Away team name cannot be the same as home team name!")
	}

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

	return
}
