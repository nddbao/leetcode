package solution

/*
	leetcode: https://leetcode.com/problems/longest-substring-with-at-most-two-distinct-characters/
*/

/*
	Similar to 340

	Time complexity: O(n)
	Space complexity: O(256) -> O(1)
*/
func lengthOfLongestSubstringTwoDistinct(s string) int {
	freq := NewFrequencies()

	var max int
	left, right := 0, 0
	for ; right < len(s); right++ {
		c := s[right]
		freq.Increase(c)

		for freq.NumberOfUniques() > 2 {
			c := s[left]
			freq.Decrease(c)
			left++
		}
		max = getMax(max, right-left+1)
	}
	return max
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func NewFrequencies() *Frequencies {
	return &Frequencies{a: make([]int, 256)}
}

type Frequencies struct {
	a     []int
	count int
}

func (f *Frequencies) Decrease(c byte) {
	key := getKey(c)

	if f.a[key] == 1 {
		f.count--
	}
	f.a[key]--
}

func (f *Frequencies) Increase(c byte) {
	key := getKey(c)

	if f.a[key] == 0 {
		f.count++
	}
	f.a[key]++
}

func (f *Frequencies) NumberOfUniques() int {
	return f.count
}

func getKey(c byte) int {
	return int(c)
}
