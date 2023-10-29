package heap

import (
	"container/heap"
	"math"
)

/*
MaxKelements2530
2530. Maximal Score After Applying K Operations
*/
//func MaxKelements2530(nums []int, k int) int64 {
//	pq := (*PriorityQueue)(&nums)
//	heap.Init(pq)
//
//	var ans int64
//	for i := 0; i < k; i++ {
//		x := heap.Pop(pq).(int)
//		ans += int64(x)
//		heap.Push(pq, (x+2)/3)
//	}
//	return ans
//}
//
//type PriorityQueue []int
//
//func (pq PriorityQueue) Len() int {
//	return len(pq)
//}
//
//func (pq PriorityQueue) Less(i, j int) bool {
//	return pq[i] > pq[j]
//}
//
//func (pq PriorityQueue) Swap(i, j int) {
//	pq[i], pq[j] = pq[j], pq[i]
//}
//
//func (pq *PriorityQueue) Push(x any) {
//	*pq = append(*pq, x.(int))
//}
//
//func (pq *PriorityQueue) Pop() any {
//	old := *pq
//	n := len(old)
//	item := old[n-1]
//	*pq = old[:n-1]
//	return item
//}

type priorityQueue []int

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x any) {
	*pq = append(*pq, x.(int))
}

func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

/*
PickGifts2558
2558. Take Gifts From the Richest Pile
*/
func PickGifts2558(gifts []int, k int) int64 {
	pq := (*priorityQueue)(&gifts)
	heap.Init(pq)

	for i := 0; i < k; i++ {
		heap.Push(pq, int(math.Sqrt(float64(heap.Pop(pq).(int)))))
	}
	ans := int64(0)
	for _, gift := range gifts {
		ans += int64(gift)
	}
	return ans
}
