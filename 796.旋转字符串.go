/*
 * @lc app=leetcode.cn id=796 lang=golang
 * @lcpr version=30104
 *
 * [796] 旋转字符串
 */
package main

import "strings"

// @lc code=start
func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if s == goal {
		return true
	}
	ss := s + s
	return strings.Contains(ss, goal)

}

// @lc code=end

/*
// @lcpr case=start
// "abcde"\n"cdeab"\n
// @lcpr case=end

// @lcpr case=start
// "abcde"\n"abced"\n
// @lcpr case=end

*/
