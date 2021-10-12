package solution

/*
	leetcode: https://leetcode.com/problems/k-closest-points-to-origin/
*/

func kClosest(points [][]int, k int) [][]int {
	return qSelect(points, 0, len(points)-1, k-1)
}

func qSelect(points [][]int, l, r int, k int) [][]int {
	p := doPivot(points, l, r)

	if p < k {
		return qSelect(points, p+1, r, k)
	} else if p > k {
		return qSelect(points, l, p-1, k)
	}

	return points[:p+1]
}

func doPivot(points [][]int, l, r int) int {
	pivot := distanceSquare(points[r])
	p := l
	for i := l; i <= r-1; i++ {
		if distanceSquare(points[i]) <= pivot {
			points[i], points[p] = points[p], points[i]
			p++
		}
	}
	points[p], points[r] = points[r], points[p]
	return p
}

func distanceSquare(a []int) int {
	return a[0]*a[0] + a[1]*a[1]
}
