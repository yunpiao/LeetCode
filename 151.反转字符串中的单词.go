/*
 * @lc app=leetcode.cn id=151 lang=golang
 * @lcpr version=30104
 *
 * [151] 反转字符串中的单词
 */
package main

import (
	"strings"
)

// @lc code=start
// func reverseWordsLast(s string) string {
// 	// 不使用额外空间
// 	// 1. 首先去除对于空格
// 	// 2. 接下来反转每一个单词
// 	// 3. 反转整个字符串
// 	bytes := []byte(s)
// 	if len(bytes) < 2 {
// 		return s
// 	}
// 	count := 0
// 	for i := 0; i < len(bytes); i++ {
// 		if bytes[i] == ' ' {
// 			count++
// 		} else {
// 			break
// 		}
// 	}

// 	bytes = bytes[count:]
// 	count = 0

// 	for i := len(bytes) - 1; i >= 0; i-- {
// 		if bytes[i] == ' ' {
// 			count++
// 		} else {
// 			break
// 		}
// 	}
// 	bytes = bytes[:len(bytes)-count]

// 	if len(bytes) < 2 {
// 		return string(bytes)
// 	}

// 	fast, slow := 1, 1
// 	for fast < len(bytes) {
// 		for fast < len(bytes) && bytes[fast-1] == bytes[fast] && bytes[fast] == ' ' {
// 			fast++
// 		}
// 		bytes[slow] = bytes[fast]
// 		slow++
// 		fast++
// 	}
// 	bytes = bytes[:slow]

// 	// 反转每一个字符
// 	reverseStr(bytes)
// 	// 反转每一个单词
// 	start, end := 0, 0
// 	for end <= len(bytes) {
// 		if end == len(bytes) || bytes[end] == ' ' {
// 			reverseStr(bytes[start:end])
// 			start = end + 1
// 		}
// 		end++
// 	}

// 	return string(bytes)

// }

//	func reverseStrLast(bytes []byte) {
//		start, end := 0, len(bytes)-1
//		for start < end {
//			bytes[start], bytes[end] = bytes[end], bytes[start]
//			start++
//			end--
//		}
//	}
func reverseWords(s string) string {
	// 目标, 按照空格为区分, 对每一个单词进行反转
	// 注意需要去掉收尾的空格

	// strs := []string{}
	// for {
	// 	index := strings.Index(s, " ")
	// 	if index == -1 {
	// 		break
	// 	}
	// 	if index == 0 {
	// 		s = s[1:]
	// 		continue
	// 	}

	// 	strs = append(strs, s[:index])
	// 	s = s[index+1:]
	// }
	// if s != "" {
	// 	strs = append(strs, s)
	// }
	index := 0
	strs := []string{}
	feildStart := index
	for index < len(s) {
		// 先跳过前导空格, 如果先找到单词开头就会有多种复杂情况, 比如在循环前去空格, 在循环结束后单独处理最后一个词
		for index < len(s) && s[index] == ' ' {
			index++
		}

		// 超过后直接退出
		if index >= len(s) {
			break
		}

		// 找到单词的首字
		feildStart = index
		for index < len(s) && s[index] != ' ' {
			index++
		}
		// 找到单词的结束
		strs = append(strs, s[feildStart:index])
	}

	// 逆序数组
	for i := 0; i < len(strs)/2; i++ {
		strs[i], strs[len(strs)-i-1] = strs[len(strs)-i-1], strs[i]
	}

	return strings.Join(strs, " ")
}

// @lc code=end

/*
// @lcpr case=start
// " the  sky is blue "\n
// @lcpr case=end

// @lcpr case=start
// "  hello world  "\n
// @lcpr case=end

// @lcpr case=start
// "a good   example"\n
// @lcpr case=end

*/
