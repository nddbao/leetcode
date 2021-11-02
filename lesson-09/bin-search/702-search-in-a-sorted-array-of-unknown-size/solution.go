package solution

/*
	leetcode: https://leetcode.com/problems/search-in-a-sorted-array-of-unknown-size/
*/
/**
 * // This is the ArrayReader's API interface.
 * // You should not implement it, or speculate about its implementation
 * type ArrayReader struct {
 * }
 *
 * func (this *ArrayReader) get(index int) int {}
 */

func search(reader ArrayReader, target int) int {
	findLenCond := func(val int) bool {
		return val > 10000 && val < -10000
	}
	n := binarySearch(reader, 10000, findLenCond)

	findTargetCond := func(val int) bool {
		return val >= target
	}
	idx := binarySearch(reader, n, findTargetCond)
	if idx == n || reader.get(idx) != target {
		return -1
	}
	return idx
}

func binarySearch(reader ArrayReader, n int, condition func(val int) bool) int {
	l, h := 0, n
	for l < h {
		mid := l + (h-l)/2
		val := reader.get(mid)
		if condition(val) {
			h = mid
		} else {
			l = mid + 1
		}
	}

	return h
}

type ArrayReader struct{}

func (this *ArrayReader) get(index int) int {}
