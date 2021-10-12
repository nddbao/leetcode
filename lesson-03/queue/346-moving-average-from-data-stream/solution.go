package solution

/*
	leetcode: https://leetcode.com/problems/moving-average-from-data-stream/
*/

type MovingAverage struct {
	q   *CircleQueue
	sum int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		q: &CircleQueue{a: make([]int, size)},
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if this.q.IsFull() {
		this.sum -= this.q.Dequeue()
	}

	this.q.Enqueue(val)
	this.sum += val
	return float64(this.sum) / float64(this.q.Len())
}

type CircleQueue struct {
	a           []int
	start, size int
}

func (c *CircleQueue) IsFull() bool {
	return c.Len() == len(c.a)
}

func (c *CircleQueue) IsEmpty() bool {
	return c.Len() == 0
}

func (c *CircleQueue) Len() int {
	return c.size
}

func (c *CircleQueue) Enqueue(v int) {
	idx := (c.start + c.size) % len(c.a)
	c.a[idx] = v
	c.size++
}

func (c *CircleQueue) Dequeue() int {
	v := c.a[c.start]
	c.start = (c.start + 1) % len(c.a)
	c.size--
	return v
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
