/*
 * @lc app=leetcode.cn id=20 lang=golang
 * @lcpr version=30005
 *
 * [20] 有效的括号
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func isValid(s string) bool {
	// 边界值判断
	if len(s) == 0 {
		return true
	}
	validMap := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if v, exist := validMap[s[i]]; exist {
			if len(stack) > 0 && stack[len(stack)-1] == v {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0

}

// @lc code=end

/*
// @lcpr case=start
// "()"\n
// @lcpr case=end

// @lcpr case=start
// "()[]{}"\n
// @lcpr case=end

// @lcpr case=start
// "(]"\n
// @lcpr case=end

// @lcpr case=start
// "([])"\n
// @lcpr case=end

*/

