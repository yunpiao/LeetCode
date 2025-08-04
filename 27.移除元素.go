/*
 * @lc app=leetcode.cn id=27 lang=golang
 * @lcpr version=30104
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	// 快慢指针
	n := len(nums)
	slow := 0
	for fast := 0; fast < n; fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast] // 复制就复制, 有分支预测, 如果出现 if, 可能会分支预测失败, 不要过早优化
			slow++
		}
	}
	return slow

	//双指针

	// n := len(nums) - 1
	// start, end := 0, n

	// for start <= end { // 需要 =, 让 start到达 val 开始的位置
	// 	if nums[start] != val {
	// 		start++
	// 		continue
	// 	}
	// 	if nums[end] == val {
	// 		end--
	// 		continue
	// 	}
	// 	nums[start], nums[end] = nums[end], nums[start]
	// 	start++
	// 	end--
	// }
	// return start
}

// @lc code=end

/*
// @lcpr case=start
// [3,2,2,3]\n3\n
// @lcpr case=end

// @lcpr case=start
// [0,1,2,2,3,0,4,2]\n2\n
// @lcpr case=end

*/

