package main

import "container/list"

type LRUCache struct {
	queue *list.List
	cache map[string]*list.Element
	cSize int
}

func (l *LRUCache) Check(s string) bool {
	if _, ok := l.cache[s]; ok {
		return true
	} else {
		return false
	}
}
func (l *LRUCache) Add(s string) {
	if l.Check(s) == true {
		val := l.cache[s]
		l.queue.MoveToFront(val)
	} else {
		if l.queue.Len() == l.cSize {
			l.Remove(s)
			l.queue.PushFront(s)
			l.cache[s] = l.queue.Front()
		} else {
			l.queue.PushFront(s)
			l.cache[s] = l.queue.Front()
		}
	}
}
func (l *LRUCache) Remove(s string) {
	for key, value := range l.cache {
		if value == l.queue.Back() {
			delete(l.cache, key)
			break
		}
	}
	l.queue.Remove(l.queue.Back())

}
