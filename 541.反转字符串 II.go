/*
 * @lc app=leetcode.cn id=541 lang=golang
 * @lcpr version=30104
 *
 * [541] 反转字符串 II
 */
package main

// @lc code=start
func reverseStr(s string, k int) string {
	buffer := []byte(s)

	n := len(buffer)

	// 单步 2k
	for i:=0; i<n; i=i+2*k {
		end := min(i+k, n) // 异常情况
		_reverseStr(buffer[i:end]) // 直接反转, 不要参数
	}
	return string(buffer)
}
// 单独的函数用于反转字符串
func _reverseStr(bytes []byte) {
	n := len(bytes)
	for i:=0; i<n/2; i++ {
		bytes[i], bytes[n-i-1] = bytes[n-i-1], bytes[i]
	}
}

// @lc code=end

/*
// @lcpr case=start
// "abcdefg"\n2\n
// @lcpr case=end

// @lcpr case=start
// "abcd"\n8\n
// @lcpr case=end

*/

