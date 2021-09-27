package main

type Queue struct {
	list []int
}

func NewQueue() *Queue {
	return &Queue{list: []int{}}
}

func (q *Queue) Enqueue(value int) {
	q.list = append(q.list, value)
}
func (q *Queue) Dequeue() int {
	if len(q.list) == 0 {
		return -1
	}
	first := q.list[0]
	//ilk sonrasını al
	q.list = q.list[1:]
	return first
}

func main() {
	//ilk giren ilk çıkar
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(3)
	q.Enqueue(5)
	println(q.Dequeue())
	println(q.Dequeue())
	println(q.Dequeue())
	println(q.Dequeue())
	println(q.Dequeue())
	println(q.Dequeue())
	println(q.Dequeue())

}
