package sets

import "github.com/divy-sh/animus/store"

func Sadd(key string, values []string) int64 {
	store.LockKeys(key)
	defer store.UnlockKeys(key)

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
	store.RLockKeys(key)
	defer store.RUnlockKeys(key)

	hashVal, ok := store.Get[string, map[string]bool](key)
	if !ok {
		return 0
	}
	return int64(len(hashVal))
}

func Sdiff(keys []string) []string {
	store.RLockKeys(keys...)
	defer store.RUnlockKeys(keys...)

	if len(keys) == 0 {
		return []string{}
	}
	baseSet, ok := store.Get[string, map[string]bool](keys[0])
	if !ok {
		return []string{}
	}
	resultSet := map[string]bool{}
	for k := range baseSet {
		resultSet[k] = true
	}
	for _, key := range keys[1:] {
		otherSet, ok := store.Get[string, map[string]bool](key)
		if !ok {
			continue
		}
		for k := range otherSet {
			delete(resultSet, k)
		}
	}
	diffValues := []string{}
	for k := range resultSet {
		diffValues = append(diffValues, k)
	}
	return diffValues
}

func Sismember(key string, value string) bool {
	store.RLockKeys(key)
	defer store.RUnlockKeys(key)

	hashVal, ok := store.Get[string, map[string]bool](key)
	if !ok {
		return false
	}
	_, exists := hashVal[value]
	return exists
}
