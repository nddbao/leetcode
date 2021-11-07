package solution

/*
	leetcode: https://leetcode.com/problems/longest-mountain-in-array/
*/

/*
	For a mountain, we need find leftRoot and rightRoot to calculate mountain's length.
	LeftRoot:
 		+ less or equal to previous element
		+ and less than next element.
	RightRoot:
		+ less or equal to next element
		+ and less than previous element.

	We go through each element to check it.
	+ rightRoot => check we have leftRoot or not. If yes, then compute length and update max.
	+ leftRoot => just record index of leftRoot
	+ check valid mountain or not. If no, reset index of leftRoot.
		For example: [0 1 2 2 1] -> not valid mountain

	Time Complexity: O(n)
	Space Complexity: O(1)

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

		if isLeftRoot(arr, i) {
			leftIdx = i
		}

		// check still valid mountain or not
		if i+1 < len(arr)-1 && arr[i] == arr[i+1] {
			leftIdx = -1
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
