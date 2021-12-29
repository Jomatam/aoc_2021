package helpers

type StringQueue []string

func (queue StringQueue) Enqueue(value string) StringQueue {
	return append(queue, value)
}

func (queue StringQueue) Dequeue() (StringQueue, string, bool) {
	if len(queue) == 0 {
		return queue, "", false
	}
	return queue[1:], queue[0], true
}
