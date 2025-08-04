/*
 * @lc app=leetcode.cn id=394 lang=golang
 * @lcpr version=30104
 *
 * [394] 字符串解码
 */
package main

import (
	"strconv"
	"strings"
)

// @lc code=start
func decodeString(s string) string {
	// 使用栈
	stack := make([]rune, 0, len(s))
	for _, ch := range s {
		if ch == ']' {
			// 先栈里面吐出字符
			currStr := []rune{}
			top := stack[len(stack)-1]

			for top >= 'a' && top <= 'z' {
				currStr = append(currStr, top)

				stack = stack[:len(stack)-1]
				top = stack[len(stack)-1]
			}
			reverseSlice(currStr)

			// 吐出 [
			if top == '[' {
				stack = stack[:len(stack)-1]
			}

			currCount := []rune{}
			top = stack[len(stack)-1]
			for top >= '0' && top <= '9' {
				currCount = append(currCount, top)
				stack = stack[:len(stack)-1]
				// 到末尾结束了
				if len(stack) == 0 {
					break
				}
				top = stack[len(stack)-1]
			}

			// 讲解压后的字符加到结果中
			reverseSlice(currCount)

			i, _ := strconv.Atoi(string(currCount))

			// 重点: 得到结果后, 讲数据写入到栈里面, 最后的所有结果都在栈里面
			for _, ch := range strings.Repeat(string(currStr), i) {
				stack = append(stack, ch)
			}
		} else {
			stack = append(stack, ch)
		}
	}

	remind := []rune{}
	for len(stack) != 0 {
		remind = append(remind, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	// 讲解压后的字符加到结果中
	reverseSlice(remind)

	return string(remind)
}

func reverseSlice(s []rune) {
	for i := 0; i < len(s)>>1; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

// @lc code=end

/*
// @lcpr case=start
// "3[a]2[bc]"\n
// @lcpr case=end

// @lcpr case=start
// "3[a2[c]]"\n
// @lcpr case=end

// @lcpr case=start
// "2[abc]3[cd]ef"\n
// @lcpr case=end

// @lcpr case=start
// "abc3[cd]xyz"\n
// @lcpr case=end

*/
