/*
 * @lc app=leetcode.cn id=22 lang=golang
 * @lcpr version=30005
 *
 * [22] 括号生成
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func generateParenthesis(n int) []string {
	// 边界值判断
	if n == 0 { return []string{}}
	ret := []string{}
	var backtrack func(path string, left, right int)
	backtrack = func(path string, left, right int) {
		if len(path) == 2*n {
			ret = append(ret, path)
		}
		// 如果还能加左边的括号, 优先加左边的括号
		if left < n {
			backtrack(path+"(", left+1, right)
		}
		
		// 如果可以加右边的括号, 左边大于右边的数量, 则加右边的数量
		if right < left {
			backtrack(path+")", left, right+1)
		}
	}
	backtrack("", 0, 0)
	return ret
}

// @lc code=end

/*
// @lcpr case=start
// 3\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/

