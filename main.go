package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CurrentRound struct {
	currentTurn string // x or o
	winner string // x or o or ""
}

var currentRound CurrentRound

var gameMatrix = [3][3]string{
	{"1", "2", "3"},
	{"4", "5", "6"},
	{"7", "8", "9"},
}




func resetGame() {
	// resetMatrix

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			gameMatrix[i][j] = strconv.Itoa((3*i + j) + 1)
		}
	}

	currentRound.winner = ""
}

func getUserInput(message string) string {
	if len(message) > 0 {
		fmt.Println(message)
	}
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')

	userInput = strings.ToUpper(strings.TrimSpace(userInput))
	return userInput
}

func chooseXO() {
	for true {

		choice := getUserInput("First Player: Choose X or O")

		if choice == "X" || choice == "O" {
			currentRound.currentTurn = choice
			break
		} else {
			fmt.Println("Invalid input, Try Again")
		}
	}
}

func displayGameMatrix() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf(`| %v |`, gameMatrix[i][j])
		}
		fmt.Println()
	}
}

func playTurn() {
	for true {
		fmt.Printf("%v's Turn \n", currentRound.currentTurn)
		playPosition := getUserInput("")

		playPositionInt, _ := strconv.Atoi(playPosition)

		rowIdx, columnIdx := detectPositionInMatrix(playPositionInt)

		if isOutOfbounds(playPositionInt) {
			fmt.Println("Out of range input, Please choose a number between 1 and 9.")
		} else if !(isPositionAvailable(rowIdx, columnIdx)) {
			fmt.Println("Position is already used")
		} else {
			// input belongs && position is available
			gameMatrix[rowIdx][columnIdx] = currentRound.currentTurn
			break
		}
	}
}

func changeTurn() {
	if currentRound.currentTurn == "X" {
		currentRound.currentTurn = "O"
	} else {
		currentRound.currentTurn = "X"
	}
}

// Number Theory Modular Arithmetc
func detectPositionInMatrix(playPositionInt int) (int, int) {

	rowIdx := int(math.Ceil(float64(playPositionInt) / 3))

	columnIdx := (playPositionInt % 3)

	if columnIdx == 0 {
		// a new cycle is complete, move to the last index which is the third or idx # 2
		columnIdx = 3
	}

	return rowIdx - 1, columnIdx - 1
}

func isPositionAvailable(row int, column int) bool {
	if gameMatrix[row][column] == "X" || gameMatrix[row][column] == "O" {
		return false
	} else {
		return true
	}
}

func isOutOfbounds(input int) bool {
	if input > 9 || input < 1 {
		return true
	} else {
		return false
	}
}

func checkWinCases() bool {
	if checkLeftDiagonalWin() || checkRightDiagonalWin() || checkRowsAndColumnsWin() {
		currentRound.winner = currentRound.currentTurn
		return true
	}

	return false
}

func playTurns() {
	i := 1
	for i <= 9 {
		displayGameMatrix()
		playTurn()
		winCaseFound := checkWinCases()
		if winCaseFound {
			break
		}
		changeTurn()
		i++
	}
}

func checkRowsAndColumnsWin() bool {
	i := 0
	isWin := false
	
	for i <= 2 {
				
		rowWinCase := gameMatrix[i][0] == gameMatrix[i][1] && gameMatrix[i][1] == gameMatrix[i][2]
		
		columnWinCase := gameMatrix[0][i] == gameMatrix[1][i] && gameMatrix[1][i] == gameMatrix[2][i]
		
		if rowWinCase || columnWinCase {
			isWin = true
			break
		}

		i++
	}

	return isWin
}

func checkLeftDiagonalWin() bool {
	i := 0

	isEqual := true

	for i <= 1 {
		if gameMatrix[i][i] != gameMatrix[i+1][i+1] {
			isEqual = false
		}
		i++
	}

	return isEqual
}

func checkRightDiagonalWin() bool {
	row, column := 0, 2
	isEqual := true
	for column >= 1 {
		if gameMatrix[row][column] != gameMatrix[row+1][column-1] {
			isEqual = false
			break
		}
		row++
		column--
	}

	return isEqual
}

func declareWinState() {
	displayGameMatrix()
	if len(currentRound.winner) == 0 {
		fmt.Println("Draw")
	} else {
		fmt.Printf("%v Won \n", currentRound.winner)
	}
}

func playAgain() {
	playAgain := false
	for true {
		answer := getUserInput("Do you want to play again ? Yes or No")
		if answer == "YES" {
			playAgain = true
			break
		} else if answer == "NO" {
			fmt.Println("Good Bye")
			break
		} else {
			fmt.Println("Invalid input, Please enter Yes or No")
		}
	}
	
	if playAgain {
		resetGame()
		startGame()
	}
}

func startGame() {
	chooseXO()
	playTurns()
	declareWinState()
	playAgain()
}

func main() {
	startGame()
}