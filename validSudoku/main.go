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
			if string(board[col][row]) != "." && (num > 9 || num < 1 || numsMap[num]) {
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
			if string(board[col][row]) != "." && (num > 9 || num < 1 || numsMap[num]) {
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
	for col := 1; col <= len(board); {
		numsMap := map[int]bool{}
		for row := 1; ; {
			num := byteToNum(board[col-1][row-1])
			if string(board[col-1][row-1]) != "." && (num > 9 || num < 1 || numsMap[num]) {
				return false
			}

			if string(board[col-1][row-1]) != "." {
				numsMap[num] = true
			}

			if row%3 == 0 && col%3 == 0 {
				numsMap = map[int]bool{}
			}
			if row == len(board) && col == len(board) {
				col++
				break
			} else if col%9 == 0 && row%3 == 0 {
				col -= 8
				row++
			} else if row%3 == 0 {
				row -= 2
				col++
			} else {
				row++
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
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	res := isValidSudoku(board)

	fmt.Println(res)
}
