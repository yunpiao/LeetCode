/*
 * @lc app=leetcode.cn id=32 lang=golang
 * @lcpr version=30104
 *
 * [32] 最长有效括号
 */

// @lc code=start
func longestValidParentheses(s string) int {
	// 核心 栈, 但是栈里面的元素变为了索引
	stack := make([]int, 0, len(s))
	stack = append(stack, -1) // 上一次出现错误的栈是 -1 的索引, 用于分界线

	ret := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			// 如果只有分界线, 更新分界线
			if len(stack) == 1 {
				stack[0] = i
			} else {

				stack = stack[:len(stack)-1]
				// 关键点
				ret = max(ret, i-stack[len(stack)-1]) // 使用在栈里面还未配对的索引, 可能是分切线, 也可能是之前还没有配对
			}
		} else {
			stack = append(stack, i)
		}
	}
	return ret
}

// @lc code=end

/*
// @lcpr case=start
// "(()"\n
// @lcpr case=end

// @lcpr case=start
// ")()())"\n
// @lcpr case=end

// @lcpr case=start
// ""\n
// @lcpr case=end

*/

