/*
 * @lc app=leetcode.cn id=1 lang=golang
 * @lcpr version=30002
 *
 * [1] 两数之和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func twoSum(nums []int, target int) []int {
	// 目标, nums 里面找出两个元素, 加起来可以等于 target
	// 第一种方法 暴力破解, 两个 for 循环
	// for i:=0; i<len(nums)-1; i++{
	// 	for j:=i+1; j<len(nums); j++{
	// 		if nums[i] + nums[j] == target {
	// 			return []int{i, j}
	// 		}
	// 	}
	// }
	// 第二种方法, 优化下, 空间换时间, 可以讲查找过程使用 map 数据结构替换
	cache := make(map[int]int, len(nums))
	for i:=0; i<len(nums); i++{
		// 先查找之前有没有遇到过这个数据, 没有遇到过加入到 cache 里面, 遇到了就找到了答案
		if retIndex, exist := cache[target-nums[i]]; exist {
			return []int{retIndex, i}
		}
		cache[nums[i]] = i
	}

	return []int{}
}

func twoSumLast(nums []int, target int) []int {
	// 需要一个 hash 表判断是不是当前需要的数据
	set := make(map[int]int, len(nums))
	for i, n := range nums {
		if j, exist := set[n]; exist {
			return []int{i, j}
		} else {
			set[target-n] = i
		}
	}
	return []int{}
}

// @lc code=end

/*
// @lcpr case=start
// [2,7,11,15]\n9\n
// @lcpr case=end

// @lcpr case=start
// [3,2,4]\n6\n
// @lcpr case=end

// @lcpr case=start
// [3,3]\n6\n
// @lcpr case=end

*/

