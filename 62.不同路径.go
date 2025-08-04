/*
 * @lc app=leetcode.cn id=62 lang=golang
 * @lcpr version=30104
 *
 * [62] 不同路径
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := range dp[0] {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if i-1 >= 0 && j-1 >= 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]

}

// @lc code=end

/*
// @lcpr case=start
// 3\n7\n
// @lcpr case=end

// @lcpr case=start
// 3\n2\n
// @lcpr case=end

// @lcpr case=start
// 7\n3\n
// @lcpr case=end

// @lcpr case=start
// 3\n3\n
// @lcpr case=end

*/

