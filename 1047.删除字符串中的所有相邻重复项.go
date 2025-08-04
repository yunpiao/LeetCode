/*
 * @lc app=leetcode.cn id=1047 lang=golang
 * @lcpr version=30104
 *
 * [1047] 删除字符串中的所有相邻重复项
 */

// @lc code=start
func removeDuplicates(s string) string {
	stack := make([]rune, 0, len(s))
	for _, c := range s {
		if len(stack) > 0 && stack[len(stack)-1] == c {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}
	return string(stack)

}

// @lc code=end

/*
// @lcpr case=start
// "abbaca"\n
// @lcpr case=end

*/

