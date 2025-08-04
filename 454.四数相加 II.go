/*
 * @lc app=leetcode.cn id=454 lang=golang
 * @lcpr version=30104
 *
 * [454] 四数相加 II
 */
package main

// @lc code=start
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	// 因为是 4 个, 所以可以使用两数之和类似的, 先同统计 nums1 nums2 的任意两数之和
	numsCount := make(map[int]int, len(nums1))
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			numsCount[n1+n2]++
		}
	}

	result := 0
	for _, n3 := range nums3 {
		for _, n4 := range nums4 {
			if v, exist := numsCount[-(n3 + n4)]; exist {
				result += v
			}
		}
	}

	return result
}

// @lc code=end

/*
// @lcpr case=start
// [1,2]\n[-2,-1]\n[-1,2]\n[0,2]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n[0]\n[0]\n[0]\n
// @lcpr case=end

*/
