package solution

/*
	Similar to normal solution when using two pointer L and R.
	But in this range [L..R]:
		+ have k distinct
		+ nums[L]  value appear only and only one time.

	We need prefix to know how far we go from index.
	It means [(L-prefix)....L.... R] have k distinct
	==> We know that number of subArrayWithDistinct = prefix + 1

	Time complexity: O(n)
	Space complexity: O(n)

	// TODO try another way
*/
func subarraysWithKDistinctOptimize(nums []int, k int) int {
	var count, prefix int
	checker := NewChecker()
	l, r := 0, 0
	for ; r < len(nums); r++ {
		checker.Add(nums[r])
		d := checker.DistinctNumbers()
		if d > k {
			prefix = 0
			checker.Remove(nums[l])
			l++
		}

		for checker.GetCount(nums[l]) > 1 {
			checker.Remove(nums[l])
			prefix++
			l++
		}

		if checker.DistinctNumbers() == k {
			count += prefix + 1
		}
	}
	return count
}
