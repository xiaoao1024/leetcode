package solution

import (
	"testing"
)

func TestLruCache(t *testing.T) {
	cache := NewLRUCache(2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	if cache.Get(1) != 1 {
		t.Errorf("LRU cache get(1) != 1")
	}

	cache.Put(3, 3)

	if cache.Get(2) != -1 {
		t.Errorf("LRU cache get(2) != -1")
	}

	cache.Put(4, 4)

	if cache.Get(1) != -1 {
		t.Errorf("LRU cache get(1) != -1")
	}

	if cache.Get(3) != 3 {
		t.Errorf("LRU cache get(3) != 3")
	}

	if cache.Get(4) != 4 {
		t.Errorf("LRU cache get(4) != 4")
	}
}
