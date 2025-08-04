/*
 * @lc app=leetcode.cn id=459 lang=golang
 * @lcpr version=30104
 *
 * [459] 重复的子字符串
 */

// @lc code=start
func repeatedSubstringPattern(s string) bool {
	// 如果s是由重复的子字符串构成的，那么s+s中一定包含s
	// newS := s + s
	// newS = newS[1 : len(newS)-1]

	// return strings.Contains(newS, s)
	// 更加高效的算法, 利用 kmp 中的 next 数组

	next := make([]int, len(s))
	j := 0
	next[0] = 0
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] { // 重点, 如果这个不匹配, 就检查前一个前缀匹配
			j = next[j-1]
		}
		if s[i] == s[j] { // 如果匹配, 就更新 j
			j++
		}
		next[i] = j
	}
	n := len(s)
	len := n - next[n-1] // 重复串的长度
	return next[n-1] != 0 && n%(len) == 0
}

// @lc code=end

/*
// @lcpr case=start
// "abab"\n
// @lcpr case=end

// @lcpr case=start
// "aba"\n
// @lcpr case=end

// @lcpr case=start
// "abcabcabcabc"\n
// @lcpr case=end

*/

