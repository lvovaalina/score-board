package services

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

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
		TotalScore:      initialTeamScore,
		CreatedDateTime: time.Now(),
	})

	// sort games upon add to use game id for update and finish game
	sortGames()

	return
}

func UpdateScore(gamePositionInBoard int, newHomeTeamScore int, newAwayTeamScore int) (err error) {
	if gamePositionInBoard <= 0 || gamePositionInBoard > len(board.Games) {
		return errors.New("No game with this position!")
	}

	game := &board.Games[gamePositionInBoard-1]
	game.HomeTeam.Score = newHomeTeamScore
	game.AwayTeam.Score = newAwayTeamScore
	game.TotalScore = newHomeTeamScore + newAwayTeamScore

	sortGames()

	return
}

func FinishGame(gamePositionInBoard int) (err error) {
	if gamePositionInBoard <= 0 || gamePositionInBoard > len(board.Games) {
		return errors.New("No game with this position!")
	}

	games := helpers.RemoveFromSliceByIndex(board.Games, gamePositionInBoard-1)
	board.Games = make([]models.Game, 0)
	board.Games = append(board.Games, games...)
	return
}

func ShowBoard() {
	fmt.Println(board)
}

func sortGames() {
	sort.SliceStable(board.Games, func(i int, j int) bool {
		var sortedByTotalScore, sortedByDate bool

		sortedByTotalScore = board.Games[i].TotalScore > board.Games[j].TotalScore

		if board.Games[i].TotalScore == board.Games[j].TotalScore {
			sortedByDate = board.Games[i].CreatedDateTime.Before(board.Games[j].CreatedDateTime)
			return sortedByDate
		}

		return sortedByTotalScore
	})
}
