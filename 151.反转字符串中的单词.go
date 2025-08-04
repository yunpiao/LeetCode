/*
 * @lc app=leetcode.cn id=151 lang=golang
 * @lcpr version=30104
 *
 * [151] 反转字符串中的单词
 */
package main

// @lc code=start
func reverseWords(s string) string {
	// 不使用额外空间
	// 1. 首先去除对于空格
	// 2. 接下来反转每一个单词
	// 3. 反转整个字符串
	bytes := []byte(s)
	if len(bytes) < 2 {
		return s
	}
	count := 0
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == ' ' {
			count++
		} else {
			break
		}
	}

	bytes = bytes[count:]
	count = 0

	for i := len(bytes) - 1; i >= 0; i-- {
		if bytes[i] == ' ' {
			count++
		} else {
			break
		}
	}
	bytes = bytes[:len(bytes)-count]

	if len(bytes) < 2 {
		return string(bytes)
	}

	fast, slow := 1, 1
	for fast < len(bytes) {
		for fast < len(bytes) && bytes[fast-1] == bytes[fast] && bytes[fast] == ' ' {
			fast++
		}
		bytes[slow] = bytes[fast]
		slow++
		fast++
	}
	bytes = bytes[:slow]

	// 反转每一个字符
	reverseStr(bytes)
	// 反转每一个单词
	start, end := 0, 0
	for end <= len(bytes) {
		if end == len(bytes) || bytes[end] == ' ' {
			reverseStr(bytes[start:end])
			start = end + 1
		}
		end++
	}

	return string(bytes)

}

func reverseStr(bytes []byte) {
	start, end := 0, len(bytes)-1
	for start < end {
		bytes[start], bytes[end] = bytes[end], bytes[start]
		start++
		end--
	}
}

// @lc code=end

/*
// @lcpr case=start
// "the sky is blue"\n
// @lcpr case=end

// @lcpr case=start
// "  hello world  "\n
// @lcpr case=end

// @lcpr case=start
// "a good   example"\n
// @lcpr case=end

*/
