package store

import (
	"sync"
	"time"

	lru "github.com/hashicorp/golang-lru/v2"
)

type Store struct {
	LRUCache    *lru.Cache[any, any]
	stopCleaner chan struct{}
	mutex       sync.RWMutex
	isRunning   bool
}

type Value struct {
	Val any
	TTL int64
}

var (
	store *Store
)

func init() {
	cache, _ := lru.New[any, any](100000)
	store = &Store{
		LRUCache:    cache,
		stopCleaner: make(chan struct{}),
		isRunning:   false,
	}
	// // Start the cleaner automatically
	StartExpiryCleaner()
}

func Get[K comparable, V any](key K) (V, bool) {
	store.mutex.RLock()
	val, ok := store.LRUCache.Get(key)
	store.mutex.RUnlock()
	if !ok {
		var zero V
		return zero, false
	}
	value := val.(*Value)
	if value.TTL > -1 && value.TTL <= time.Now().Unix() {
		Delete(key)
		var zero V
		return zero, false
	}
	return value.Val.(V), true
}

func GetWithTTL[K comparable, V any](key K) (V, int64, bool) {
	store.mutex.RLock()
	val, ok := store.LRUCache.Get(key)
	store.mutex.RUnlock()
	if !ok {
		var zero V
		return zero, -1, false
	}
	value := val.(*Value)
	if value.TTL > -1 && value.TTL <= time.Now().Unix() {
		Delete(key)
		var zero V
		return zero, -1, false
	}
	return value.Val.(V), value.TTL, true
}

func Set[K comparable, V any](key K, value V) {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	store.LRUCache.Add(key, &Value{value, -1})
}

func SetWithTTL[K comparable, V any](key K, value V, ttl int64) {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	store.LRUCache.Add(key, &Value{value, ttl + time.Now().Unix()})
}

func Delete[K comparable](key K) {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	store.LRUCache.Remove(key)
}

func GetKeys[K comparable]() *[]K {
	keys := store.LRUCache.Keys()
	kKeys := []K{}
	for _, key := range keys {
		kKeys = append(kKeys, key.(K))
	}
	return &kKeys
}
