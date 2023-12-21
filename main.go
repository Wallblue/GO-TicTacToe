package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Player int

const (
	Player1 Player = iota // 0
	Player2               // 1
)

type Board struct {
	boxes [3][3]rune
}

func (b *Board) UpdateBox(line int, column int, currPlayer Player) {
	if currPlayer == Player1 {
		b.boxes[line][column] = 'X'
	} else {
		b.boxes[line][column] = 'O'
	}
}

func main() {
	var board Board
	var currentPlayer Player = Player1
	var turnNumber int = 0
	in := bufio.NewScanner(os.Stdin)

	board.boxes = initializeBoardBoxes(board.boxes)

	for {
		turnNumber++
		fmt.Printf("\nTurn of Player %v\n", currentPlayer+1)
		printBoard(board)

		board = boxSelection(currentPlayer, board, in) //Players play here
		switchPlayer(currentPlayer)

		if isGameOver(board) {
			fmt.Println("\nPLAYER", currentPlayer+1, "WON !!!")
			printBoard(board)
			break
		}

		if turnNumber == 9 {
			fmt.Println("\nTIE ! No winner.")
			printBoard(board)
			break
		}
		currentPlayer = switchPlayer(currentPlayer)
	}
}

func initializeBoardBoxes(boxes [3][3]rune) [3][3]rune {
	for i := 0; i < len(boxes); i++ {
		for j := 0; j < len(boxes[i]); j++ {
			boxes[i][j] = ' ' //Initializing every box with a space, if played it becomes an O or an X
		}
	}
	return boxes
}

func printBoard(b Board) {
	for i := 0; i < len(b.boxes); i++ {
		fmt.Println("+---+---+---+")
		for j := 0; j < len(b.boxes[i]); j++ {
			fmt.Printf("| %c ", b.boxes[i][j])
		}
		fmt.Print("|\n")
	}
	fmt.Println("+---+---+---+")
}

func boxSelection(currPlayer Player, currBoard Board, in *bufio.Scanner) Board {
	var (
		line, column = 0, 0
		err          error
	)
	for isOccupied(line, column, currBoard) {
		line, column = 0, 0
		for line == 0 { //We left when line has not its default value anymore
			fmt.Print("Choose the line you want to play on : ")
			in.Scan()
			line, err = strconv.Atoi(in.Text())
			if !isLinesColumnsOk(line, err, currBoard) {
				line = 0
				continue
			}
		}

		for column == 0 {
			fmt.Print("Choose the column you want to play on : ")
			in.Scan()
			column, err = strconv.Atoi(in.Text())
			if !isLinesColumnsOk(column, err, currBoard) {
				column = 0
				continue
			}
		}
	}
	line-- //Decrementing to make it match with the true indexes
	column--
	currBoard.UpdateBox(line, column, currPlayer)
	return currBoard
}

func isLinesColumnsOk(data int, err error, b Board) bool {
	if err != nil {
		fmt.Println("Line must be an integer.")
		return false
	}
	if data > len(b.boxes) || data < 1 {
		fmt.Printf("Line must be included between 1 and %d\n", len(b.boxes))
		return false
	}
	return true
}

func isOccupied(line int, column int, board Board) bool {
	if line == 0 && column == 0 {
		return true
	}
	if board.boxes[line-1][column-1] == ' ' {
		return false
	}
	fmt.Println("This box is already used.")
	return true
}

func switchPlayer(current Player) Player {
	if current == Player1 {
		return Player2
	} else {
		return Player1
	}
}

func isGameOver(b Board) bool {
	//This function is used to find out if a player won.
	//We use loops instead of if statements because in that way the code do not depend on the size of the double array
	for i := 0; i < len(b.boxes); i++ {
		testLine, testCol, testLeftDiag, testRightDiag := true, true, true, true //A variable for each case
		for j := 1; j < len(b.boxes); j++ {
			testLine = testLine && b.boxes[i][j-1] == b.boxes[i][j] && b.boxes[i][j] != ' '
			testCol = testCol && (b.boxes[j-1][i] == b.boxes[j][i] && b.boxes[j][i] != ' ')
			testLeftDiag = testLeftDiag && (b.boxes[j-1][j-1] == b.boxes[j][j] && b.boxes[j][j] != ' ')
			testRightDiag = testRightDiag && (b.boxes[j-1][(len(b.boxes)-1)-(j-1)] == b.boxes[j][(len(b.boxes)-1)-j] && b.boxes[j][(len(b.boxes)-1)-j] != ' ')
		}
		if testLine || testCol || testLeftDiag || testRightDiag {
			return true
		}
	}
	return false
}
