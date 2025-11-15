package lists

import (
	"errors"
	"strconv"

	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/store"
)

func getOrCreate(key string) *Deque[string] {
	dq, ok := store.Get[string, *Deque[string]](key)
	if !ok {
		dq = NewDeque[string](4)
		store.Set(key, dq)
	}
	return dq
}

func get(key string) (*Deque[string], error) {
	dq, ok := store.Get[string, *Deque[string]](key)
	if !ok {
		return nil, errors.New(common.ERR_LIST_NOT_FOUND)
	}
	return dq, nil
}

func Lindex(key string, index int64) (string, error) {
	store.LockKeys(key)
	defer store.UnlockKeys(key)

	dq, err := get(key)
	if err != nil {
		return "", err
	}

	if index < 0 {
		index = int64(dq.Len()) + index
	}
	v, ok := dq.Get(int(index))
	if !ok {
		return "", errors.New(common.ERR_INDEX_OUT_OF_RANGE)
	}
	return v, nil
}

func Linsert(key string, position string, pivot string, value string) (int64, error) {
	store.LockKeys(key)
	defer store.UnlockKeys(key)

	dq, err := get(key)
	if err != nil {
		return 0, err
	}

	pivotIdx := -1
	for i := 0; i < dq.Len(); i++ {
		v, _ := dq.Get(i)
		if v == pivot {
			pivotIdx = i
			break
		}
	}
	if pivotIdx == -1 {
		return int64(dq.Len()), nil
	}

	switch position {
	case "BEFORE":
		dq.InsertAt(pivotIdx, value)
	case "AFTER":
		dq.InsertAt(pivotIdx+1, value)
	default:
		return 0, errors.New(common.ERR_WRONG_ARGUMENT_COUNT)
	}

	return int64(dq.Len()), nil
}

func LLen(key string) (int64, error) {
	store.LockKeys(key)
	defer store.UnlockKeys(key)

	dq, err := get(key)
	if err != nil {
		return 0, err
	}

	return int64(dq.Len()), nil
}

func Lmove(source, destination, direction string) (string, error) {
	store.LockKeys(source, destination)
	defer store.UnlockKeys(source, destination)

	src, err := get(source)
	if err != nil || src.Len() == 0 {
		return "", errors.New(common.ERR_SOURCE_KEY_NOT_FOUND)
	}

	var val string
	var ok bool

	switch direction {
	case "RIGHT":
		val, ok = src.PopBack()
	case "LEFT":
		val, ok = src.PopFront()
	default:
		return "", errors.New(common.ERR_WRONG_ARGUMENT_COUNT)
	}
	if !ok {
		return "", errors.New(common.ERR_SOURCE_KEY_NOT_FOUND)
	}

	dest := getOrCreate(destination)

	if direction == "RIGHT" {
		dest.PushFront(val)
	} else {
		dest.PushBack(val)
	}

	return val, nil
}

func Lmpop(source string, destinations []string, direction string) (map[string]string, error) {
	lockKeys := make([]string, 0, 1+len(destinations))

	store.LockKeys(lockKeys...)
	defer store.UnlockKeys(lockKeys...)

	src, err := get(source)
	if err != nil || src.Len() == 0 {
		return nil, errors.New(common.ERR_SOURCE_KEY_NOT_FOUND)
	}

	result := make(map[string]string)

	for _, destKey := range destinations {
		if src.Len() == 0 {
			break
		}

		var val string
		var ok bool
		switch direction {
		case "RIGHT":
			val, ok = src.PopBack()
		case "LEFT":
			val, ok = src.PopFront()
		default:
			return nil, errors.New(common.ERR_WRONG_ARGUMENT_COUNT)
		}
		if !ok {
			break
		}

		dest := getOrCreate(destKey)
		if direction == "RIGHT" {
			dest.PushFront(val)
		} else {
			dest.PushBack(val)
		}

		result[destKey] = val
	}

	return result, nil
}

func LPop(key string, count string) ([]string, error) {
	store.LockKeys(key)
	defer store.UnlockKeys(key)

	dq, err := get(key)
	if err != nil {
		return nil, err
	}

	cnt, err := strconv.ParseInt(count, 10, 64)
	if err != nil || cnt <= 0 || cnt > int64(dq.Len()) {
		return nil, errors.New("ERR invalid count")
	}

	out := make([]string, cnt)
	for i := int64(0); i < cnt; i++ {
		v, _ := dq.PopFront()
		out[i] = v
	}
	return out, nil
}

func Lpos(key string, value string) (int64, error) {
	store.LockKeys(key)
	defer store.UnlockKeys(key)

	dq, err := get(key)
	if err != nil {
		return -1, err
	}

	for i := 0; i < dq.Len(); i++ {
		v, _ := dq.Get(i)
		if v == value {
			return int64(i), nil
		}
	}
	return -1, nil
}

func LPush(key string, values *[]string) {
	store.LockKeys(key)
	defer store.UnlockKeys(key)

	dq := getOrCreate(key)

	for i := len(*values) - 1; i >= 0; i-- {
		dq.PushFront((*values)[i])
	}
}

func LPushx(key string, values *[]string) error {
	store.LockKeys(key)
	defer store.UnlockKeys(key)

	dq, err := get(key)
	if err != nil {
		return err
	}

	for i := len(*values) - 1; i >= 0; i-- {
		dq.PushFront((*values)[i])
	}
	return nil
}

func Lrange(key string, start int64, stop int64) ([]string, error) {
	store.LockKeys(key)
	defer store.UnlockKeys(key)

	dq, err := get(key)
	if err != nil {
		return nil, err
	}

	l := int64(dq.Len())

	if start < 0 {
		start = l + start
	}
	if stop < 0 {
		stop = l + stop
	}
	if start < 0 {
		start = 0
	}
	if stop >= l {
		stop = l - 1
	}
	if start > stop || start >= l {
		return []string{}, nil
	}

	return dq.SliceRange(int(start), int(stop)), nil
}

func Lrem(key string, count int64, value string) (int64, error) {
	store.LockKeys(key)
	defer store.UnlockKeys(key)

	dq, err := get(key)
	if err != nil {
		return 0, err
	}

	removed := int64(0)

	if count == 0 {
		// remove all
		newdq := NewDeque[string](dq.Len())
		for i := 0; i < dq.Len(); i++ {
			v, _ := dq.Get(i)
			if v == value {
				removed++
			} else {
				newdq.PushBack(v)
			}
		}
		store.Set(key, newdq)
		return removed, nil
	}

	if count > 0 {
		// left to right
		newdq := NewDeque[string](dq.Len())
		for i := 0; i < dq.Len(); i++ {
			v, _ := dq.Get(i)
			if v == value && removed < count {
				removed++
			} else {
				newdq.PushBack(v)
			}
		}
		store.Set(key, newdq)
		return removed, nil
	}

	// count < 0 — remove from right
	newdq := NewDeque[string](dq.Len())
	toRemove := -count

	// collect values
	values := dq.ToSlice()
	for i := dq.Len() - 1; i >= 0; i-- {
		if values[i] == value && removed < toRemove {
			removed++
			values[i] = "" // skip
		}
	}

	for _, v := range values {
		if v != "" {
			newdq.PushBack(v)
		}
	}

	store.Set(key, newdq)
	return removed, nil
}

func Lset(key string, index int64, value string) error {
	store.LockKeys(key)
	defer store.UnlockKeys(key)

	dq, err := get(key)
	if err != nil {
		return err
	}

	if index < 0 {
		index = int64(dq.Len()) + index
	}
	if !dq.Set(int(index), value) {
		return errors.New(common.ERR_INDEX_OUT_OF_RANGE)
	}
	return nil
}

func Ltrim(key string, start int64, stop int64) error {
	store.LockKeys(key)
	defer store.UnlockKeys(key)

	dq, err := get(key)
	if err != nil {
		return err
	}

	l := int64(dq.Len())

	if start < 0 {
		start = l + start
	}
	if stop < 0 {
		stop = l + stop
	}
	if start < 0 {
		start = 0
	}
	if stop >= l {
		stop = l - 1
	}

	if start > stop || start >= l {
		store.Set(key, NewDeque[string](4))
		return nil
	}

	out := NewDeque[string](int(stop-start) + 1)
	for i := start; i <= stop; i++ {
		v, _ := dq.Get(int(i))
		out.PushBack(v)
	}

	store.Set(key, out)
	return nil
}

func RPop(key string, count string) ([]string, error) {
	store.LockKeys(key)
	defer store.UnlockKeys(key)

	dq, err := get(key)
	if err != nil {
		return nil, err
	}

	cnt, err := strconv.ParseInt(count, 10, 64)
	if err != nil || cnt <= 0 || cnt > int64(dq.Len()) {
		return nil, errors.New("ERR invalid count")
	}

	out := make([]string, cnt)
	for i := int64(0); i < cnt; i++ {
		v, _ := dq.PopBack()
		out[i] = v
	}
	return out, nil
}

func RPush(key string, values *[]string) {
	store.LockKeys(key)
	defer store.UnlockKeys(key)

	dq := getOrCreate(key)

	for _, v := range *values {
		dq.PushBack(v)
	}
}

func RPushx(key string, values *[]string) error {
	store.LockKeys(key)
	defer store.UnlockKeys(key)

	dq, err := get(key)
	if err != nil {
		return err
	}

	for _, v := range *values {
		dq.PushBack(v)
	}
	return nil
}
