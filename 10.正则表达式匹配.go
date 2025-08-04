/*
 * @lc app=leetcode.cn id=10 lang=golang
 * @lcpr version=30104
 *
 * [10] 正则表达式匹配
 */
package main

// @lc code=start
func isMatch(s string, p string) bool {
	m := len(s)
	n := len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true

	for ni := 1; ni <= n; ni++ {
		if p[ni-1] == '*' && ni-2 >= 0 {
			// 如果当前是 *, 根据 * 前的数据来看
			dp[0][ni] = dp[0][ni-2]
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			pChar := p[j-1]
			sChar := s[i-1]
			if pChar == '.' || pChar == sChar {
				dp[i][j] = dp[i-1][j-1]
			} else if pChar == '*' && j >= 2 {
				ppChar := p[j-2]
				// 匹配 0 次
				dp[i][j] = dp[i][j-2]
				// 匹配多次
				if dp[i-1][j] && (ppChar == '.' || ppChar == sChar) {
					dp[i][j] = true
				}
			}
		}
	}
	return dp[m][n]

}

// @lc code=end

/*
// @lcpr case=start
// "aa"\n"a"\n
// @lcpr case=end

// @lcpr case=start
// "aa"\n"a*"\n
// @lcpr case=end

// @lcpr case=start
// "ab"\n".*"\n
// @lcpr case=end

*/
