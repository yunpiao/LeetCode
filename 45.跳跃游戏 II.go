/*
 * @lc app=leetcode.cn id=45 lang=golang
 * @lcpr version=30104
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 采用两个两个变量记录
	// 一个记录当前层级能达到的最远距离-每次都更新
	// 一个记录当前层级能达到的最远距离-只有到达最远距离的时候更新
	count := 0
	currMax := 0  // 每次都更新, 用于在到达下一层级别的时候,标记当前层级能达到的最远距离
	levelMax := 0 // 只有到达最远距离的时候更新, 用于标记当前层级能达到的最远距离
	for i := 0; i < len(nums)-1; i++ {
		currMax = max(currMax, i+nums[i])

		// 关键点, 到达当前层级最远距离的时候, 更新下一层级最远距离
		if i == levelMax {
			levelMax = currMax
			count++
			// 关键点, 如果当前层级最远距离已经大于等于数组长度, 则直接返回
			if currMax >= len(nums)-1 {
				return count
			}
		}
	}
	return count

}

// @lc code=end

/*
// @lcpr case=start
// [2,3,1,1,4]\n
// @lcpr case=end

// @lcpr case=start
// [2,3,0,1,4]\n
// @lcpr case=end

*/

