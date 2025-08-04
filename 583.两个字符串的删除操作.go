/*
 * @lc app=leetcode.cn id=583 lang=golang
 * @lcpr version=30104
 *
 * [583] 两个字符串的删除操作
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
	}
	for j := range dp[0] {
		dp[0][j] = j
	}

	// 删除的递推关系, 不相等可以由三个方向推出来, 取最小的就可以
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+2)
			}
		}
	}
	return dp[len(word1)][len(word2)]

}

// @lc code=end

/*
// @lcpr case=start
// "sea"\n"eat"\n
// @lcpr case=end

// @lcpr case=start
// "leetcode"\n"etco"\n
// @lcpr case=end

*/

