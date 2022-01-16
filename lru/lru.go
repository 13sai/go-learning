package lru

import (
	"container/list"
)

type Link struct {
	ID         int
	ActiveTime int
}

type Lru struct {
	maxSize int
	list    *list.List
	cache   map[*Link]*list.Element
}

func newLru(max int) *Lru {
	return &Lru{
		maxSize: max,
		cache:   map[*Link]*list.Element{},
		list:    list.New(),
	}
}

func (l *Lru) Push(key *Link) {
	if e, ok := l.cache[key]; ok {
		l.list.MoveToFront(e)
	} else {
		row := l.list.PushFront(key)
		l.cache[key] = row
	}
	for l.maxSize > 0 && l.list.Len() > l.maxSize {
		l.removePassive()
	}
}

func (l *Lru) CheckPassive() (*Link, bool) {
	e := l.list.Back()
	if e == nil {
		return nil, false
	}
	link := e.Value.(*Link)
	return link, true
}

func (l *Lru) Remove(key *Link) {
	if e, ok := l.cache[key]; ok {
		l.list.Remove(e)
		delete(l.cache, key)
	}
}

func (l *Lru) Len() int {
	return l.list.Len()
}

func (l *Lru) removePassive() {
	e := l.list.Back()
	l.list.Remove(e)
	delete(l.cache, e.Value.(*Link))
}
