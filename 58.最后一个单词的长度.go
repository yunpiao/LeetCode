/*
 * @lc app=leetcode.cn id=58 lang=golang
 * @lcpr version=30104
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	startSpace := true
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if startSpace && s[i] == ' ' {
			continue
		}
		startSpace = false
		if s[i] == ' ' {
			return count
		} else {
			count++
		}

	}
	return count

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

