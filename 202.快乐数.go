/*
 * @lc app=leetcode.cn id=202 lang=golang
 * @lcpr version=30104
 *
 * [202] 快乐数
 */

// @lc code=start
func isHappy(n int) bool {
	// 用于判断一个数的每一个位置的平方加和
	help := func(n int) int {
		sum := 0
		for n != 0 {
			i := n % 10
			sum += i * i
			n = n / 10
		}
		return sum
	}
	// 用于避免循环
	hasSeen := make(map[int]struct{})

	// 判断快乐数
	for n != 1 {
		n = help(n)

		// 之前已经出现过, 代表循环了, 直接返回
		if _, exist := hasSeen[n]; exist {
			return false
		} else {
			hasSeen[n] = struct{}{}
		}
	}
	return true

}

// @lc code=end

/*
// @lcpr case=start
// 19\n
// @lcpr case=end

// @lcpr case=start
// 2\n
// @lcpr case=end

*/

