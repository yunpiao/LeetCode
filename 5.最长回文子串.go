/*
 * @lc app=leetcode.cn id=5 lang=golang
 * @lcpr version=30003
 *
 * [5] 最长回文子串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func longestPalindrome(s string) string {
	// 中心扩展法, 每次遍历都作为回文的中心节点去探索, 每次探索和记录值进行比对
	n := len(s)
	if n == 0 {
		return ""
	}
	ret := s[0:1]
	maxLength := 1

	// 闭包处理, 代码回更加清晰, 但是这种用法会破坏可读性
	huiwen := func(left, right int) {
		for left >= 0 && right < n && s[left] == s[right] {
			left--
			right++
		}

		if right-left+1 > maxLength {
			maxLength = right - left + 1
			ret = s[left+1 : right]
		}
	}

	for i := 0; i < n; i++ {
		huiwen(i, i)
		huiwen(i, i+1)
	}
	return ret
}

// @lc code=end

/*
// @lcpr case=start
// "babad"\n
// @lcpr case=end

// @lcpr case=start
// "cbbd"\n
// @lcpr case=end

*/

