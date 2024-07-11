package main

import (
	"fmt"
	"strconv"
)

func isValidSudoku(board [][]byte) bool {
	return checkBox(board) && checkCol(board) && checkRow(board)
}

func checkRow(board [][]byte) bool {
	for col := 0; col < len(board); col++ {
		numsMap := map[int]bool{}
		for row := 0; row < len(board[col]); row++ {
			num := byteToNum(board[col][row])
			if string(board[col][row]) != "." && num > 9 && num < 1 && numsMap[num] {
				return false
			}

			if string(board[col][row]) != "." {
				numsMap[num] = true
			}
		}
	}

	return true
}

func checkCol(board [][]byte) bool {
	for row := 0; row < len(board); row++ {
		numsMap := map[int]bool{}
		for col := 0; col < len(board[row]); col++ {
			num := byteToNum(board[col][row])
			if string(board[col][row]) != "." && num > 9 && num < 1 && numsMap[num] {
				return false
			}

			if string(board[col][row]) != "." {
				numsMap[num] = true
			}
		}
	}

	return true
}

func checkBox(board [][]byte) bool {
	for col := 1; col < len(board); {
		for row := 1; row < len(board); {
			if row%4 == 0 && col%3 == 0 {
				col -= 2
			} else if row%4 != 0 && col%4 != 0 {
				fmt.Print(string(board[col-1][row-1]))
				row++
			} else if row%4 == 0 && col%4 != 0 {
				row -= 3
				col++
				fmt.Println()
			} else {
				break
			}
		}
	}

	return true
}

func byteToNum(b byte) int {
	num, _ := strconv.Atoi(string(b))

	return num
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	res := checkBox(board)

	fmt.Println(res)
}
