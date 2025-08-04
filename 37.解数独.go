/*
 * @lc app=leetcode.cn id=37 lang=golang
 * @lcpr version=30104
 *
 * [37] 解数独
 */

// @lc code=start
func solveSudoku(board [][]byte) {
	var dfs func(board [][]byte) bool
	dfs = func(board [][]byte) bool {
		for row := 0; row < 9; row++ {
			for col := 0; col < 9; col++ {
				if board[row][col] == '.' {
					for c := '1'; c <= '9'; c++ {
						if isValid(board, row, col, byte(c)) {
							board[row][col] = byte(c)
							if dfs(board) {
								return true
							}
							board[row][col] = '.'
						}
					}
					return false
				}
			}
		}
		return true
	}
	dfs(board)
}

func isValid(board [][]byte, row int, col int, c byte) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] == c {
			return false
		}
		if board[row][i] == c {
			return false
		}
	}
	rowStart := (row / 3) * 3
	colStart := (col / 3) * 3
	for i := rowStart; i < rowStart+3; i++ {
		for j := colStart; j < colStart+3; j++ {
			if board[i][j] == c {
				return false
			}
		}
	}
	return true
}

// @lc code=end

/*
// @lcpr case=start
// [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]\n
// @lcpr case=end

*/

