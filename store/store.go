package store

// TODO fix the auto deletion algorithm
import (
	"math/rand"
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
	sharedInstance *Store
	once           sync.Once
)

func getSharedStore() *Store {
	once.Do(func() {
		cache, _ := lru.New[any, any](100000)
		sharedInstance = &Store{
			LRUCache:    cache,
			stopCleaner: make(chan struct{}),
			isRunning:   false,
		}
		// // Start the cleaner automatically
		// StartExpiryCleaner()
	})
	return sharedInstance
}

func Get[K comparable, V any](key K) (V, bool) {
	store := getSharedStore()
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

func Set[K comparable, V any](key K, value V) {
	store := getSharedStore()
	store.mutex.Lock()
	defer store.mutex.Unlock()
	store.LRUCache.Add(key, &Value{value, -1})
}

func SetWithTTL[K comparable, V any](key K, value V, ttl int64) {
	store := getSharedStore()
	store.mutex.Lock()
	defer store.mutex.Unlock()
	store.LRUCache.Add(key, &Value{value, ttl + time.Now().Unix()})
}

func Delete[K comparable](key K) {
	store := getSharedStore()
	store.mutex.Lock()
	defer store.mutex.Unlock()
	store.LRUCache.Remove(key)
}

func StartExpiryCleaner() {
	store := getSharedStore()
	store.mutex.Lock()
	defer store.mutex.Unlock()
	if store.isRunning {
		return
	}
	store.isRunning = true
	store.stopCleaner = make(chan struct{})
	go expiryCleanerLoop(store)
}

func StopExpiryCleaner() {
	store := getSharedStore()
	store.mutex.Lock()
	defer store.mutex.Unlock()
	if !store.isRunning {
		return
	}
	store.isRunning = false
	close(store.stopCleaner)
}

func expiryCleanerLoop(store *Store) {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			cleanExpiredKeys()
		case <-store.stopCleaner:
			return
		}
	}
}

func cleanExpiredKeys() {
	const (
		sampleSize    = 20
		targetPercent = 25.0
		maxIterations = 3
	)
	store := getSharedStore()
	now := time.Now().Unix()
	store.mutex.RLock()
	allKeys := store.LRUCache.Keys()
	store.mutex.RUnlock()
	if len(allKeys) == 0 {
		return
	}
	for iteration := 0; iteration < maxIterations; iteration++ {
		keysToCheck := sampleRandomKeys(allKeys, sampleSize)
		if len(keysToCheck) == 0 {
			return
		}
		// Count expired keys and delete them
		expiredCount := 0
		for _, key := range keysToCheck {
			store.mutex.RLock()
			val, ok := store.LRUCache.Get(key)
			store.mutex.RUnlock()
			if !ok {
				continue
			}
			value := val.(*Value)
			if value.TTL > -1 && value.TTL < now {
				expiredCount++
				store.mutex.Lock()
				store.LRUCache.Remove(key)
				store.mutex.Unlock()
			}
		}
		if len(keysToCheck) == 0 {
			return
		}
		expiryPercentage := float64(expiredCount) / float64(len(keysToCheck)) * 100
		if expiryPercentage < targetPercent {
			return
		}
	}
}

func sampleRandomKeys(keys []any, sampleSize int) []any {
	if len(keys) == 0 {
		return nil
	}
	if len(keys) <= sampleSize {
		return keys
	}
	shuffled := make([]any, len(keys))
	copy(shuffled, keys)
	for i := range shuffled {
		j := rand.Intn(i + 1)
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}
	return shuffled[:sampleSize]
}
