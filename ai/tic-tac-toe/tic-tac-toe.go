package tic_tac_toe

import "fmt"

var playerRune uint8
var enemyRune uint8

func main() {
	var player string
	fmt.Scan(&player)

	if player == "X" {
		playerRune = 'X'
		enemyRune = 'O'
	} else {
		playerRune = 'O'
		enemyRune = 'X'
	}

	board := make([]string, 3)
	for i := range board {
		fmt.Scan(&board[i])
	}

	nextMove(player, board)
}

func nextMove(player string, board []string) (result string) {
	if hasEmptyCenter(board) {
		result = "1 1"
	} else if hasLastEmptyCell(board) {
		result = findEmptyCell(board)
	} else {
		weights := createEmptyWeights()

		weights = calcWeightOfRows(board, weights)
		weights = calcWeightOfColumns(board, weights)
		weights = calcWeightOfRowsForEnemy(board, weights)
		weights = calcWeightOfColumnsForEnemy(board, weights)
		weights = calcWeightOfRowsForPlayer(board, weights)
		weights = calcWeightOfColumnsForPlayer(board, weights)
		// fmt.Println(weights) // debug

		if isAllZeroWeights(weights) {
			result = findEmptyCell(board)
		} else {
			result = findMaxWeight(weights)
		}
	}

	fmt.Println(result)

	return result
}

func createEmptyWeights() (weights []int) {
	weights = make([]int, 9)
	for i := range weights {
		weights[i] = 0
	}

	return weights
}

func hasEmptyCenter(board []string) bool {
	return board[1][1] == '_'
}

func findMaxWeight(weights []int) string {
	max := 0
	maxIndex := 0
	for i, weight := range weights {
		if weight > max {
			max = weight
			maxIndex = i
		}
	}

	return fmt.Sprintf("%d %d", maxIndex/3, maxIndex%3)
}

func calcWeightOfRows(board []string, weights []int) []int {
	for rowIndex := 0; rowIndex < 3; rowIndex++ {
		for colIndex := 0; colIndex < 3; colIndex++ {
			if board[rowIndex][colIndex] == enemyRune {
				for i := 0; i < 3; i++ {
					weights[rowIndex*3+i] = 0
				}
				break
			}

			if board[rowIndex][colIndex] == '_' {
				for i := 0; i < 3; i++ {
					weights[rowIndex*3+i] += 1
				}
			}

			if board[rowIndex][colIndex] == playerRune {
				for i := 0; i < 3; i++ {
					weights[rowIndex*3+i] += 2
				}
				weights[rowIndex*3+colIndex] = 0
			}
		}
	}

	return weights
}

func calcWeightOfColumns(board []string, weights []int) []int {
	for colIndex := 0; colIndex < 3; colIndex++ {
		for rowIndex := 0; rowIndex < 3; rowIndex++ {
			if board[rowIndex][colIndex] == enemyRune {
				for i := 0; i < 3; i++ {
					weights[i*3+colIndex] = 0
				}
				break
			}

			if board[rowIndex][colIndex] == '_' {
				for i := 0; i < 3; i++ {
					weights[i*3+colIndex] += 1
				}
			}

			if board[rowIndex][colIndex] == playerRune {
				for i := 0; i < 3; i++ {
					weights[i*3+colIndex] += 2
				}
				weights[rowIndex+colIndex] = 0
			}
		}
	}

	return weights
}

func calcWeightOfRowsForEnemy(board []string, weights []int) []int {
	for rowIndex := 0; rowIndex < 3; rowIndex++ {
		freeColumnIndex := -1
		totalEnemies := 0

		for colIndex := 0; colIndex < 3; colIndex++ {
			if board[rowIndex][colIndex] == enemyRune {
				totalEnemies++
			}

			if board[rowIndex][colIndex] == '_' {
				freeColumnIndex = colIndex
			}

			if board[rowIndex][colIndex] == playerRune {
				weights[rowIndex*3+colIndex] = 0
			}
		}

		if freeColumnIndex != -1 && totalEnemies == 2 {
			weights[rowIndex*3+freeColumnIndex] += 10
		}
	}

	return weights
}

func calcWeightOfColumnsForEnemy(board []string, weights []int) []int {
	for colIndex := 0; colIndex < 3; colIndex++ {
		freeRowIndex := -1
		totalEnemies := 0

		for rowIndex := 0; rowIndex < 3; rowIndex++ {
			if board[rowIndex][colIndex] == enemyRune {
				totalEnemies++
			}

			if board[rowIndex][colIndex] == '_' {
				freeRowIndex = rowIndex
			}

			if board[rowIndex][colIndex] == playerRune {
				weights[rowIndex*3+colIndex] = 0
			}
		}

		if freeRowIndex != -1 && totalEnemies == 2 {
			weights[freeRowIndex*3+colIndex] += 10
		}
	}

	return weights
}

func calcWeightOfRowsForPlayer(board []string, weights []int) []int {
	for rowIndex := 0; rowIndex < 3; rowIndex++ {
		freeColumnIndex := -1
		totalPlayers := 0

		for colIndex := 0; colIndex < 3; colIndex++ {
			if board[rowIndex][colIndex] == playerRune {
				totalPlayers++
			}

			if board[rowIndex][colIndex] == '_' {
				freeColumnIndex = colIndex
			}

			if board[rowIndex][colIndex] == enemyRune {
				weights[rowIndex*3+colIndex] = 0
			}
		}

		if freeColumnIndex != -1 && totalPlayers == 2 {
			weights[rowIndex*3+freeColumnIndex] += 20
		}
	}

	return weights
}

func calcWeightOfColumnsForPlayer(board []string, weights []int) []int {
	for colIndex := 0; colIndex < 3; colIndex++ {
		freeRowIndex := -1
		totalPlayers := 0

		for rowIndex := 0; rowIndex < 3; rowIndex++ {
			if board[rowIndex][colIndex] == playerRune {
				totalPlayers++
			}

			if board[rowIndex][colIndex] == '_' {
				freeRowIndex = rowIndex
			}

			if board[rowIndex][colIndex] == enemyRune {
				weights[rowIndex*3+colIndex] = 0
			}
		}

		if freeRowIndex != -1 && totalPlayers == 2 {
			weights[freeRowIndex*3+colIndex] += 20
		}
	}

	return weights
}

func hasLastEmptyCell(board []string) bool {
	numberOfEmpties := 0
	for _, row := range board {
		for _, r := range row {
			if r == '_' {
				numberOfEmpties++
			}
		}
	}

	return numberOfEmpties <= 1
}

func findEmptyCell(board []string) string {
	for rowIndex, row := range board {
		for columnIndex, r := range row {
			if r == '_' {
				return fmt.Sprintf("%d %d", rowIndex, columnIndex)
			}
		}
	}

	return ""
}

func isAllZeroWeights(weights []int) bool {
	for _, weight := range weights {
		if weight != 0 {
			return false
		}
	}

	return true
}
