package Queue

import "math"

type queue struct {
	queueList []int
	Head      int
	Tail      int
	Size      int
	Limit     int
}

func (q *queue) Init(k int) {
	q.Tail, q.Head, q.Size = 0, 0, 0
	q.Limit = k
}

func (q *queue) enQueue(v int) bool {
	if q.isFull() {
		return false
	} else {
		q.queueList[q.Tail] = v
		if q.Tail == q.Limit-1 {
			q.Tail = 0
		} else {
			q.Tail++
		}
		q.Size++
		return true
	}
}

func (q *queue) deQueue() (bool, int) {
	ans := math.MaxInt
	if q.isFull() {
		return false, ans
	} else {
		ans = q.queueList[q.Head]
		if q.Head == q.Limit-1 {
			q.Head = 0
		} else {
			q.Head++
		}
		q.Size--
		return true, ans
	}
}

func (q *queue) isFull() bool {
	if q.Size == q.Limit {
		return true
	}
	return false
}

func (q *queue) isEmpty() bool {
	if q.Size == 0 {
		return true
	}
	return false
}
