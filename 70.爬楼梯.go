/*
 * @lc app=leetcode.cn id=70 lang=golang
 * @lcpr version=30100
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	// 动态规划, 完全背包问题, 0 1 可以无限的取, 同时求的是组合数量
	//
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for j := 2; j <= n; j++ {
		for i := 1; i <= 2; i++ {
			if j-i >= 0 {
				dp[j] = dp[j] + dp[j-i]
			}
		}
	}
	return dp[n]
}

// @lc code=end

/*
// @lcpr case=start
// 2\n
// @lcpr case=end

// @lcpr case=start
// 3\n
// @lcpr case=end

*/

