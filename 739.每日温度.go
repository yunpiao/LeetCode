/*
 * @lc app=leetcode.cn id=739 lang=golang
 * @lcpr version=30104
 *
 * [739] 每日温度
 */

// @lc code=start
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	stack := make([]int, 0, n)
	ret := make([]int, n)
	if n == 0 {
		return []int{}
	}

	// 单调栈, 每次都把小于当前的数据栈顶元素push 出来
	// 单调队列, 每次都把小于当前元素的队列头部 push 出来

	for i, temp := range temperatures {
		// 比当前小的全部出栈
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			index := stack[len(stack)-1]
			ret[index] = i - index
			stack = stack[:len(stack)-1]
		}

		// 入栈
		stack = append(stack, i)
	}
	return ret

}

// @lc code=end

/*
// @lcpr case=start
// [73,74,75,71,69,72,76,73]\n
// @lcpr case=end

// @lcpr case=start
// [30,40,50,60]\n
// @lcpr case=end

// @lcpr case=start
// [30,60,90]\n
// @lcpr case=end

*/

