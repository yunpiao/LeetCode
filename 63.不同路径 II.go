/*
 * @lc app=leetcode.cn id=63 lang=golang
 * @lcpr version=30104
 *
 * [63] 不同路径 II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])

	dp := make([]int, n)
	// 初始化
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	dp[0] = 1
	for i := 1; i < n; i++ {
		if obstacleGrid[0][i] != 0 {
			dp[i] = 0
		} else {
			dp[i] = dp[i-1]
		}
	}
	for i := 1; i < m; i++ {
		// 初始化第一列, 如果第一列就是 0, 代表数据不能继承, 需要重置
		if obstacleGrid[i][0] == 1 {
			dp[0] = 0
		}
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else {
				dp[j] = dp[j] + dp[j-1]
			}
		}
	}
	return dp[n-1]

}

// @lc code=end

/*
// @lcpr case=start
// [[0],[1]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,1],[0,0]]\n
// @lcpr case=end

*/

