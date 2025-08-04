/*
 * @lc app=leetcode.cn id=221 lang=golang
 * @lcpr version=30104
 *
 * [221] 最大正方形
 */

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	ret := 0

	// 初始化dp数组
	// 如果matrix[i][j] == '1', 则dp[i][j] = 1
	// 否则dp[i][j] = 0
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[i]))
		for j, v := range matrix[i] {
			dp[i][j] = int(v - '0')
			ret = max(ret, dp[i][j])
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			// 遍历， 如果dp[i][j] != 0, 则dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1 关键点， 三个正方形才有可能凑成一个更大的正方形
			if dp[i][j] != 0 {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				ret = max(ret, dp[i][j])
				// 更新最大正方形边长
			}

		}
	}

	// 返回最大正方形的面积
	return ret * ret
}

// @lc code=end

/*
// @lcpr case=start
// [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]\n
// @lcpr case=end

// @lcpr case=start
// [["0","1"],["1","0"]]\n
// @lcpr case=end

// @lcpr case=start
// [["0"]]\n
// @lcpr case=end

*/

