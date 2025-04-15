package store

import (
	"fmt"
	"time"

	lru "github.com/hashicorp/golang-lru/v2"
)

type Store struct {
	LRUCache *lru.Cache[any, any]
}

type Value struct {
	Val any
	TTL int64
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
	if !ok || (val.(*Value).TTL > -1 && val.(*Value).TTL < time.Now().Unix()) {
		fmt.Println("testing")
		Delete(key)
		var zero V
		return zero, false
	}
	return val.(*Value).Val.(V), ok
}

func Set[K comparable, V any](key K, value V) {
	store := getSharedStore()
	store.LRUCache.Add(key, &Value{value, -1})
}

func SetWithTTL[K comparable, V any](key K, value V, ttl int64) {
	store := getSharedStore()
	store.LRUCache.Add(key, &Value{value, ttl + time.Now().Unix()})
}

func Delete[K comparable](key K) {
	store := getSharedStore()
	store.LRUCache.Remove(key)
}
