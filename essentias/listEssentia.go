package essentias

import (
	"errors"
	"strconv"
	"sync"
	"time"

	"github.com/divy-sh/animus/store"
)

type ListEssentia struct {
	lists store.Store[string, []string]
	lock  sync.RWMutex
}

func NewListEssentia() *ListEssentia {
	return &ListEssentia{
		lists: *store.NewStore[string, []string](),
	}
}

func (l *ListEssentia) RPop(key string, count string) ([]string, error) {
	vals, ok := l.lists.Get(key)
	if !ok {
		return nil, errors.New("ERR list does not exist")
	}
	cnt, err := strconv.ParseInt(count, 10, 64)
	if err != nil || cnt <= 0 || cnt > int64(len(vals)) {
		return nil, errors.New("ERR invalid count")
	}
	return vals[len(vals)-int(cnt):], nil
}

func (l *ListEssentia) RPush(key string, values *[]string) {
	l.lock.Lock()
	defer l.lock.Unlock()
	vals, ok := l.lists.Get(key)
	if !ok {
		vals = *values
	} else {
		vals = append(vals, *values...)
	}
	l.lists.Set(key, vals, time.Now().AddDate(1000, 0, 0))
}
