/*
 * @lc app=leetcode.cn id=392 lang=golang
 * @lcpr version=30103
 *
 * [392] 判断子序列
 */

// @lc code=start
func isSubsequence(s string, t string) bool {

	// 可以使用最长公共子序列, 最后子序列长度是不是 len(s) 判断
	/*
		dp := make([][]int, len(s)+1)
		for i := range dp {
			dp[i] = make([]int, len(t)+1)
		}
		for i := 1; i <= len(s); i++ {
			for j := 1; j <= len(t); j++ {
				if s[i-1] == t[j-1] {
					dp[i][j] = dp[i-1][j-1] + 1
				} else {
					dp[i][j] = max(dp[i-1][j], dp[i][j-1])
				}
			}
		}
		return dp[len(s)][len(t)] == len(s)
	*/

	// 可以使用双指针, 一个指针指向 s, 一个指针指向 t, 如果 s[i] == t[j], 则 i++, j++, 最后判断 i 是否等于 len(s)

	/*
		i, j := 0, 0
		for i < len(s) && j < len(t) {
			if s[i] == t[j] {
				i++
				j++
			} else {
				j++
			}
		}
		return i == len(s)
	*/

	// 使用动态规划
	//dp[i][j] 表示以下标i-1为结尾的字符串s，和以下标j-1为结尾的字符串t，相同子序列的长度为dp[i][j]。
	// 状态转移方程:
	// 如果 s[i-1] == t[j-1] 则 dp[i][j] = dp[i-1][j-1] + 1
	// 否则 dp[i][j] = dp[i][j-1] 不匹配的情况下, 只需要 j-1 就可以, 也就是去掉 t 的最后一个字符
	// 初始化: dp[0][j] = 0, dp[i][0] = 0
	// 遍历顺序: 外层遍历 s, 内层遍历 t
	// 结果: dp[len(s)][len(t)] == len(s)

	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
	}
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// dp[i][j] = dp[i][j-1]
				dp[i][j] = max(dp[i-1][j], dp[i][j-1]) // 不匹配的情况下, 只需要 j-1 就可以, 也就是去掉 t 的最后一个字符

			}
		}
	}
	return dp[len(s)][len(t)] == len(s)
}

// @lc code=end

/*
// @lcpr case=start
// "abc"\n"ahbgdc"\n
// @lcpr case=end

// @lcpr case=start
// "axc"\n"ahbgdc"\n
// @lcpr case=end

*/

