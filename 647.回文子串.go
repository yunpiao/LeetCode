/*
 * @lc app=leetcode.cn id=647 lang=golang
 * @lcpr version=30104
 *
 * [647] 回文子串
 */
package main
// @lc code=start
func countSubstrings(s string) int {
	// 还是中心扩展法, 每次判断技术和偶数的两种情况
	n := len(s)
	ret := 0
	for i:=0; i<n; i++ {
		ret += _countSubstrings(s, i, i)
		ret += _countSubstrings(s, i, i+1)
	}
	return ret
}

func _countSubstrings(s string, start, end int) int {
	ret := 0
	for start>=0 && end<len(s) && s[start] == s[end] {
		start--
		end++
		ret++
	}
	return ret
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

