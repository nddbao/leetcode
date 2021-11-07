package solution

/*
	leetcode: https://leetcode.com/problems/longest-mountain-in-array/
*/

func longestMountain(arr []int) int {
	if len(arr) < 3 {
		return 0
	}

	max, leftIdx := 0, -1
	for i := 0; i < len(arr); i++ {
		if isRightRoot(arr, i) && leftIdx >= 0 {
			max = getMax(max, i-leftIdx+1)
			leftIdx = -1
		}

		if i+1 < len(arr)-1 && arr[i] == arr[i+1] {
			leftIdx = -1
			continue
		}

		if isLeftRoot(arr, i) {
			leftIdx = i
		}
	} // end loop

	return max
}

func isRightRoot(a []int, i int) bool {
	return (i == len(a)-1 && a[i-1] > a[i]) ||
		(i-1 >= 0 && i+1 <= len(a)-1 && a[i-1] > a[i] && a[i] <= a[i+1])
}

func isLeftRoot(a []int, i int) bool {
	return (i == 0 && a[i] < a[i+1]) ||
		(i-1 >= 0 && i+1 <= len(a)-1 && a[i-1] >= a[i] && a[i] < a[i+1])
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
