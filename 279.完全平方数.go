/*
 * @lc app=leetcode.cn id=279 lang=golang
 * @lcpr version=30100
 *
 * [279] 完全平方数
 */

// @lc code=start
func numSquares(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	// 	初始化也很关键 求最小值, 要初始化很大的值
	for i := 0; i <= n; i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 0; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			// 递推, 需要其中遍历最小的
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]

}

// @lc code=end

/*
// @lcpr case=start
// 12\n
// @lcpr case=end

// @lcpr case=start
// 13\n
// @lcpr case=end

*/

