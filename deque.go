package dsa

import (
	"container/list"
)

type Deque struct {
	data *list.List
}

func NewDeque() *Deque {
	return &Deque{data: list.New()}
}

func (d *Deque) Append(val int) {
	d.data.PushBack(val)
}

func (d *Deque) Popleft() (int, bool) {
	if d.data.Len() == 0 {
		return 0, false
	}
	elem := d.data.Front()
	d.data.Remove(elem)
	return elem.Value.(int), true
}
