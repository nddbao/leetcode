package solution

/*
	leetcode: https://leetcode.com/problems/range-sum-query-immutable/
*/

type NumArray struct {
	preSum []int
}

func Constructor(nums []int) NumArray {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i-1] + nums[i]
	}
	return NumArray{preSum: nums}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left-1 < 0 {
		return this.preSum[right]
	}

	return this.preSum[right] - this.preSum[left-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
