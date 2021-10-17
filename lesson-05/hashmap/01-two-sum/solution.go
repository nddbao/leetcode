package solution

func twoSum(nums []int, target int) []int {
	distinct := make(map[int]int)
	for i1, v := range nums {
		i2, exist := distinct[target-v]
		if exist {
			return []int{i1, i2}
		}

		distinct[v] = i1
	}

	return nil
}
