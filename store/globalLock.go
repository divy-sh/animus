package store

import "sync"

var (
	GlobalLock sync.RWMutex
)
