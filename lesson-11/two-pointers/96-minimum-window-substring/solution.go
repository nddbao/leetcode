package solution

/*
	Suppose that we have a checker to count character and remove it.
	It also support for checking current range [L..R] is sub window or not.

   Based on this checker, we use two pointers L and R to keep track our range.
   For each iteration,
		+ increasing R and add s[R] to checker.
		+ if current sub range [L..R] is qualified, we will update our Min Range.
          Remove s[L] from checker as well as increasing L until [L..R] not qualified

	For building checker, we have map that contains occurrences for each char in t string.
	We also need a counter for checking sub range [L..R] have enough character not.

  Time complexity: O(m+n) whether m = len(s), n = len(t)
		Init Checker O(n)
		Find Min O(m)
  Space complexity: O(256)
*/

/*
	leetcode: https://leetcode.com/problems/minimum-window-substring/
*/
func minWindow(s string, t string) string {
	m, n := len(s), len(t)
	if n > m {
		return ""
	}

	checker := NewChecker(t)
	minL, minR := 0, m
	l, r := 0, 0
	for ; r < m; r++ {
		checker.Count(s[r])

		for checker.IsQualified() {
			minL, minR = updateMin(minL, minR, l, r)
			checker.Remove(s[l])
			l++
		} // end loop
	} // end loop

	if windowLen(minL, minR) > m {
		return ""
	}

	return s[minL : minR+1]
}

func NewChecker(t string) *Checker {
	m := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		m[t[i]]--
	}
	return &Checker{mChar: m, count: 0}
}

type Checker struct {
	mChar map[byte]int
	count int
}

func (c *Checker) Count(key byte) {
	v, ok := c.mChar[key]
	if !ok {
		return
	}

	if v == -1 {
		c.count++
	}

	c.mChar[key]++
}

func (c *Checker) Remove(key byte) {
	v, ok := c.mChar[key]
	if !ok {
		return
	}

	if v == 0 {
		c.count--
	}

	c.mChar[key]--
}

func (c *Checker) IsQualified() bool {
	return c.count == len(c.mChar)
}

func windowLen(l, r int) int {
	return r - l + 1
}

func updateMin(minL, minR, l, r int) (int, int) {
	if windowLen(l, r) < windowLen(minL, minR) {
		return l, r
	}

	return minL, minR
}
