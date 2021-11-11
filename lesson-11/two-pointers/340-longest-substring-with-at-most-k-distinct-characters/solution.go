package solution

/*
	leetcode: https://leetcode.com/problems/longest-substring-with-at-most-k-distinct-characters/
*/

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	if k == 0 {
		return 0
	}

	freq := NewFrequencies()
	var max int
	left, right := 0, 0
	for ; right < len(s); right++ {
		c := int(s[right])
		freq.Increase(c)

		for freq.NumberOfUniques() > k {
			c := int(s[left])
			freq.Decrease(c)
			left++
		}

		max = getMax(max, right-left+1)
	}

	return max
}

func NewFrequencies() *Frequencies {
	return &Frequencies{a: make([]int, 256)}
}

type Frequencies struct {
	a     []int
	count int
}

func (d *Frequencies) Increase(key int) {
	d.a[key]++
	if d.a[key] == 1 {
		d.count++
	}
}

func (d *Frequencies) Decrease(key int) {
	d.a[key]--
	if d.a[key] == 0 {
		d.count--
	}
}

func (d *Frequencies) NumberOfUniques() int {
	return d.count
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
