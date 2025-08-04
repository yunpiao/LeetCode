/*
 * @lc app=leetcode.cn id=51 lang=golang
 * @lcpr version=30104
 *
 * [51] N 皇后
 */

// @lc code=start
package main

import "strings"

func solveNQueens(n int) [][]string {
	board := make([]string, n)
	for i :=0; i<n; i++{
		board[i] = strings.Repeat(".", n)
	}

	ret := [][]string{}

	var dfs func(board []string, row int)
	dfs = func(board []string, row int) {
		if row == n {
			tmp := make([]string, n)
			copy(tmp, board)
			ret = append(ret, tmp)
			return
		}
		for col :=0; col<n; col++{
			if isValid(board, row, col){
				board[row] = board[row][:col] + "Q" + board[row][col+1:]
				dfs(board, row+1)
				board[row] = board[row][:col] + "." + board[row][col+1:]
			}
		}
	}
	dfs(board, 0)
	return ret
	
}

func isValid(board []string, row int, col int) bool {
	//不能同一列
	for i := 0; i < len(board); i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	// 不能左斜上方
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	// 不能右斜上方
	for i, j := row-1, col+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	return true
}

// @lc code=end

/*
// @lcpr case=start
// 4\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/
