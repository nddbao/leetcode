package solution

/*
	leetcode: https://leetcode.com/problems/number-of-1-bits/
*/

func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		count++
		num = clearLastSetBit(num)
	}
	return count
}

func clearLastSetBit(num uint32) uint32 {
	return num & (num - 1)
}
