package generics

import (
	"errors"
	"time"

	"github.com/divy-sh/animus/store"
)

func Copy(source, destination string) (int, error) {
	lock := store.GetLock(source)
	lock.RLock()
	value, ok := store.Get[any, any](source)
	lock.RUnlock()

	if !ok {
		return 0, errors.New("ERR source key not found, or expired")
	}

	lock2 := store.GetLock(destination)
	lock2.Lock()
	defer lock2.Unlock()
	store.Set(destination, value, time.Now().AddDate(1000, 0, 0))
	return 1, nil
}
