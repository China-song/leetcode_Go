package queue

import "container/list"

type FrontMiddleBackQueue struct {
	//a []int
	left  *list.List
	right *list.List
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{
		left:  list.New(),
		right: list.New(),
	}
}

// 0 <= q.right.Len() - q.left.Len() <= 1
func (q *FrontMiddleBackQueue) balance() {
	if q.left.Len() > q.right.Len() {
		q.right.PushFront(q.left.Remove(q.left.Back()))
	} else if q.right.Len() > q.left.Len()+1 {
		q.left.PushBack(q.right.Remove(q.right.Front()))
	}
}

func (q *FrontMiddleBackQueue) PushFront(val int) {
	q.left.PushFront(val)
	q.balance()
}

func (q *FrontMiddleBackQueue) PushMiddle(val int) {
	q.left.PushBack(val)
	q.balance()
}

func (q *FrontMiddleBackQueue) PushBack(val int) {
	q.right.PushBack(val)
	q.balance()
}

func (q *FrontMiddleBackQueue) PopFront() (val int) {
	if q.right.Len() == 0 {
		return -1
	}
	if q.left.Len() > 0 {
		val = q.left.Remove(q.left.Front()).(int)
	} else {
		val = q.right.Remove(q.right.Front()).(int)
	}
	q.balance()
	return val
}

func (q *FrontMiddleBackQueue) PopMiddle() (val int) {
	if q.right.Len() == 0 {
		return -1
	}
	if q.left.Len() == q.right.Len() {
		val = q.left.Remove(q.left.Back()).(int)
	} else {
		val = q.right.Remove(q.right.Front()).(int)
	}
	return val
}

func (q *FrontMiddleBackQueue) PopBack() (val int) {
	if q.right.Len() == 0 {
		return -1
	}
	val = q.right.Remove(q.right.Back()).(int)
	q.balance()
	return val
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */
