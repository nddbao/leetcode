package solution

/*
	leetcode: https://leetcode.com/problems/subarray-sums-divisible-by-k/
*/

/*
	We do preSum of array and then count frequency of preSum  when mod k.
	We have:
		val =(a%k + r) - (b%k + r)
	-> val will be divisible by k
	To find the answer, we will sum up all frequency.

	We also remember to add freq[0] .

*/
func subarraysDivByK(nums []int, k int) int {
	freq := make([]int, k)
	for i := 0; i < len(nums); i++ {
		if i != 0 {
			nums[i] += nums[i-1]
		}

		r := getRemainder(nums[i], k)
		freq[r]++
	}

	result := freq[0]
	for _, v := range freq {
		if v > 1 {
			result += (v - 1) * v / 2
		}
	}

	return result
}

func getRemainder(a, k int) int {
	// r := a % k
	// if r < 0 {
	//     return k+r
	// }
	// return r

	return (a%k + k) % k
}
