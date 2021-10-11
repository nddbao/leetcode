package solution

/*
	leetcode: https://leetcode.com/problems/sliding-window-maximum/
*/

func maxSlidingWindow(nums []int, k int) []int {
	var result, qVal, qIdx []int
	for i := 0; i < len(nums); i++ {
		// remove nums[i-k]
		if i >= k && qIdx[0] == i-k {
			qVal = qVal[1:]
			qIdx = qIdx[1:]
		}

		// remove items that smalller  or equal to nums[i]
		for len(qVal) != 0 && qVal[len(qVal)-1] <= nums[i] {
			qVal = qVal[:len(qVal)-1]
			qIdx = qIdx[:len(qIdx)-1]
		}

		// add nums[i] to the end
		qVal = append(qVal, nums[i])
		qIdx = append(qIdx, i)

		if i >= k-1 {
			result = append(result, qVal[0])
		}
	}

	return result
}
