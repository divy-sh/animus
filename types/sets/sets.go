package sets

import "github.com/divy-sh/animus/store"

func Sadd(key string, values []string) int64 {
	store.GlobalLock.Lock()
	defer store.GlobalLock.Unlock()
	hashVal, ok := store.Get[string, map[string]bool](key)
	count := 0
	if !ok {
		hashVal = map[string]bool{}
		for _, value := range values {
			hashVal[value] = true
			count++
		}
	} else {
		for _, value := range values {
			if _, exists := hashVal[value]; exists {
				continue
			}
			hashVal[value] = true
			count++
		}
	}
	store.Set(key, hashVal)
	return int64(count)
}

func Scard(key string) int64 {
	store.GlobalLock.RLock()
	defer store.GlobalLock.RUnlock()
	hashVal, ok := store.Get[string, map[string]bool](key)
	if !ok {
		return 0
	}
	return int64(len(hashVal))
}
