package store

import (
	"fmt"
	"sort"
	"sync"
)

var (
	lock sync.Map
)

// lockKey returns a RWMutex for a given key
func (s *Store) lockKey(key any) *sync.RWMutex {
	l, ok := lock.Load(key)
	if ok {
		return l.(*sync.RWMutex)
	}
	newLock := &sync.RWMutex{}
	actual, _ := lock.LoadOrStore(key, newLock)
	return actual.(*sync.RWMutex)
}

func RLockKeys(keys ...string) []*sync.RWMutex {
	sortedKeys := sortKeys(keys)
	locks := make([]*sync.RWMutex, 0, len(sortedKeys))
	for _, key := range sortedKeys {
		l := store.lockKey(key)
		l.RLock()
		locks = append(locks, l)
	}
	return locks
}

func RUnlockKeys(keys ...string) {
	interfaceKeys := sortKeys(keys)
	for _, key := range interfaceKeys {
		l := store.lockKey(key)
		l.RUnlock()
	}
}

func LockKeys(keys ...string) []*sync.RWMutex {
	sortedKeys := sortKeys(keys)
	locks := make([]*sync.RWMutex, 0, len(sortedKeys))
	for _, key := range sortedKeys {
		l := store.lockKey(key)
		l.Lock()
		locks = append(locks, l)
	}
	return locks
}

func UnlockKeys(keys ...string) {
	interfaceKeys := sortKeys(keys)
	for _, key := range interfaceKeys {
		l := store.lockKey(key)
		l.Unlock()
	}
}

func sortKeys(keys []string) []string {
	sortedKeys := make([]string, len(keys))
	copy(sortedKeys, keys)
	sort.Slice(sortedKeys, func(i, j int) bool {
		return fmt.Sprintf("%v", sortedKeys[i]) < fmt.Sprintf("%v", sortedKeys[j])
	})
	return sortedKeys
}
