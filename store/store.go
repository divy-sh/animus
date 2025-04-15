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

// StartExpiryCleaner starts the Redis-like expiry cleaner in a separate goroutine
func StartExpiryCleaner() {
	store := getSharedStore()
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.isRunning {
		return // Already running
	}

	store.isRunning = true
	store.stopCleaner = make(chan struct{}) // Reset channel
	go expiryCleanerLoop(store)
}

// StopExpiryCleaner stops the expiry cleaner goroutine
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
	ticker := time.NewTicker(100 * time.Millisecond) // Redis checks keys 10 times per second
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

// cleanExpiredKeys implements the Redis probabilistic key expiration algorithm
func cleanExpiredKeys() {
	const (
		// Sampling parameters based on Redis defaults
		sampleSize    = 20   // Number of keys to check per cycle
		targetPercent = 25.0 // If more than 25% of keys are expired, repeat
		maxIterations = 3    // Maximum number of iterations to prevent infinite loops
	)

	store := getSharedStore()
	now := time.Now().Unix()

	// Get a snapshot of all keys
	store.mutex.RLock()
	allKeys := store.LRUCache.Keys()
	store.mutex.RUnlock()

	if len(allKeys) == 0 {
		return // No keys to check, exit immediately
	}

	// Limit the number of iterations to prevent infinite loops
	for iteration := 0; iteration < maxIterations; iteration++ {
		// Choose random keys to check
		keysToCheck := sampleRandomKeys(allKeys, sampleSize)
		if len(keysToCheck) == 0 {
			return // No keys to check, exit immediately
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

		// If no keys were sampled or checked, exit
		if len(keysToCheck) == 0 {
			return
		}

		// Calculate expiry percentage
		expiryPercentage := float64(expiredCount) / float64(len(keysToCheck)) * 100

		// If expiry percentage is less than target, stop the cycle
		if expiryPercentage < targetPercent {
			return
		}

		// Otherwise, continue with next iteration
		// We'll eventually exit due to maxIterations limit
	}
}

// sampleRandomKeys selects random keys from the available keys
func sampleRandomKeys(keys []any, sampleSize int) []any {
	if len(keys) == 0 {
		return nil // No keys to sample, return nil
	}

	if len(keys) <= sampleSize {
		return keys // Return all keys if we have fewer than sampleSize
	}

	// Create a copy to avoid modifying the original slice
	shuffled := make([]any, len(keys))
	copy(shuffled, keys)

	// Shuffle using Fisher-Yates algorithm
	for i := range shuffled {
		j := rand.Intn(i + 1)
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}

	return shuffled[:sampleSize]
}
