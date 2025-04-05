package essentias

import (
	"errors"
	"strconv"

	"github.com/divy-sh/animus/store"
)

func RPop(key string, count string) ([]string, error) {
	lock := store.GetLock(key)
	lock.Lock()
	defer lock.Unlock()
	vals, ok := store.Get[string, []string](key)
	if !ok {
		return nil, errors.New("ERR list does not exist")
	}
	cnt, err := strconv.ParseInt(count, 10, 64)
	if err != nil || cnt <= 0 || cnt > int64(len(vals)) {
		return nil, errors.New("ERR invalid count")
	}
	store.Set(key, vals[len(vals)-int(cnt):])
	return vals[len(vals)-int(cnt):], nil
}

func RPush(key string, values *[]string) {
	lock := store.GetLock(key)
	lock.Lock()
	defer lock.Unlock()
	vals, ok := store.Get[string, []string](key)
	if !ok {
		vals = *values
	} else {
		vals = append(vals, *values...)
	}
	store.Set(key, vals)
}
