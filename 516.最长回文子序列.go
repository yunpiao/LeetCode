/*
 * @lc app=leetcode.cn id=516 lang=golang
 * @lcpr version=30104
 *
 * [516] 最长回文子序列
 */

// @lc code=start
func longestPalindromeSubseq(s string) int {
	// 中心拓展法失败
	// 动态规划
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
		for j := range dp[i] {
			dp[i][j] = 1
		}
	}

	for length := 2; length <= len(s); length++ {
		for start := 0; start <= len(s)-length; start++ {
			end := start + length - 1
			if s[start] == s[end] {
				if length == 2 {
					dp[start][end] = 2
				} else {
					dp[start][end] = dp[start+1][end-1] + 2
				}
			} else {
				dp[start][end] = max(dp[start+1][end], dp[start][end-1])
			}
		}

	}
	return dp[0][n-1]
}

// @lc code=end

/*
// @lcpr case=start
// "bbbab"\n
// @lcpr case=end

// @lcpr case=start
// "cbbd"\n
// @lcpr case=end

*/

