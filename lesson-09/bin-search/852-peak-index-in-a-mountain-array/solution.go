package solution

/*
	leetcode: https://leetcode.com/problems/peak-index-in-a-mountain-array/
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
