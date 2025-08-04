/*
 * @lc app=leetcode.cn id=344 lang=golang
 * @lcpr version=30104
 *
 * [344] 反转字符串
 */

// @lc code=start
func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}

}

// @lc code=end

/*
// @lcpr case=start
// ["h","e","l","l","o"]\n
// @lcpr case=end

// @lcpr case=start
// ["H","a","n","n","a","h"]\n
// @lcpr case=end

*/

