package solution

/*
	leetcode: https://leetcode.com/problems/sort-array-by-parity-ii/
*/

func sortArrayByParityII(nums []int) []int {
	k := 0
	for i := 0; i < len(nums); i += 2 {
		if nums[i]%2 == 0 {
			nums[i], nums[k] = nums[k], nums[i]
			k += 2
		}
	}

	for i := 1; i < len(nums); i += 2 {
		if nums[i]%2 == 0 {
			nums[i], nums[k] = nums[k], nums[i]
			k += 2
		}
	}

	return nums
}

func sortArrayByParityIIV2(nums []int) []int {
	i, j := 0, 1
	for i < len(nums) && j < len(nums) {
		if nums[i]%2 == 0 {
			i += 2
			continue
		}

		if nums[j]%2 != 0 {
			j += 2
			continue
		}

		nums[i], nums[j] = nums[j], nums[i]
		i, j = i+2, j+2
	}
	return nums
}
