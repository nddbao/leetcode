package solution

/*
	leetcode: https://leetcode.com/problems/longest-substring-with-at-most-k-distinct-characters/
*/

/*
	We define a Frequencies to keep occurrences of character using []int with 256 length.
	We use two pointers to keep track range [left ... right] that have number of unique characters <= k.
    Every time, moving right pointer ahead by one. We add one to its occurrence in Freq.
	Then we check numberOfUniques in Freq:
		+ if numberOfUniques <= k: do nothing
		+ otherwise: we will move left by one until numberOfUniques <= k.
					We also need to update its occurrence in Freq.
    We have max variable to keep track our longest length.
	Each time we move for right pointer, will update it.

	Time complexity: O(n)
	Space complexity: O(256) -> O(1)
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
