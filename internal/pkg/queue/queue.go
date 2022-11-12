package queue

import (
	"container/list"

	"github.com/wishperera/race-tracks/internal/models"
)

type Queue struct {
	list *list.List
}

func NewQueue() *Queue {
	return &Queue{
		list: list.New(),
	}
}

func (q *Queue) Enqueue(hop models.Hop) {
	q.list.PushBack(hop)
}

func (q *Queue) Dequeue() models.Hop {
	front := q.list.Front()
	if front == nil {
		return models.Hop{}
	}

	q.list.Remove(front)
	return front.Value.(models.Hop)
}

func (q *Queue) Empty() bool {
	return q.list.Front() == nil
}
