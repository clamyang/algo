package stack_queue

import (
	"container/heap"
	"fmt"
)

// TopK 主要考察 堆 这种数据结构
func TopK() {
	arr := []int{1, 1, 1, 2, 2, 3}
	k := 2

	frequent := make(map[int]int)
	for _, val := range arr {
		counter, ok := frequent[val]
		if !ok {
			frequent[val] = 1
		} else {
			counter++
			frequent[val] = counter
		}
	}

	// 实现包含 k 个元素的小顶堆
	// 小顶堆其实是排序后的完全二叉树
	// 完全二叉树可以用数组的形式存储
	//heap.Init()

	target := new(IHeap)
	heap.Init(target)

	for key, val := range frequent {
		heap.Push(target, [2]int{key, val})
		if target.Len() > k {
			heap.Pop(target)
		}
	}

	res:=make([]int,k)
	//按顺序返回堆中的元素
	for i:=0;i<k;i++{
		res[k-i-1]=heap.Pop(target).([2]int)[0]
	}
	fmt.Println(res)
}

type IHeap [][2]int

func (h IHeap) Len() int {
	return len(h)
}

// Less 使用频率做比较
func (h IHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
