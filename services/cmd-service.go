package services

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Command string

const (
	StartGameCmd   Command = "start"
	UpdateScoreCmd         = "update"
	FinishGameCmd          = "finish"
	QuitCmd                = "quit"
)

func (c Command) isCommand() bool {
	switch c {
	case StartGameCmd:
		return true
	case UpdateScoreCmd:
		return true
	case FinishGameCmd:
		return true
	case QuitCmd:
		return true
	}
	return false
}

var reader *bufio.Reader

func StartBoardCmd() {

	fmt.Println("Welcome to Football World Cup Score Board!")
	fmt.Println("Availible commands:\n" +
		"Start game: " + StartGameCmd + "\n" +
		"Update score: " + UpdateScoreCmd + "\n" +
		"Finish game: " + FinishGameCmd + "\n" +
		"Quit: " + QuitCmd)

	command := make(chan string)

	go func(in chan string) {
		reader = bufio.NewReader(os.Stdin)
		for {
			s, err := reader.ReadString('\n')
			if err != nil {
				close(in)
				fmt.Println("Error in read string", err)
			}
			in <- s
		}
	}(command)

exit:
	for {
		select {
		case in := <-command:
			in = strings.TrimSpace(in)
			if in == QuitCmd {
				break exit
			}
			fmt.Println("Read from stdin: ", in)
		}
	}

	fmt.Println("Thanks for using Football World Cup Score Board! Have a nice day!")
}

func ExecuteCommand(commandStr string) {
	if !(Command(strings.TrimSuffix(commandStr, "\n")).isCommand()) {
		fmt.Println("Unknown command")
	}
}
