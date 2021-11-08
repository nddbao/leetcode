package solution

/*
	leetcode: https://leetcode.com/problems/fruit-into-baskets/
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
