/*
 * @lc app=leetcode.cn id=738 lang=golang
 * @lcpr version=30104
 *
 * [738] 单调递增的数字
 */
package main

import (
	"strconv"
)

// @lc code=start
func monotoneIncreasingDigits(n int) int {
	// 转成字符串处理
	c := strconv.Itoa(n)
	chars := []byte(c)

	mark := len(chars)
	for i := len(chars) - 2; i >= 0; i-- {
		if chars[i] > chars[i+1] {
			mark = i
			chars[i]--
		}
	}
	for i := mark + 1; i < len(chars); i++ {
		chars[i] = '9'
	}

	ret, _ := strconv.Atoi(string(chars))

	return ret

	// // 	1. 基本步骤, 如果 i-1 的数字 < i 的数字, i-1 位置--, i 位置等于 9
	// last := n % 10
	// n = n / 10

	// ret := 0

	// tmp := []int{last}
	// for n != 0 {
	// 	curr := n % 10
	// 	n = n / 10
	// 	if curr > last {
	// 		for i := range tmp {
	// 			tmp[i] = 9
	// 		}
	// 		tmp = append([]int{curr - 1}, tmp...)
	// 	} else {
	// 		tmp = append([]int{curr}, tmp...)
	// 		last = curr
	// 	}
	// }

	// for i := range tmp {
	// 	ret = ret*10 + tmp[i]
	// }
	// return ret
}

// @lc code=end

/*
// @lcpr case=start
// 10\n
// @lcpr case=end

// @lcpr case=start
// 1234\n
// @lcpr case=end

// @lcpr case=start
// 332\n
// @lcpr case=end

*/
