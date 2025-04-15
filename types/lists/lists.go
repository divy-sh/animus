package lists

import (
	"errors"
	"strconv"

	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/store"
)

func RPop(key string, count string) ([]string, error) {
	store.GlobalLock.Lock()
	defer store.GlobalLock.Unlock()
	vals, ok := store.Get[string, []string](key)
	if !ok {
		return nil, errors.New(common.ERR_LIST_NOT_FOUND)
	}
	cnt, err := strconv.ParseInt(count, 10, 64)
	if err != nil || cnt <= 0 || cnt > int64(len(vals)) {
		return nil, errors.New("ERR invalid count")
	}
	store.Set(key, vals[len(vals)-int(cnt):])
	return vals[len(vals)-int(cnt):], nil
}

func RPush(key string, values *[]string) {
	store.GlobalLock.Lock()
	defer store.GlobalLock.Unlock()
	vals, ok := store.Get[string, []string](key)
	if !ok {
		vals = *values
	} else {
		vals = append(vals, *values...)
	}
	store.Set(key, vals)
}
