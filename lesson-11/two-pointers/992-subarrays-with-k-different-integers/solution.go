package solution

/*
	leetcode: https://leetcode.com/problems/subarrays-with-k-different-integers/
*/
func subarraysWithKDistinct(nums []int, k int) int {
	var count int
	checker := NewChecker()
	l, r := 0, 0
	for ; r < len(nums); r++ {
		checker.Add(nums[r])
		d := checker.DistinctNumbers()
		if d < k {
			continue
		}

		l = updateLeft(nums, checker, k, l, r)
		count += countSubArray(nums, checker, k, l, r)
	}
	return count
}

func countSubArray(nums []int, checker *Checker, k, l, r int) int {
	var count int
	if checker.DistinctNumbers() == k {
		key := nums[l]
		checker.Remove(key)
		count = 1 + countSubArray(nums, checker, k, l+1, r)
		checker.Add(key)
	}

	return count
}

func updateLeft(nums []int, checker *Checker, k, l, r int) int {
	for checker.DistinctNumbers() > k {
		checker.Remove(nums[l])
		l++
	}

	return l
}

func NewChecker() *Checker {
	return &Checker{arr: make([]int, 2*1e4+1)}
}

type Checker struct {
	arr      []int
	distinct int
}

func (c *Checker) Add(key int) {
	c.arr[key]++
	if c.arr[key] == 1 {
		c.distinct++
	}
}

func (c *Checker) Remove(key int) {
	c.arr[key]--
	if c.arr[key] == 0 {
		c.distinct--
	}
}

func (c *Checker) DistinctNumbers() int {
	return c.distinct
}
