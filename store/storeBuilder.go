package store

import (
	"container/list"

	"github.com/google/btree"
)

/* StoreConfig - Configuration settings for Store */
type StoreConfig struct {
	MaxSize int
}

/* StoreBuilder - Allows configuring store before initialization */
type StoreBuilder struct {
	config StoreConfig
}

// NewStoreBuilder creates a new builder
func NewStoreBuilder() *StoreBuilder {
	return &StoreBuilder{config: storeConfig} // Start with default config
}

// SetMaxSize configures the max size
func (b *StoreBuilder) SetMaxSize(size int) *StoreBuilder {
	b.config.MaxSize = size
	return b
}

// Build sets the configuration and initializes the singleton
func (b *StoreBuilder) Build() {
	once.Do(func() {
		storeConfig = b.config // Apply configuration before initialization
		sharedInstance = &Store{
			dict:    make(map[any]*list.Element),
			queue:   list.New(),
			expTree: btree.New(2),
			maxSize: storeConfig.MaxSize,
		}
	})
}
