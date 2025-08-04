/*
 * @lc app=leetcode.cn id=115 lang=golang
 * @lcpr version=30104
 *
 * [115] 不同的子序列
 */

// @lc code=start
func numDistinct(s string, t string) int {
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
		dp[i][0] = 1
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j] // 使用 i-1 (dp[i-1][j-1]) 和 不使用 i-1(dp[i-1][j] 前 i-1 字符)
			} else {
				dp[i][j] = dp[i-1][j] // 不相等 只删除 s 最后一个的结果
			}
		}
	}
	return dp[len(s)][len(t)]
}

// @lc code=end

/*
// @lcpr case=start
// "rabbbit"\n"rabbit"\n
// @lcpr case=end

// @lcpr case=start
// "babgbag"\n"bag"\n
// @lcpr case=end

*/

