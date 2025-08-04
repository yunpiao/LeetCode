/*
 * @lc app=leetcode.cn id=150 lang=golang
 * @lcpr version=30104
 *
 * [150] 逆波兰表达式求值
 */

// @lc code=start
func evalRPN(tokens []string) int {
	stack := make([]string, 0, len(tokens))
	for i := 0; i < len(tokens); i++ {
		if len(stack) >= 2 && (tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/") {
			tmp := 0
			t1, _ := strconv.Atoi(stack[len(stack)-2])
			t2, _ := strconv.Atoi(stack[len(stack)-1])

			if tokens[i] == "+" {
				tmp = t1 + t2
			}
			if tokens[i] == "-" {
				tmp = t1 - t2
			}
			if tokens[i] == "*" {
				tmp = t1 * t2
			}
			if tokens[i] == "/" {
				tmp = t1 / t2
			}
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] = strconv.Itoa(tmp)
		} else {
			stack = append(stack, tokens[i])
		}
	}
	if len(stack) != 1 {
		return -1
	}

	ret, _ := strconv.Atoi(stack[0])
	return ret

}

// @lc code=end

/*
// @lcpr case=start
// ["2","1","+","3","*"]\n
// @lcpr case=end

// @lcpr case=start
// ["4","13","5","/","+"]\n
// @lcpr case=end

// @lcpr case=start
// ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]\n
// @lcpr case=end

*/

