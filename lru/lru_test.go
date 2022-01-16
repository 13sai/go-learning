package lru

import "testing"

func TestLRU(t *testing.T) {
	lru := newLru(10)

	for i := 0; i < 20; i++ {
		link1 := &Link{i, i + 1642262400}
		lru.Push(link1)
	}

	old, _ := lru.CheckPassive()
	t.Logf("CheckPassive:%+v", old)
	t.Log("len=", lru.Len())

	for true {
		item, ok := lru.CheckPassive()
		if !ok {
			break
		}
		t.Logf("Remove:%+v", item)
		lru.Remove(item)
	}

	old, _ = lru.CheckPassive()
	t.Logf("CheckPassive:%+v", old)
	t.Log("len=", lru.Len())
}
