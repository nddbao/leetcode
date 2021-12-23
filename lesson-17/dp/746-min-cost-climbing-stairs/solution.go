package solution

/*
	leetcode: https://leetcode.com/problems/min-cost-climbing-stairs
*/

/*
	Time complexity: O(N) whether N is len of cost
	Space coplexity: O(1)
*/

func minCostClimbingStairs(cost []int) int {
	t0, t1 := cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		t := getMin(t0, t1) + cost[i]
		t0, t1 = t1, t
	}

	return getMin(t0, t1)
}

func getMin(a, b int) int {
	if a > b {
		return b
	}

	return a
}
