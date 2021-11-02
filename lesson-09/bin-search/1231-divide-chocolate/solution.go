package solution

/*
	https://leetcode.com/problems/divide-chocolate/
*/

func maximizeSweetness(sweetness []int, k int) int {
	l, h := 0, int(10)
	for l < h {
		mid := l + (h-l)/2
		if !condition(sweetness, k, mid) {
			h = mid
		} else {
			l = mid + 1
		}
	}

	return h - 1
}

func condition(sweetness []int, k, target int) bool {
	count := 0
	sum := 0
	for _, v := range sweetness {
		sum += v
		if sum >= target {
			count++
			sum = 0
		}
	}

	return count >= k+1
}
