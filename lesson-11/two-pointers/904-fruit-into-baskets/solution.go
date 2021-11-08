package solution

/*
	leetcode: https://leetcode.com/problems/fruit-into-baskets/
*/

/*
	We move from beginning of the array. We will two pointer (left, right) to keep tracking a window.
	That window has to be satisfied condition: number of fruit types in range is not greater than 2.

    When moving to next element:
		if not satisfied condition, just shifting left pointer ahead util condition is true.
		We also calculate len of window and update max result.

	Time complexity: O(n)
	Space complexity: O(3) -> O(1)
*/

func totalFruit(fruits []int) int {
	freq := make(map[int]int)
	var max, left int
	for right, v := range fruits {
		freq[v]++
		left = updateLeft(fruits, freq, left)
		max = getMax(max, right-left+1)
	}
	return max
}

func updateLeft(fruits []int, freq map[int]int, left int) int {
	for len(freq) > 2 {
		v := fruits[left]
		c, _ := freq[v]
		if c > 1 {
			freq[v]--
		} else {
			delete(freq, v)
		}

		left++
	}
	return left
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
