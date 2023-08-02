package two

import (
	"bufio"
	"fmt"
	"os"
)

var myPoints int = 0

var choice = map[string]string{
	"A": "ROCK",
	"B": "PAPER",
	"C": "SCISSORS",
	"X": "ROCK",
	"Y": "PAPER",
	"Z": "SCISSORS",
}

var outcome = map[string]int{
	"X": -1,
	"Y": 0,
	"Z": 1,
}

var points = map[string]int{
	"ROCK":     1,
	"PAPER":    2,
	"SCISSORS": 3,
}

func DayTwo1() {

	strategyTxt := "./day-2/big-input.txt"

	buffer, err := os.Open(strategyTxt)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(buffer)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		roundTxt := fileScanner.Text()

		opponentPlay := choice[roundTxt[0:1]]

		myPlay := choice[roundTxt[2:]]

		calculatePoints(opponentPlay, myPlay)

	}

	buffer.Close()

	fmt.Println(myPoints)
}

func DayTwo2() {

	strategyTxt := "./day-2/big-input.txt"

	buffer, err := os.Open(strategyTxt)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(buffer)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		roundTxt := fileScanner.Text()

		opponentPlay := choice[roundTxt[0:1]]

		roundOutcome := outcome[roundTxt[2:]]

		myPlay := determineMyPlay(opponentPlay, roundOutcome)

		calculatePoints(opponentPlay, myPlay)

	}

	buffer.Close()

	fmt.Println(myPoints)
}

func calculatePoints(opponentPlay, myPlay string) {
	myPoints = myPoints + points[myPlay]
	if opponentPlay == myPlay {
		myPoints = myPoints + 3
	} else if (myPlay == "ROCK" && opponentPlay == "SCISSORS") ||
		(myPlay == "PAPER" && opponentPlay == "ROCK") ||
		(myPlay == "SCISSORS" && opponentPlay == "PAPER") {

		myPoints = myPoints + 6
	}
}

func determineMyPlay(opponentPlay string, outcome int) string {
	if outcome == 0 {
		return opponentPlay
	} else if (outcome == 1 && opponentPlay == "SCISSORS") ||
		(outcome == -1 && opponentPlay == "PAPER") {
		return "ROCK"
	} else if (outcome == 1 && opponentPlay == "ROCK") ||
		(outcome == -1 && opponentPlay == "SCISSORS") {
		return "PAPER"
	} else if (outcome == 1 && opponentPlay == "PAPER") ||
		(outcome == -1 && opponentPlay == "ROCK") {
		return "SCISSORS"
	} else {
		return "unknown"
	}
}
