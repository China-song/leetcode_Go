package heap

import "container/heap"

/*
MaxKelements2530
2530. Maximal Score After Applying K Operations
*/
func MaxKelements2530(nums []int, k int) int64 {
	pq := (*PriorityQueue)(&nums)
	heap.Init(pq)

	var ans int64
	for i := 0; i < k; i++ {
		x := heap.Pop(pq).(int)
		ans += int64(x)
		heap.Push(pq, (x+2)/3)
	}
	return ans
}

type PriorityQueue []int

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(int))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}
