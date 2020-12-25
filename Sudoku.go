package main

import (
	"fmt"
)

func inRange(board [][]int) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] > 0 && board[i][j] <= 9 {
				return true
			}
		}
	}
	return false
}
func validSudoku(board [][]int) bool {
	if inRange(board) == false {
		return false
	}
	n := len(board)
	var differentRow [10]bool
	differentRow[0] = true
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			m := board[i][j]
			//fmt.Println(differentRow)
			if differentRow[m] == true && i == 9 {
				return false
			}
			differentRow[m] = true
		}
	}
	var differentCol [10]bool
	differentCol[0] = true
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			N := board[j][i]
			if differentCol[N] == true && i == 9 {
				return false
			}
			differentCol[N] = true
		}
	}
	var differentBox [10]bool
	differentBox[0] = true
	for i := 0; i < n-2; i += 3 {
		for j := 0; j < n-2; j += 3 {

			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					o := board[k+i][j+l]
					if differentBox[o] == true && i == 9 {
						return false
					}
					differentBox[o] = true
				}
			}
		}
	}
	return true
}

func main() {
	board := [][]int{[]int{7, 9, 2, 1, 5, 4, 3, 8, 6},
		[]int{6, 4, 3, 8, 2, 7, 1, 5, 9},
		[]int{8, 5, 1, 3, 9, 6, 7, 2, 4},
		[]int{2, 6, 5, 9, 7, 3, 8, 4, 1},
		[]int{4, 8, 9, 5, 6, 1, 2, 7, 3},
		[]int{3, 1, 7, 4, 8, 2, 9, 6, 5},
		[]int{1, 3, 6, 7, 4, 8, 5, 9, 2},
		[]int{9, 7, 4, 2, 1, 5, 6, 3, 8},
		[]int{5, 2, 8, 6, 3, 9, 4, 1, 7}}
	if validSudoku(board) {
		fmt.Println("valid")
	} else {
		fmt.Println("Invalid")
	}
}
