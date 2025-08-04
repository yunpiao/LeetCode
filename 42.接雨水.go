/*
 * @lc app=leetcode.cn id=42 lang=golang
 * @lcpr version=30104
 *
 * [42] 接雨水
 */
package main

// @lc code=start
func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	// 最简单的算法 暴力,每次找到index 最左侧和最右侧的最大值, 求出每个 index 的 min(lmax, rmax) - h[index] 面积
	// sum := 0
	// for i, h := range height {
	// 	if i == 0 || i == len(height)-1 {
	// 		continue
	// 	}
	// 	lMax := MaxInt(height[:i])
	// 	rMax := MaxInt(height[i+1:])
	// 	if lMax < h || rMax < h {
	// 		continue
	// 	} else {
	// 		sum += (min(lMax, rMax) - h)
	// 	}
	// }
	// return sum

	// 第二种, 优化下, 先记录下左右两边的最大值, 之后遍历的时候使用这部分直接计算面积
	// maxLeft := make([]int, len(height))
	// maxRight := make([]int, len(height))
	// for i := range maxLeft {
	// 	maxLeft[i] = -1
	// 	maxRight[i] = -1
	// }

	// maxLeft[0] = height[0]
	// for i := 1; i < len(height); i++ {
	// 	maxLeft[i] = max(height[i], maxLeft[i-1])
	// }

	// maxRight[len(height)-1] = height[len(height)-1]
	// for i := len(height) - 2; i >= 0; i-- {
	// 	maxRight[i] = max(height[i], maxRight[i+1])
	// }

	// sum := 0
	// for i, v := range height {
	// 	lmax := maxLeft[i]
	// 	rmax := maxRight[i]
	// 	minWater := min(lmax, rmax)
	// 	if minWater > height[i] {
	// 		sum += (minWater - v)
	// 	}
	// }
	// return sum
	// 以上都是每次都从列来计算的, 累加每一列的水量
	// 第三种是按照行来计算, 单调栈的方式 一旦发现添加的柱子高度大于栈头元素了，此时就出现凹槽了，栈头元素就是凹槽底部的柱子，栈头第二个元素就是凹槽左边的柱子，而添加的元素就是凹槽右边的柱子。
	// stack := make([]int, 0, len(height))
	// sum := 0
	// for i, v := range height {
	// 	for len(stack) > 0 && height[stack[len(stack)-1]] < v {
	// 		topIndex := stack[len(stack)-1]
	// 		stack = stack[:len(stack)-1]
	// 		// 每次有凹槽, 就可以计算面积, 面积根据为出栈和当前值中间的面积
	// 		if len(stack) > 0 {
	// 			leftIndex := height[stack[len(stack)-1]]          // 找到左边的面积
	// 			h := min(height[leftIndex], v) - height[topIndex] // 高是弹出后的栈顶索引 和 当前元素的较小值
	// 			w := i - stack[len(stack)-1] - 1                  // 宽是弹出后的栈顶索引到当前再-1
	// 			sum += h * w                                      // 面积 长x宽
	// 		}

	// 	}
	// 	stack = append(stack, i)
	// }
	// 第四种, 双指针,最优雅的解法, 还是按照列进行计算
	sum := 0
	l, r := 0, len(height)-1
	lm, rm := height[l], height[r]
	for l < r {
		// 处理左边
		// lm = max(lm, height[l])
		// rm = max(rm, height[r])

		// 重点, --- 在双指针移动过程中，至少有一个指针的当前高度是其遍历路径中的最大值
		if height[l] < height[r] { // height[r] 为当前遍历到的最大值
			if height[l] < lm {
				// 有凹槽, 左右两边都有比当前值大的数
				sum += lm - height[l] //   height[l] < lm <  height[r](局部最大值), 所以直接取 lm 而不是 min(lm, height[r])
			} else {
				lm = height[l]
			}
			l++
		} else { // height[l] 为当前遍历到的最大值
			if height[r] < rm {
				sum += rm - height[r] // height[i] 是扫描过的最大值, 所以直接 rm 就可以
			} else {
				rm = height[r]
			}
			r--
		}

	}
	return sum
}

// func MaxInt(nums []int) int {
// 	if len(nums) == 0 {
// 		return -1
// 	}
// 	ret := nums[0]
// 	for _, v := range nums {
// 		ret = max(ret, v)
// 	}
// 	return ret
// }

// @lc code=end

/*
// @lcpr case=start
// [0,1,0,2,1,0,1,3,2,1,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [4,2,0,3,2,5]\n
// @lcpr case=end

*/
