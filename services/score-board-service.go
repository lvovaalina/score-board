package services

import (
	"errors"
	"fmt"
	"strings"

	"github.com/lvovaalina/score-board/helpers"
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

	var isAnyTeamAlreadyPlaingArr []bool = helpers.Map(board.Games, func(game models.Game) bool {
		return (game.AwayTeam.Name == homeTeamName || (game.HomeTeam.Name == homeTeamName) ||
			(game.AwayTeam.Name == awayTeamName) || (game.HomeTeam.Name == awayTeamName))
	})

	if helpers.Contains(isAnyTeamAlreadyPlaingArr, true) {
		return errors.New("Team already plaing!")
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

func ShowBoard() {
	fmt.Println(board)
}
