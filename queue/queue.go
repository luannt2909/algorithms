package queue

type Queue struct {
	Items []interface{}
}

func (q *Queue) Enqueue(item interface{}) *Queue {
	q.Items = append(q.Items, item)
	return q
}

func (q *Queue) Len() int {
	return len(q.Items)
}

func (q *Queue) Empty() bool {
	return q.Len() <= 0
}

func (q *Queue) Dequeue() interface{} {
	if q.Len() <= 0 {
		return nil
	}
	item := q.Items[0]
	q.Items = q.Items[1:]
	return item
}
