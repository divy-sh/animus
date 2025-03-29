package essentias

import (
	"errors"
	"strconv"
	"sync"
	"time"

	"github.com/divy-sh/animus/store"
)

type ListEssentia struct {
	locks sync.Map
}

func NewListEssentia() *ListEssentia {
	return &ListEssentia{}
}

func (l *ListEssentia) getLock(key string) *sync.RWMutex {
	actual, _ := l.locks.LoadOrStore(key, &sync.RWMutex{})
	return actual.(*sync.RWMutex)
}

func (l *ListEssentia) RPop(key string, count string) ([]string, error) {
	lock := l.getLock(key)
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
	store.Set(key, vals[len(vals)-int(cnt):], time.Now().AddDate(1000, 0, 0))
	return vals[len(vals)-int(cnt):], nil
}

func (l *ListEssentia) RPush(key string, values *[]string) {
	lock := l.getLock(key)
	lock.Lock()
	defer lock.Unlock()
	vals, ok := store.Get[string, []string](key)
	if !ok {
		vals = *values
	} else {
		vals = append(vals, *values...)
	}
	store.Set(key, vals, time.Now().AddDate(1000, 0, 0))
}
