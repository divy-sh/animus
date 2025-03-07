package types

import (
	"errors"
	"strconv"
	"time"

	"github.com/divy-sh/animus/store"
)

type ListType struct {
	lists store.Store[string, []string]
}

func NewListType() *ListType {
	return &ListType{
		lists: *store.NewStore[string, []string](),
	}
}

func (l *ListType) RPop(key string, count string) ([]string, error) {
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

func (l *ListType) RPush(key string, values *[]string) {
	vals, ok := l.lists.Get(key)
	if !ok {
		vals = *values
	} else {
		vals = append(vals, *values...)
	}
	l.lists.Set(key, vals, time.Now().AddDate(1000, 0, 0))
}
