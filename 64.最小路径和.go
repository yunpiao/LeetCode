/*
 * @lc app=leetcode.cn id=64 lang=golang
 * @lcpr version=30104
 *
 * [64] 最小路径和
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	sumRow := 0
	sumCol := 0
	for i := range dp {
		sumCol += grid[i][0]
		dp[i][0] = sumCol
	}
	for j := range dp[0] {
		sumRow += grid[0][j]
		dp[0][j] = sumRow
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
		}
	}
	return dp[m-1][n-1]

}

// @lc code=end

/*
// @lcpr case=start
// [[1,3,1],[1,5,1],[4,2,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,3],[4,5,6]]\n
// @lcpr case=end

*/

