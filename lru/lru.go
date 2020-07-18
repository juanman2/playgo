// Copyright 2020 Juan Tellez All rights reserved.

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

package lrucache

import "container/list"

// LRUCache is an Implementation of LRUCache with O(1) map fetch
type LRUCache struct {
	m   map[int]*list.Element
	l   *list.List
	cap int
}

// KV is Key value type
type KV struct {
	k int
	v int
}

// Constructor provides a way to create the cache, passing capacity
func Constructor(capacity int) LRUCache {
	cache := LRUCache{nil, nil, capacity}
	cache.m = make(map[int]*list.Element)
	cache.l = list.New()

	if capacity < 0 {
		cache.cap = 0
	}

	return cache
}

// Get the value (will always be positive) of the key if the key
// exists in the cache, otherwise return -1.
func (this *LRUCache) Get(key int) int {
	e := this.m[key]

	if e == nil {
		return -1
	}
	this.l.MoveToFront(e)
	kv := e.Value.(*KV)
	return kv.v
}

// Put adds an entry to the LRU. Set or insert the value if the key is
// not already present. When the cache reached its capacity, it should
// invalidate the least recently used item before inserting a new
// item
func (this *LRUCache) Put(key int, value int) {

	e, ok := this.m[key]

	// if it already exists set and return
	if ok {
		kv := e.Value.(*KV)
		kv.v = value
		this.l.MoveToFront(e)
		return
	}

	// insert

	// Reached max capacity, delete last
	if len(this.m) == this.cap {
		e = this.l.Back()
		this.l.Remove(e)
		kv := e.Value.(*KV)
		delete(this.m, kv.k)
	}

	kv := KV{key, value}
	e = this.l.PushFront(&kv)
	this.m[key] = e
}
