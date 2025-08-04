/*
 * @lc app=leetcode.cn id=135 lang=golang
 * @lcpr version=30104
 *
 * [135] 分发糖果
 */

// @lc code=start
func candy(ratings []int) int {
	// 局部最优, 只看当前和上一个, 如果当前比上一个大, 则当前的糖果数为上一个加1, 如果小, 则当前的糖果数为1
	candies := make([]int, len(ratings))
	for i := range candies {
		candies[i] = 1
	}

	for i := range ratings {
		if i == 0 {
			continue
		}
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}
	// 从后往前遍历, 如果当前比后一个大, 则当前的糖果数为后一个加1 和 当前的糖果数中较大的一个
	ret := candies[len(ratings)-1]
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			// 关键点, 这个点可能本身已经比前一个大, 所以需要取max
			//如果 ratings[i] > ratings[i + 1]，此时candyVec[i]（第i个小孩的糖果数量）就有两个选择了，一个是candyVec[i + 1] + 1（从右边这个加1得到的糖果数量），一个是candyVec[i]（之前比较右孩子大于左孩子得到的糖果数量）。
			// 那么又要贪心了，局部最优：取candyVec[i + 1] + 1 和 candyVec[i] 最大的糖果数量，保证第i个小孩的糖果数量既大于左边的也大于右边的。全局最优：相邻的孩子中，评分高的孩子获得更多的糖果。

			// 局部最优可以推出全局最优。

			// 所以就取candyVec[i + 1] + 1 和 candyVec[i] 最大的糖果数量，candyVec[i]只有取最大的才能既保持对左边candyVec[i - 1]的糖果多，也比右边candyVec[i + 1]的糖果多。
			candies[i] = max(candies[i], candies[i+1]+1)
		}
		ret += candies[i]
	}
	return ret
}

// @lc code=end

/*
// @lcpr case=start
// [1,0,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,2]\n
// @lcpr case=end

*/

