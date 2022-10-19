package stack_queue

// 单调队列
type queue struct {
	data []int
}

func (q queue) empty() bool {
	return len(q.data) == 0
}

func (q queue) front() int {
	return q.data[0]
}

// 窗口右移纳进来的值
func (q *queue) push(val int) {
	idx := 0
	for _, d := range q.data {
		if d < val {
			idx++
		} else {
			break
		}
	}
	q.data = append(q.data[idx:], val)
}

// 参数 val 是从窗口的左边弹出去的
// 每次弹出的时候，比较当前要弹出的数值是否等于队列出口元素的数值，如果相等则弹出。
func (q *queue) pop(val int) {
	if !q.empty() && q.front() == val {
		q.data = q.data[1:]
	}
}

func MaxSliding() {
	data := []int{1, 3, -1, -3, 5, 3, 6, 7}
	q := new(queue)

	k := 3

	for i := 0; i < k; i++ {
		q.push(data[i])
	}

	var result []int
	result = append(result, q.front())
	for i := k; i < len(data); i++ {
		q.pop(data[i-k])
		q.push(data[i])
		result = append(result, q.front())
	}
	println(result)
}
