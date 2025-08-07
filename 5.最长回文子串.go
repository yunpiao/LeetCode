/*
 * @lc app=leetcode.cn id=5 lang=golang
 * @lcpr version=30003
 *
 * [5] 最长回文子串
 */
package main

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func longestPalindrome(s string) string {
	// 目标, 找出最长的子串
	// 逐个遍历, 中心扩展, 找到最长的子串
	ret := ""
	for i:=0; i<len(s); i++{
		palindrome := getPalindrome(s, i, i)
		if len(palindrome) > len(ret) {
			ret = palindrome
		}
		palindrome = getPalindrome(s, i, i+1)
		if len(palindrome) > len(ret) {
			ret = palindrome
		}
	}
	return ret
}

// 找到最大的回文子串, 有两种情况, 中心为基数或者中心为偶数
func getPalindrome(s string, start, end int) string{
	if end >= len(s) {
		return ""
	}
	for start >= 0 && end < len(s) && s[start] == s[end]{
			start--
			end++
	}
	return s[start+1:end]
}

// @lc code=end

/*
// @lcpr case=start
// "aacabdkacaa"\n
// @lcpr case=end

// @lcpr case=start
// "cbbd"\n
// @lcpr case=end

*/

