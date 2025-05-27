package inmemory

import (
	"sync"

	"github.com/kiritocyanpine/go-tiny-url/persistant"
)

type InMemory struct {
	db map[string]any
	m  *sync.RWMutex
}

func CreateDB() *InMemory {
	return &InMemory{
		db: make(map[string]any),
		m:  &sync.RWMutex{},
	}
}

func (m *InMemory) Get(key string) (any, error) {
	m.m.RLock()
	value, exist := m.db[key]
	if !exist {
		return nil, persistant.ErrKeyNotFound
	}
	m.m.RUnlock()

	return value, nil
}

func (m *InMemory) Set(key string, value any) error {
	m.m.Lock()
	_, exist := m.db[key]
	if exist {
		return persistant.ErrKeyCollision
	}

	m.db[key] = value

	m.m.Unlock()

	return nil
}
