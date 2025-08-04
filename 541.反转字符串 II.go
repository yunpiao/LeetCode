/*
 * @lc app=leetcode.cn id=541 lang=golang
 * @lcpr version=30104
 *
 * [541] 反转字符串 II
 */

// @lc code=start
func reverseStr(s string, k int) string {
	n := len(s)
	bytes := []byte(s)

	for i := 0; i < n; i = i + 2*k {
		start := i
		// 	特殊场景在这里体现
		end := min(i+k-1, n-1)
		//双指针反转
		for start < end {
			bytes[start], bytes[end] = bytes[end], bytes[start]
			start++
			end--
		}
	}
	return string(bytes)

}

// @lc code=end

/*
// @lcpr case=start
// "abcdefg"\n2\n
// @lcpr case=end

// @lcpr case=start
// "abcd"\n2\n
// @lcpr case=end

*/

