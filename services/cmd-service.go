package services

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/lvovaalina/score-board/helpers"
)

type Command string

const (
	StartGameCmd   Command = "start"
	UpdateScoreCmd         = "update"
	FinishGameCmd          = "finish"
	ShowBoardCmd           = "show"
	QuitCmd                = "quit"
)

var reader *bufio.Reader
var wg sync.WaitGroup
var isTaskExecuting bool

func StartBoardCmd() {

	fmt.Println("Welcome to Football World Cup Score Board!")
	fmt.Println("Availible commands:\n" +
		"Start game: " + StartGameCmd + "\n" +
		"Update score: " + UpdateScoreCmd + "\n" +
		"Finish game: " + FinishGameCmd + "\n" +
		"Quit: " + QuitCmd + "\n\n\n")

	reader = bufio.NewReader(os.Stdin)
exit:
	for {
		fmt.Print(">")
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error in read string", err)
		}
		command = strings.TrimSpace(command)
		if command == QuitCmd {
			break exit
		}
		ExecuteCommand(command)
	}
}

func ExecuteCommand(commandStr string) {
	Command(strings.TrimSuffix(commandStr, "\n")).Execute()
	isTaskExecuting = false
}

func (c Command) Execute() {
	switch c {
	case StartGameCmd:
		startGameCmdHandler()
		return
	case UpdateScoreCmd:
		UpdateScoreCmdHandler()
		return
	case FinishGameCmd:
		FinishGameCmdHandler()
		return
	case ShowBoardCmd:
		ShowBoardCmdHandler()
	case QuitCmd:
		return
	default:
		fmt.Println("Unknown command")
		return
	}
}

func startGameCmdHandler() {
	fmt.Print(">Home team name: ")
	homeTeamName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading home team name", err)
		return
	}

	fmt.Print(">Away team name: ")
	awayTeamName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading away team name", err)
		return
	}

	errStart := StartGame(strings.TrimSuffix(homeTeamName, "\n"), strings.TrimSuffix(awayTeamName, "\n"))
	if errStart != nil {
		fmt.Println("Game start faled: ", errStart)
		return
	}

	fmt.Println("Game started succesfully!")
}

func FinishGameCmdHandler() {
	fmt.Print(">Game to finish position: ")
	positionString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading game position", err)
		return
	}

	positionNumber, err := helpers.ParseToInt(strings.TrimSuffix(positionString, "\n"))
	if err == nil {
		finishErr := FinishGame(positionNumber)
		if finishErr != nil {
			fmt.Println("Finish game failed: ", finishErr)
			return
		}
		fmt.Println("Game finished succesfully!")
	}
}

func UpdateScoreCmdHandler() {
	fmt.Print(">Game to update position: ")

	positionString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading game position", err)
		return
	}

	positionNumber, err := helpers.ParseToInt(strings.TrimSuffix(positionString, "\n"))
	if err != nil {
		return
	}

	fmt.Print(">Home team score: ")
	homeTeamScoreString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading team score", err)
		return
	}

	homeTeamScore, err := helpers.ParseToInt(strings.TrimSuffix(homeTeamScoreString, "\n"))
	if err != nil {
		return
	}

	fmt.Print(">Home team score: ")
	awayTeamScoreString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading team score", err)
		return
	}
	awayTeamScore, err := helpers.ParseToInt(strings.TrimSuffix(awayTeamScoreString, "\n"))
	if err != nil {
		return
	}

	updErr := UpdateScore(positionNumber, homeTeamScore, awayTeamScore)
	if updErr != nil {
		fmt.Println("Game score update failed: ", updErr)
		return
	}
	fmt.Println("Game updated succesfully!")
}

func ShowBoardCmdHandler() {
	ShowBoard()
}
