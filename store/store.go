package store

import (
	lru "github.com/hashicorp/golang-lru/v2"
)

type Store struct {
	LRUCache *lru.Cache[any, any]
}

var (
	sharedInstance *Store
)

func getSharedStore() *Store {
	if sharedInstance == nil {
		sharedInstance = &Store{}
		sharedInstance.LRUCache, _ = lru.New[any, any](100000)
	}
	return sharedInstance
}

func Get[K comparable, V any](key K) (V, bool) {
	store := getSharedStore()
	val, ok := store.LRUCache.Get(key)
	if !ok {
		var zero V
		return zero, ok
	}
	return val.(V), ok
}

func Set[K comparable, V any](key K, value V) {
	store := getSharedStore()
	store.LRUCache.Add(key, value)
}

func SetWithTTL[K comparable, V any](key K, value V, ttl int64) {
	store := getSharedStore()
	store.LRUCache.Add(key, value)
}

func Delete[K comparable](key K) {
	store := getSharedStore()
	store.LRUCache.Remove(key)
}
