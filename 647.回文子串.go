/*
 * @lc app=leetcode.cn id=647 lang=golang
 * @lcpr version=30104
 *
 * [647] 回文子串
 */

// @lc code=start
func countSubstrings(s string) int {

	// 	还有更简单的 双指针法则
	// 从中心向两边扩散
	ret := 0

	countString := func(s string, left int, right int) int {
		ret := 0
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
			ret++
		}
		return ret
	}

	for i := 0; i < len(s); i++ {
		// 奇数情况下
		ret += countString(s, i, i)
		// 偶数情况下
		ret += countString(s, i, i+1)
	}

	return ret

	// dp := make([][]bool, len(s))
	// for i := range dp {
	// 	dp[i] = make([]bool, len(s))
	// }

	// ret := 0

	// for i := len(s) - 1; i >= 0; i-- {
	// 	for j := i; j < len(s); j++ {
	// 		if i == j {
	// 			dp[i][j] = true // 单个字符是回文串
	// 		} else if s[i] == s[j] {
	// 			if j-i <= 1 {
	// 				dp[i][j] = true // 两个字符是回文串
	// 			} else {
	// 				dp[i][j] = dp[i+1][j-1] // 多个字符是回文串
	// 			}
	// 		}
	// 		if dp[i][j] {
	// 			ret++
	// 		}
	// 	}
	// }

	// return ret
}

// @lc code=end

/*
// @lcpr case=start
// "fdsklf"\n
// @lcpr case=end

// @lcpr case=start
// "aaa"\n
// @lcpr case=end

*/

