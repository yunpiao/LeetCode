/*
 * @lc app=leetcode.cn id=796 lang=golang
 * @lcpr version=30104
 *
 * [796] 旋转字符串
 */
package main

import "fmt"

// @lc code=start
func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if s == goal {
		return true
	}
	for i:=0; i<len(s); i++ {
		if s[i:len(s)] == goal[:len(s)-i] && 
		s[:i] == goal[len(s)-i:] {
			return true
		}
	}
	return false

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
