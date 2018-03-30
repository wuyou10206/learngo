package queue

//这是一个后进先出的队列
type Queue []interface{}

//向队列中添加数据
//   Push(123)
func (q *Queue) Push(value interface{}) {
	*q = append(*q, value)
}

//从队列中取出一个数据并从队列中删除
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	// return head.(int)
	return head
}

//判断队列是否为空
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
