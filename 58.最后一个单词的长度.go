/*
 * @lc app=leetcode.cn id=58 lang=golang
 * @lcpr version=30104
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	n := len(s)
	ret := 0
	for i:=n-1; i>=0; i-- {
		if s[i] == ' ' {
			if ret == 0 {
				continue
			} else {
				break
			}
		}
		ret++
	}
	return ret
}

// @lc code=end

/*
// @lcpr case=start
// "Hello World"\n
// @lcpr case=end

// @lcpr case=start
// "   fly me   to   the moon  "\n
// @lcpr case=end

// @lcpr case=start
// "luffy is still joyboy"\n
// @lcpr case=end

*/

