/*
 * @lc app=leetcode.cn id=581 lang=golang
 * @lcpr version=30200
 *
 * [581] 最短无序连续子数组
 */
package main

// @lc code=start
func findUnsortedSubarray(nums []int) int {
	// 在一个序列中, 有一段落是无序的, 一般思路是, 行开头找, 找到第一个逆序的和最后一个逆序的, 之后找到其中的最大值和最小值, 之后最接近最大最小值的索引,计算出长度
	// 有一种优化的思路, 找到最后一个低于当前遍历过的值的(低于就记录下来, 改与就更新最大值), 最开始的一个大于从后往前遍历的 minNums(高于就记录下来, 低于就更新值)
	// 最后就得到了正确答案
	n := len(nums)
	if n < 2 {
		return 0
	}
	maxNums := nums[0]
	end := -2
	for i:=1; i<n; i++{
		if nums[i] < maxNums {
			end = i
		} else {
			maxNums = nums[i]
		}
	}
	minNums := nums[n-1]
	start :=-1
	for i:=n-2; i>=0; i--{
		if nums[i] > minNums {
			start = i
		} else {
			minNums = nums[i]
		}
	}
	return end-start+1



	// n := len(nums)
	// if n < 2 {
	// 	return 0
	// }

	// start, end := -1, -2 // 保证已排序时返回0
	// maxVal, minVal := nums[0], nums[n-1]

	// // 正向遍历，寻找右边界
	// for i := 0; i < n; i++ {
	// 	if nums[i] < maxVal {
	// 		end = i
	// 	} else {
	// 		maxVal = nums[i]
	// 	}
	// }

	// // 反向遍历，寻找左边界
	// for i := n - 1; i >= 0; i-- {
	// 	if nums[i] > minVal {
	// 		start = i
	// 	} else {
	// 		minVal = nums[i]
	// 	}
	// }

	// return end - start + 1
}

// @lc code=end

/*
// @lcpr case=start
// [1, 3, 2, 2, 2]\n
// @lcpr case=end

// @lcpr case=start
// [2,1]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/

