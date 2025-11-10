package store

import (
	"math/rand"
	"time"
)

func StartExpiryCleaner() {
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
	now := time.Now().Unix()
	store.mutex.RLock()
	allKeys := store.LRUCache.Keys()
	store.mutex.RUnlock()
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
		expiryPercentage := float64(expiredCount) / float64(len(keysToCheck)) * 100
		if expiryPercentage < targetPercent {
			return
		}
	}
}

func sampleRandomKeys(keys []any, sampleSize int) []any {
	if len(keys) == 0 {
		return []any{}
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
