/*
 * @lc app=leetcode.cn id=376 lang=golang
 * @lcpr version=30104
 *
 * [376] 摆动序列
 */

// @lc code=start
func wiggleMaxLength(nums []int) int {
	// 贪心, 这里使用的是贪心, 而不是动态规划

	n := len(nums)
	if n <= 1 {
		return n
	}

	// 记录上一个峰值
	preDiff := nums[1] - nums[0]
	count := 2
	if preDiff == 0 {
		count = 1
	}
	for i := 2; i < n; i++ {
		diff := nums[i] - nums[i-1]
		if (diff > 0 && preDiff <= 0) || (diff < 0 && preDiff >= 0) {
			preDiff = diff
			count++
		}
	}
	return count

	// // 使用两个 dp 变量, 来推出最大最小值

	// up := make([]int, len(nums))
	// down := make([]int, len(nums))
	// for i := range up {
	// 	up[i] = 1
	// 	down[i] = 1
	// }
	// maxUp := 1
	// maxDown := 1

	// for i := 1; i < len(nums); i++ {
	// 	if nums[i] > nums[i-1] {
	// 		up[i] = max(up[i], down[i-1]+1)
	// 	}
	// 	if nums[i] < nums[i-1] {
	// 		down[i] = max(down[i], up[i-1]+1)
	// 	}
	// 	maxUp = max(maxUp, up[i])
	// 	maxDown = max(maxDown, down[i])
	// }
	// return max(maxUp, maxDown)

	// n := len(nums)
	// if n <= 1 {
	// 	return n
	// }

	// l := 0
	// for i := 1; i < len(nums); i++ {
	// 	if nums[l] != nums[i] {
	// 		l++
	// 		nums[l] = nums[i]
	// 	}
	// }
	// if l == 0 {
	// 	return 1
	// }

	// for i := 0; i < l; i++ {
	// 	nums[i] = nums[i+1] - nums[i]
	// }

	// ret := 0

	// flag := nums[0] > 0
	// for i := 1; i < l; i++ {
	// 	if nums[i] > 0 && !flag {
	// 		ret++
	// 		flag = !flag
	// 	} else if nums[i] < 0 && flag {
	// 		ret++
	// 		flag = !flag
	// 	}
	// }
	// return ret + 2

}

// @lc code=end

/*
// @lcpr case=start
// [1,7,4,9,2,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,17,5,10,13,15,10,5,16,8]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5,6,7,8,9]\n
// @lcpr case=end

*/

