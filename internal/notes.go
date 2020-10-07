package internal

import (
	"sync"
)

type noteQueue struct {
	mu    sync.Mutex
	notes []interface{}
}

func (queue *noteQueue) Add(note interface{}) (one bool) {
	queue.mu.Lock()
	queue.notes = append(queue.notes, note)
	length := len(queue.notes)
	queue.mu.Unlock()
	return length == 1
}
