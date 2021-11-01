package solution

/*
	leetcode: https://leetcode.com/problems/peak-index-in-a-mountain-array/
*/

/*
	Array: [...ascending......PEAK.....descending......]

	We can check each element in array one by one. But it will take O(n).
	To improve runtime, we can use binary search to speed up our search.

	We have low and high to keep track segment that we are searching.
	Finding middle: mid = (low + high) / 2
	There are 2 two cases for arr[mid]:
		+ mid belongs to ascending part: update low = mid + 1
		+ mid blogs to descending part: update high = mid
	We keep going if low still less than high.

	We know that low belongs to ascending part, high belongs descending part.
	when low == high, that is result that we want.

	Time complexity: O(log(n))
	Space complexity: O(1)
*/

func peakIndexInMountainArray(arr []int) int {
	l, h := 0, len(arr)-1
	for l < h {
		mid := l + (h-l)/2
		if arr[mid] < arr[mid+1] {
			l = mid + 1
		} else {
			h = mid
		}
	}
	return l
}
