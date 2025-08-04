/*
 * @lc app=leetcode.cn id=1143 lang=golang
 * @lcpr version=30103
 *
 * [1143] 最长公共子序列
 */

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}
	// 二维 dp
	dp := make([][]int, len(text1)+1)
	for i := 0; i < len(text1)+1; i++ {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// 这里的技巧是, 如果两个不相等,代表 t1 t2 其中的任意一个去掉之后可能相等,
				// 所以取 max(dp[i-1][j], dp[i][j-1]), 数据中, 就是将小于 ij 的任意数据的最大值
				// 将比 i j 小的区域中的最大值取出来
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}

	}
	return dp[len(text1)][len(text2)]

}

// @lc code=end

/*
// @lcpr case=start
// "abcde"\n"ace"\n
// @lcpr case=end

// @lcpr case=start
// "abc"\n"abc"\n
// @lcpr case=end

// @lcpr case=start
// "abc"\n"def"\n
// @lcpr case=end

*/

