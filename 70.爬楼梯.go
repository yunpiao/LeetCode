/*
 * @lc app=leetcode.cn id=70 lang=golang
 * @lcpr version=30100
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	// if n <= 1 { return n }
	// dp1, dp2 := 1, 1
	// 精简二维 dp 后的版本
	// for i:=2; i<n+1; i++ {
	// 	dp2, dp1 = dp1 + dp2, dp2
	// }
	// return dp2
	// 排列问题, 使用 1 和 2 填充 n 有多少种排列方式
	// 完全背包, 填满这个背包, 有多少种排列方法
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

